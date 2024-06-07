package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
	"github.com/grantlerduck/go-aws-lambda-dynamo/stacks"
)

func main() {
	defer jsii.Close()
	app := awscdk.NewApp(nil)
	service := "go-aws-lambda-dynamo"
	group := "grantlerduck"

	stacks.NewPipelineStack(app, fmt.Sprintf("%s-pipeline-stack", service), &stacks.PipelineStackProps{
		StackProps: awscdk.StackProps{
			Env: env(),
		},
		PipelineName: fmt.Sprintf("%s-pipeline", service),
		RepositoryName: fmt.Sprintf("%s/%s", group, service),
		ServiceName: service,
		ConnectionArnImport: "account-setup-github-connection-github-connection-arn",
	})
	print("Hello World!")

	app.Synth(nil)
}

func env() *awscdk.Environment {
	// use your prefered region and your accountId hardocded if you know where to deploy to
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
