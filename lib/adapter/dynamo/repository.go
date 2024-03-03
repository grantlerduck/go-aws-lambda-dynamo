package dynamo

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/grantlerduck/go-was-lambda-dyanmo/lib/domain/booking"
	"go.uber.org/zap"
)

type EventRepository struct {
	dynamoClient *ClientWrapper
	logger       *zap.Logger
	tableName    string
}

func (repo *EventRepository) Insert(event *booking.Event) (*booking.Event, error) {
	av, marshalErr := attributevalue.MarshalMap(FromDomainBooking(event))
	if marshalErr != nil {
		repo.logger.Error("failed to marshal item",
			zap.Any("item", event),
			zap.Error(marshalErr),
		)
		return event, marshalErr
	}
	_, putItemErr := repo.dynamoClient.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(repo.tableName),
		Item:      av,
	})
	if putItemErr != nil {
		repo.logger.Error("failed to put item", zap.Error(putItemErr))
		return event, putItemErr
	}
	return event, nil
}

func NewEventRepository(region string, tableName string, logger *zap.Logger) *EventRepository {
	client := NewClientWrapper(region)
	return &EventRepository{dynamoClient: client, logger: logger, tableName: tableName}
}
