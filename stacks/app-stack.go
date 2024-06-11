package stacks

import (
	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewBookingEventsLambdaStage(scope constructs.Construct, props cdk.StackProps) cdk.Stage {
	stage := cdk.NewStage(scope, jsii.String("booking-event"), nil)
	dataStackOutputs := NewDataStack(stage, "data-stack", &DataStackProps{props})
	NewAppStack(stage, "app-lambda-stack", &AppStackProps{
		StackProps:         props,
		BookingEventsTable: dataStackOutputs.BookingEventsTable,
	})
	return stage
}

type AppStackProps struct {
	StackProps         cdk.StackProps
	BookingEventsTable awsdynamodb.Table
}

func NewAppStack(scope constructs.Construct, id string, props *AppStackProps) cdk.Stack {
	var sprops cdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := cdk.NewStack(scope, &id, &sprops)

	lambda := awscdklambdagoalpha.NewGoFunction(
		stack,
		jsii.String("BookingProcessorLambda"),
		&awscdklambdagoalpha.GoFunctionProps{
			Architecture: awslambda.Architecture_ARM_64(),
			LogRetention: awslogs.RetentionDays_ONE_DAY,
			Entry:        jsii.String("lambdas/booking-handler"),
			MemorySize:   jsii.Number(128),
			Timeout:      cdk.Duration_Seconds(jsii.Number(10)),
			Environment: &map[string]*string{
				"DYNAMO_BOOKING_TABLE_NAME": props.BookingEventsTable.TableName(),
			},
		},
	)
	props.BookingEventsTable.GrantReadWriteData(lambda)
	return stack
}
