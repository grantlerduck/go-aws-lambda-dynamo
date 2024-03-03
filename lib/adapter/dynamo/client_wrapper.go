package dynamo

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DbClient interface {
	PutItem(params *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
}

type ClientWrapper struct {
	dynamoClient *dynamodb.Client
}

func (client *ClientWrapper) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	output, err := client.dynamoClient.PutItem(context.TODO(), input)
	return output, err
}

func NewClientWrapper(region string) *ClientWrapper {
	awsConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}
	dynamoClient := dynamodb.NewFromConfig(awsConfig, func(opt *dynamodb.Options) {
		opt.Region = region
	})
	return &ClientWrapper{dynamoClient: dynamoClient}
}
