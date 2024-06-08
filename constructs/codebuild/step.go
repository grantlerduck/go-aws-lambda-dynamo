package codebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
	"github.com/aws/jsii-runtime-go"
)

type GoStepProps struct {
	CacheBucket awss3.IBucket
	CachePrefix string
	Commands    *[]*string
}

type GoSynthStepProps struct {
	ConnectionArn   *string
	RepoName        string
	TriggerBranches string
	Commands        *[]*string
}

func NewGoLintStep(props GoStepProps) pipelines.CodeBuildStep {
	var cmds *[]*string = props.Commands
	if cmds == nil {
		cmds = &[]*string{
			jsii.String("export PATH=\"$GOPATH/bin:$PATH\""),
			jsii.String("go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"),
			jsii.String("golangci-lint run ./..."),
		}
	}
	return pipelines.NewCodeBuildStep(jsii.String("TestLint"), &pipelines.CodeBuildStepProps{
		BuildEnvironment: NewUbuntuBuildEnv(jsii.String("MEDIUM")),
		PartialBuildSpec: NewDefaultBuildRuntimes(),
		Commands:         cmds,
		Cache: awscodebuild.Cache_Bucket(props.CacheBucket, &awscodebuild.BucketCacheOptions{
			Prefix: jsii.Sprintf("%s/lint", props.CachePrefix),
		}),
	})
}

func NewGoTestReportStep(props GoStepProps) pipelines.CodeBuildStep {
	var cmds *[]*string = props.Commands
	if cmds == nil {
		cmds = &[]*string{
			jsii.String("go version"),
			jsii.String("export PATH=\"$GOPATH/bin:$PATH\""),
			jsii.String("go install gotest.tools/gotestsum@latest"),
			jsii.String("go install github.com/onsi/ginkgo/v2/ginkgo@latest"),
			jsii.String("go install github.com/t-yuki/gocover-cobertura@latest"),
			jsii.String("go mod download"),
			jsii.String("go mod tidy"),
			jsii.String("gotestsum --junitfile unit-tests.xml -- -coverprofile=cover.out -covermode count ./..."),
			jsii.String("go tool cover -html=cover.out -o coverage.html"),
			jsii.String("gocover-cobertura < cover.out > coverage.xml"),
		}
	}
	return pipelines.NewCodeBuildStep(jsii.String("TestReport"), &pipelines.CodeBuildStepProps{
		BuildEnvironment: NewDefaultBuildEnv(jsii.String("MEDIUM")),
		PartialBuildSpec: NewDefaultBuildRuntimes(),
		Commands:         cmds,
		Cache: awscodebuild.Cache_Bucket(props.CacheBucket, &awscodebuild.BucketCacheOptions{
			Prefix: jsii.Sprintf("%s/test-build-report", props.CachePrefix),
		}),
	})
}

func NewGoSynthStep(props GoSynthStepProps) pipelines.ShellStep {
	var cmds *[]*string = props.Commands
	if cmds == nil {
		cmds = &[]*string{
			jsii.String("go version"),
			jsii.String("npm install -g aws-cdk@latest"),
			jsii.String("go mod download"),
			jsii.String("go mod tidy"),
			jsii.String("npx cdk synth >> /dev/null"),
		}
	}
	return pipelines.NewShellStep(jsii.String("Synth"), &pipelines.ShellStepProps{
		Input: pipelines.CodePipelineSource_Connection(jsii.String(props.RepoName), jsii.String(props.TriggerBranches), &pipelines.ConnectionSourceOptions{
			ConnectionArn: props.ConnectionArn,
			TriggerOnPush: jsii.Bool(true),
		}),
		Commands: cmds,
	})
}
