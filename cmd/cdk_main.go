package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
	"github.com/grantlerduck/go-aws-lambda-dynamo/stacks"
)

func main() {
	defer jsii.Close()
	app := awscdk.NewApp(nil)
	stacks.NewAppStack(app, "CdkStack", &stacks.AppStackProps{
		StackProps: awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		//Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Account: jsii.String("12345678912"),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
