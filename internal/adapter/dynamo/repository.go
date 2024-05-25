package dynamo

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/grantlerduck/go-aws-lambda-dynamo/internal/domain/booking"
	"go.uber.org/zap"
)

type EventRepository struct {
	dynamoClient *ClientWrapper
	logger       *zap.Logger
	tableName    string
}

func (repo *EventRepository) Insert(event *booking.Event) (*booking.Event, error) {
	item := new(Item).fromDomainBooking(event)
	av, marshalErr := attributevalue.MarshalMap(item)
	if marshalErr != nil {
		repo.logger.Error("failed to marshal item",
			zap.String("bookingId", item.BookingId),
			zap.String("eventId", item.EventId),
			zap.Error(marshalErr),
		)
		return event, marshalErr
	}
	_, putItemErr := repo.dynamoClient.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(repo.tableName),
		Item:      av,
	})
	if putItemErr != nil {
		repo.logger.Error("failed to PutItem",
			zap.String("bookingId", item.BookingId),
			zap.String("eventId", item.EventId),
			zap.Error(putItemErr))
		return event, putItemErr
	}
	return event, nil
}

func (repo *EventRepository) GetByKey(bookingId string, eventId string) (*booking.Event, error) {
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(repo.tableName),
		Key: map[string]types.AttributeValue{
			ItemHasKeyAttribute:  &types.AttributeValueMemberS{Value: eventId},
			ItemSortKeyAttribute: &types.AttributeValueMemberS{Value: bookingId},
		},
	}
	output, getItemErr := repo.dynamoClient.GetItem(getItemInput)
	if getItemErr != nil {
		repo.logger.Error("failed to GetItem",
			zap.String("bookingId", bookingId),
			zap.String("eventId", eventId),
			zap.Error(getItemErr),
		)
		return nil, getItemErr
	}
	item := new(Item)
	unmarshalErr := attributevalue.UnmarshalMap(output.Item, item)
	if unmarshalErr != nil {
		repo.logger.Error("failed to unmarshal GetItem output",
			zap.String("bookingId", bookingId),
			zap.String("eventId", eventId),
			zap.Error(unmarshalErr),
		)
		return nil, unmarshalErr
	}
	return item.toBookingDomain(), nil
}

func (repo *EventRepository) GetBookingEventsByBID(bookingId string) (*[]booking.Event, error) {
	keyCondition := fmt.Sprintf("%s = :key", ItemGsi1KeyAttribute)
	queryInput := dynamodb.QueryInput{
		TableName:              aws.String(repo.tableName),
		IndexName:              aws.String(ItemGsi1IndexName),
		KeyConditionExpression: aws.String(keyCondition),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":key": &types.AttributeValueMemberS{Value: bookingId},
		},
	}
	output, queryErr := repo.dynamoClient.QueryItem(&queryInput)
	if queryErr != nil {
		repo.logger.Error("failed to Query GSI by",
			zap.String("bookingId", bookingId),
			zap.Error(queryErr),
		)
		return nil, queryErr
	}
	return repo.handleQueryAvs(output.Items, bookingId)
}

func (repo *EventRepository) handleQueryAvs(avs []map[string]types.AttributeValue, bookingId string) (*[]booking.Event, error) {
	var items []Item = (make([]Item, len(avs)))
	var itemsPointr = &items
	unmarshalErr := attributevalue.UnmarshalListOfMaps(avs, itemsPointr)
	if unmarshalErr != nil {
		repo.logger.Error("failed to unmarshal Query output",
			zap.String("bookingId", bookingId),
			zap.Error(unmarshalErr),
		)
		return nil, unmarshalErr
	}
	var events []booking.Event
	for i := range items {
		item := items[i]
		events = append(events, *item.toBookingDomain())

	}
	return &events, nil
}

func NewEventRepository(region string, tableName string, logger *zap.Logger) *EventRepository {
	client := NewClientWrapper(region)
	return &EventRepository{dynamoClient: client, logger: logger, tableName: tableName}
}

func NewLocalEventRepository(dynamoClient *dynamodb.Client, tableName string, logger *zap.Logger) *EventRepository {
	client := NewClientWrapperFromClient(dynamoClient)
	return &EventRepository{dynamoClient: client, logger: logger, tableName: tableName}
}
