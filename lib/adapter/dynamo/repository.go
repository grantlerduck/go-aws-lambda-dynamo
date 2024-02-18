package dynamo

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/grantlerduck/go-was-lambda-dyanmo/lib/domain/booking"
	"go.uber.org/zap"
)

type EventRepository struct {
	dynamoClient *dynamodb.Client
	logger       *zap.Logger
	tableName    string
}

func (repo *EventRepository) Insert(event *booking.Event) (*booking.Event, error) {
	av, marshalErr := attributevalue.MarshalMap(FromDomainBooking(event))
	if marshalErr != nil {
		repo.logger.Error("Failed to marshal item",
			zap.Any("item", event),
			zap.Error(marshalErr),
		)
		return event, marshalErr
	}
	_, putItemErr := repo.dynamoClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(repo.tableName),
		Item:      av,
	})
	if putItemErr != nil {
		repo.logger.Error("Failed to put item", zap.Error(putItemErr))
		return event, putItemErr
	}
	return event, nil
}

func NewEventRepository(region string, tableName string, logger *zap.Logger) *EventRepository {
	awsConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}
	dynamoClient := dynamodb.NewFromConfig(awsConfig, func(opt *dynamodb.Options) {
		opt.Region = region
	})
	return &EventRepository{dynamoClient: dynamoClient, logger: logger, tableName: tableName}
}
