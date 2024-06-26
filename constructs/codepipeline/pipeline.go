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
	DEV_BRANCH  = "dev"
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
		CodePipeline: awscodepipeline.NewPipeline(scope, jsii.Sprintf("%sPipeline", id), &awscodepipeline.PipelineProps{
			PipelineType:   awscodepipeline.PipelineType_V2,
			PipelineName:   jsii.Sprintf("%s-main-pipeline", props.ServiceName),
			ArtifactBucket: props.ArtifactBucket,
			ExecutionMode:  awscodepipeline.ExecutionMode_QUEUED,
			RestartExecutionOnUpdate: jsii.Bool(true),
		}),
		Synth: codebuild.NewGoSynthStep(codebuild.GoSynthStepProps{
			ConnectionArn:   props.ConnectionArn,
			TriggerBranches: MAIN_BRANCH,
			RepoName:        props.RepoName,
			Commands:        nil,
			TriggerOnPush:   true,
		}),
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
		CodePipeline: awscodepipeline.NewPipeline(scope, jsii.Sprintf("%sPipeline", id), &awscodepipeline.PipelineProps{
			PipelineType:   awscodepipeline.PipelineType_V2,
			PipelineName:   jsii.Sprintf("%s-branch-pr-pipeline", props.ServiceName),
			ArtifactBucket: props.ArtifactBucket,
			ExecutionMode:  awscodepipeline.ExecutionMode_PARALLEL,
		}),
		Synth: codebuild.NewGoSynthStep(codebuild.GoSynthStepProps{
			ConnectionArn:   props.ConnectionArn,
			TriggerBranches: DEV_BRANCH,
			RepoName:        props.RepoName,
			Commands:        nil,
			TriggerOnPush:   true,
		}),
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
	addBranchTrigger(branchPipeline)
	return branchPipeline
}

// adds a branch based trigger to the pipeline referencing the source action
// workaround since aws cdk does not yet support push triggers 
func addBranchTrigger(pipeline pipelines.CodePipeline) {
	sourceStage := pipeline.Pipeline().Stage(jsii.String("Source"))
	actions := sourceStage.Actions()
	branches := &[]*string{
		jsii.String("dev*"),
		jsii.String("feat*"),
		jsii.String("feat/*"),
		jsii.String("chore*"),
		jsii.String("chore/*"),
		jsii.String("bug*"),
		jsii.String("bug/*"),
	}
	if actions != nil {
		acts := (*actions)
		if len(acts) > 0 {
			sourceAction := acts[0]
			includes := &map[string]*[]*string{"Includes": branches}
			branchConfig := &[]map[string]any{{"Branches": includes}}
			override := &[]any{
				&map[string]any{
					"GitConfiguration": map[string]any{
						"Push":             branchConfig,
						"SourceActionName": sourceAction.ActionProperties().ActionName,
					},
					"ProviderType": jsii.String("CodeStarSourceConnection"),
				},
			}
			var cfnPipeline awscodepipeline.CfnPipeline
			runtime.Get(interface{}(pipeline.Pipeline().Node()), "defaultChild", interface{}(&cfnPipeline))
			cfnPipeline.AddPropertyOverride(jsii.String("Triggers"), override)
		}
	}
}
