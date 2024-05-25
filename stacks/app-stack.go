package stacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type AppStackProps struct {
	StackProps cdk.StackProps
}

func NewAppStack(scope constructs.Construct, id string, props *AppStackProps) cdk.Stack {
	var sprops cdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := cdk.NewStack(scope, &id, &sprops)

	awscdklambdagoalpha.NewGoFunction(
		stack,
		jsii.String("BookingProcessorLambda"),
		&awscdklambdagoalpha.GoFunctionProps{
			Architecture: awslambda.Architecture_ARM_64(),
			LogRetention: awslogs.RetentionDays_ONE_DAY,
			Entry: jsii.String("lambdas/booking-handler"),
			MemorySize: jsii.Number(248),
			Timeout: cdk.Duration_Seconds(jsii.Number(10)),
		},
	)

	return stack
}