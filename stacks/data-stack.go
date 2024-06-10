package stacks

import (
	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type DataStackProps struct {
	StackProps cdk.StackProps
}

type DataStackOutPuts struct {
	Stack cdk.Stack
	BookingEventsTabel awsdynamodb.Table
}

func NewDataStack(scope constructs.Construct, id string, props *DataStackProps) DataStackOutPuts {
	var sprops cdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := cdk.NewStack(scope, &id, &sprops)
	bookingEventsTable := awsdynamodb.NewTable(stack, jsii.String("BookingEventsTable"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{Name: jsii.String("pk"), Type: awsdynamodb.AttributeType_STRING},
		SortKey: &awsdynamodb.Attribute{Name: jsii.String("sk"), Type: awsdynamodb.AttributeType_STRING},
		BillingMode: awsdynamodb.BillingMode_PROVISIONED,
		ReadCapacity: jsii.Number(5), // default value
		WriteCapacity: jsii.Number(5), // default value
		RemovalPolicy: cdk.RemovalPolicy_DESTROY, // don't do that in production
	})
	bookingEventsTable.AddGlobalSecondaryIndex(&awsdynamodb.GlobalSecondaryIndexProps{
		IndexName: jsii.String("GSI1"),
		PartitionKey: &awsdynamodb.Attribute{Name: jsii.String("gsi1_pk"), Type: awsdynamodb.AttributeType_STRING},
		ProjectionType: awsdynamodb.ProjectionType_ALL,
		ReadCapacity: jsii.Number(5), // default value
		WriteCapacity: jsii.Number(5), // default value
	})

	return DataStackOutPuts{Stack: stack, BookingEventsTabel: bookingEventsTable}
}
