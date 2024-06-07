package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/aws/jsii-runtime-go/runtime"
)

type PipelineStackProps struct {
	StackProps          cdk.StackProps
	PipelineName        string
	RepositoryName      string
	ServiceName         string
	ConnectionArnImport string
}

func NewPipelineStack(scope constructs.Construct, id string, props *PipelineStackProps) cdk.Stack {
	var sprops cdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := cdk.NewStack(scope, &id, &sprops)

	githubConnectionArn := cdk.Fn_ImportValue(jsii.String(props.ConnectionArnImport))

	// pipeline buckets for artifacts and caching
	artifactBucket := awss3.NewBucket(stack, jsii.String("ArtifactBucket"), &awss3.BucketProps{
		EnforceSSL: jsii.Bool(true),
		LifecycleRules: &[]*awss3.LifecycleRule{
			{
				Enabled:    jsii.Bool(true),
				Expiration: awscdk.Duration_Days(jsii.Number(7)),
			},
		},
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY, // don't use this removal polcicy for real world scenarios unless you are really sure
	})
	cacheBucket := awss3.NewBucket(stack, jsii.String("CacheBucket"), &awss3.BucketProps{
		EnforceSSL: jsii.Bool(true),
		LifecycleRules: &[]*awss3.LifecycleRule{
			{
				Enabled:    jsii.Bool(true),
				Expiration: awscdk.Duration_Days(jsii.Number(7)),
			},
		},
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY, // don't use this removal polcicy for real world scenarios unless you are really sure
	})

	// the main pipeline
	pipeline := pipelines.NewCodePipeline(stack, jsii.String("MainPipeline"), &pipelines.CodePipelineProps{
		SelfMutation: jsii.Bool(true),
		SynthCodeBuildDefaults: &pipelines.CodeBuildOptions{
			BuildEnvironment: defaultBuidEnv(),
			PartialBuildSpec: defaultBuildRuntimes(),
		},
		Synth: pipelines.NewShellStep(jsii.String("Synth"), &pipelines.ShellStepProps{
			Input: pipelines.CodePipelineSource_Connection(jsii.String(props.RepositoryName), jsii.String("main"), &pipelines.ConnectionSourceOptions{
				ConnectionArn: githubConnectionArn,
				TriggerOnPush: jsii.Bool(true),
			}),
			Commands: &[]*string{
				jsii.String("go version"),
				jsii.String("npm install -g aws-cdk"),
				jsii.String("go mod download"),
				jsii.String("go mod tidy"),
				// jsii.String("go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"),
				// jsii.String("${GOPATH}/bin/golangci-lint run ./..."),
				jsii.String("npx cdk synth >> /dev/null"),
			},
		}),
		ArtifactBucket: artifactBucket,
		PipelineName:   jsii.String(props.PipelineName),
	})

	wave := pipeline.AddWave(jsii.String("TestLambda"), nil)
	buildStep := pipelines.NewCodeBuildStep(jsii.String("Test"), &pipelines.CodeBuildStepProps{
		BuildEnvironment: defaultBuidEnv(),
		PartialBuildSpec: defaultBuildRuntimes(),
		Commands: &[]*string{
			jsii.String("echo Hello World!"),
		},
		Cache: awscodebuild.Cache_Bucket(cacheBucket, nil),
	})
	wave.AddPost(buildStep)

	pipeline.BuildPipeline()

	// use V2 pipeline since it is cheaper for things that just run occasionally
	var cfnPipeline awscodepipeline.CfnPipeline
	runtime.Get(interface{}(pipeline.Pipeline().Node()), "defaultChild", interface{}(&cfnPipeline))
	cfnPipeline.AddPropertyOverride(jsii.String("PipelineType"), jsii.String("V2"))

	return stack
}

func defaultBuidEnv() *awscodebuild.BuildEnvironment {
	return &awscodebuild.BuildEnvironment{
		BuildImage:  awscodebuild.LinuxBuildImage_AMAZON_LINUX_2_ARM_3(),
		ComputeType: awscodebuild.ComputeType_SMALL,
	}
}

func defaultBuildRuntimes() awscodebuild.BuildSpec {
	return awscodebuild.BuildSpec_FromObject(&map[string]interface{}{
		"phases": map[string]interface{}{
			"install": map[string]interface{}{
				"runtime-versions": map[string]interface{}{
					"nodejs": "20",
					"golang": "1.21",
				},
			},
		},
	})
}
