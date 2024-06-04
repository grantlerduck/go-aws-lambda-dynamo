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
)


type PipelineStackProps struct {
	StackProps cdk.StackProps
	PipelineName string
	RepositoryName string
	ServiceName string
}

func NewPipelineStack(scope constructs.Construct, id string, props *PipelineStackProps) cdk.Stack {
	var sprops cdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := cdk.NewStack(scope, &id, &sprops)

	githubConnectionArn := cdk.Fn_ImportValue(jsii.String("account-setup-github-connection-github-connection-arn"))

	artifactBucket := awss3.NewBucket(stack, jsii.String("ArtifactBucket"), &awss3.BucketProps{
		EnforceSSL: jsii.Bool(true),
		LifecycleRules: &[]*awss3.LifecycleRule{
			{
				Enabled:    jsii.Bool(true),
				Expiration: awscdk.Duration_Days(jsii.Number(7)),
			},
		},
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})
	// cacheBucket := awss3.NewBucket(stack, jsii.String("CacheBucket"), &awss3.BucketProps{
	// 	EnforceSSL: jsii.Bool(true),
	// 	LifecycleRules: &[]*awss3.LifecycleRule{
	// 		{
	// 			Enabled:    jsii.Bool(true),
	// 			Expiration: awscdk.Duration_Days(jsii.Number(7)),
	// 		},
	// 	},
	// 	RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	// })
	pipeline := pipelines.NewCodePipeline(stack, jsii.String("MainPipeline"), &pipelines.CodePipelineProps{
		SelfMutation: jsii.Bool(true),
		SynthCodeBuildDefaults: &pipelines.CodeBuildOptions{
			BuildEnvironment: &awscodebuild.BuildEnvironment{
				BuildImage: awscodebuild.LinuxArmLambdaBuildImage_AMAZON_LINUX_2_GO_1_21(),
				ComputeType: awscodebuild.ComputeType_LAMBDA_10GB,
			},
		},
		Synth: pipelines.NewShellStep(jsii.String("Synth"), &pipelines.ShellStepProps{
			Input: pipelines.CodePipelineSource_Connection(jsii.String(props.RepositoryName), jsii.String("main"), &pipelines.ConnectionSourceOptions{
				ConnectionArn: githubConnectionArn,
				TriggerOnPush: jsii.Bool(true),
			}),
			Commands: &[]*string{
				jsii.String("go mod download"),
				jsii.String("go mod tidy"),
				jsii.String("go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"),
				jsii.String("${GOPATH}/bin/golangci-lint run ./..."),
				jsii.String("npx cdk synth"),
			},
		}),
		ArtifactBucket: artifactBucket,
		PipelineName: jsii.String(props.PipelineName),
	})
	pipeline.BuildPipeline()
	var cfnPipeline awscodepipeline.CfnPipeline
	jsii.Get(pipeline.Pipeline().Node(), "defaultChild", &cfnPipeline)
	cfnPipeline.AddPropertyOverride(jsii.String("PipelineType"), jsii.String("V2"))
	return stack
}