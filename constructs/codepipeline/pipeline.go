package codepipeline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/aws/jsii-runtime-go/runtime"
	"github.com/grantlerduck/go-aws-lambda-dynamo/constructs/codebuild"
)

type GoPipelineProps struct {
	ArtifactBucket awss3.IBucket
	CacheBucket    awss3.IBucket
	ConnectionArn  *string
	ServiceName    string
	RepoName       string
	Stages         []DeployableStage
}

type DeployableStage struct {
	Stage        awscdk.Stage
	StageOptions *pipelines.AddStageOpts
}

const (
	MAIN_BRANCH = "main"
	BRANCHES    = "dev/*,feat/*,chore/*"
	BRANCH      = "branch"
)

func NewGoV2MainPipeline(scope constructs.Construct, id string, props GoPipelineProps) pipelines.CodePipeline {
	mainPipeline := pipelines.NewCodePipeline(scope, jsii.String(id), &pipelines.CodePipelineProps{
		SelfMutation: jsii.Bool(true),
		SynthCodeBuildDefaults: &pipelines.CodeBuildOptions{
			BuildEnvironment: codebuild.NewDefaultBuildEnv(nil, jsii.Bool(false)),
			PartialBuildSpec: codebuild.NewDefaultBuildRuntimes(),
			Cache: awscodebuild.Cache_Bucket(props.CacheBucket, &awscodebuild.BucketCacheOptions{
				Prefix: jsii.Sprintf("%s/synth", MAIN_BRANCH),
			}),
		},
		Synth: codebuild.NewGoSynthStep(codebuild.GoSynthStepProps{
			ConnectionArn:   props.ConnectionArn,
			TriggerBranches: MAIN_BRANCH,
			RepoName:        props.RepoName,
			Commands:        nil,
		}),
		ArtifactBucket: props.ArtifactBucket,
		PipelineName:   jsii.Sprintf("%s-pipeline", props.ServiceName),
	})

	testWave := mainPipeline.AddWave(jsii.String("Test"), nil)
	lintStep := codebuild.NewGoLintStep(codebuild.GoStepProps{
		CacheBucket: props.CacheBucket,
		CachePrefix: MAIN_BRANCH,
		Commands:    nil,
	})
	testStep := codebuild.NewGoTestReportStep(codebuild.GoStepProps{
		CacheBucket: props.CacheBucket,
		CachePrefix: MAIN_BRANCH,
		Commands:    nil,
	})
	testWave.AddPost(lintStep, testStep)

	if len(props.Stages) > 0 {
		deployWave := mainPipeline.AddWave(jsii.String("Deploy"), nil)
		for _, stage := range props.Stages {
			deployWave.AddStage(stage.Stage, stage.StageOptions)

		}
	}

	mainPipeline.BuildPipeline()
	// use V2 pipeline since it is cheaper for things that just run occasionally
	toV2Pipeline(mainPipeline)
	return mainPipeline
}

func NewGoV2BranchPipeline(scope constructs.Construct, id string, props GoPipelineProps) pipelines.CodePipeline {
	branchPipeline := pipelines.NewCodePipeline(scope, jsii.String(id), &pipelines.CodePipelineProps{
		SelfMutation: jsii.Bool(false),
		SynthCodeBuildDefaults: &pipelines.CodeBuildOptions{
			BuildEnvironment: codebuild.NewDefaultBuildEnv(nil, jsii.Bool(false)),
			PartialBuildSpec: codebuild.NewDefaultBuildRuntimes(),
			Cache: awscodebuild.Cache_Bucket(props.CacheBucket, &awscodebuild.BucketCacheOptions{
				Prefix: jsii.Sprintf("%s/synth", BRANCH),
			}),
		},
		Synth: codebuild.NewGoSynthStep(codebuild.GoSynthStepProps{
			ConnectionArn:   props.ConnectionArn,
			TriggerBranches: BRANCHES,
			RepoName:        props.RepoName,
			Commands:        nil,
		}),
		ArtifactBucket: props.ArtifactBucket,
		PipelineName:   jsii.Sprintf("%s-branch-pipeline", props.ServiceName),
	})

	testWave := branchPipeline.AddWave(jsii.String("Test"), nil)
	lintStep := codebuild.NewGoLintStep(codebuild.GoStepProps{
		CacheBucket: props.CacheBucket,
		CachePrefix: BRANCH,
		Commands:    nil,
	})
	testStep := codebuild.NewGoTestReportStep(codebuild.GoStepProps{
		CacheBucket: props.CacheBucket,
		CachePrefix: BRANCH,
		Commands:    nil,
	})
	testWave.AddPost(lintStep, testStep)
	branchPipeline.BuildPipeline()
	// use V2 pipeline since it is cheaper for things that just run occasionally
	toV2Pipeline(branchPipeline)
	toParallelExecution(branchPipeline)
	return branchPipeline
}

func toV2Pipeline(pipeline pipelines.CodePipeline) {
	var cfnPipeline awscodepipeline.CfnPipeline
	runtime.Get(interface{}(pipeline.Pipeline().Node()), "defaultChild", interface{}(&cfnPipeline))
	cfnPipeline.AddPropertyOverride(jsii.String("PipelineType"), jsii.String("V2"))
}

func toParallelExecution(pipeline pipelines.CodePipeline) {
	var cfnPipeline awscodepipeline.CfnPipeline
	runtime.Get(interface{}(pipeline.Pipeline().Node()), "defaultChild", interface{}(&cfnPipeline))
	cfnPipeline.AddPropertyOverride(jsii.String("ExecutionMode"), jsii.String("PARALLEL"))
}
