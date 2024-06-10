package stacks

import (
	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/grantlerduck/go-aws-lambda-dynamo/constructs/codepipeline"
	"github.com/grantlerduck/go-aws-lambda-dynamo/constructs/s3"
)

type PipelineStackProps struct {
	StackProps          cdk.StackProps
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
	artifactBucket := s3.NewRemoveableBucket(stack, "ArtifactBucket")
	cacheBucket := s3.NewRemoveableBucket(stack, "CacheBucket")

	stage := NewBookingEventsLambdaStage(stack, props.StackProps)
	deployAbleStages := []codepipeline.DeployableStage{
		codepipeline.DeployableStage{Stage: stage, StageOptions: nil},
	}

	// the main pipeline
	codepipeline.NewGoV2MainPipeline(stack, "MainExecutionPipeline", codepipeline.GoPipelineProps{
		ArtifactBucket: artifactBucket,
		CacheBucket:    cacheBucket,
		ConnectionArn:  githubConnectionArn,
		ServiceName:    props.ServiceName,
		RepoName:       props.RepositoryName,
		Stages:         deployAbleStages,
	})

	// the branch pipeline
	codepipeline.NewGoV2BranchPipeline(stack, "PRPipeline", codepipeline.GoPipelineProps{
		ArtifactBucket: artifactBucket,
		CacheBucket:    cacheBucket,
		ConnectionArn:  githubConnectionArn,
		ServiceName:    props.ServiceName,
		RepoName:       props.RepositoryName,
	})

	return stack
}
