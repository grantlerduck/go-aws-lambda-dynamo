package dynamo

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
	"github.com/grantlerduck/go-aws-lambda-dynamo/internal/domain/booking"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/localstack"
	"go.uber.org/zap"
)

const localstackImage string = "localstack/localstack:3.4.0"
const portProtocol nat.Port = "4566/tcp"
const testTable string = "booking-events-table"
const regionLocal string = "eu-west-1"

var _ = Describe("Given event repository", Ordered, func() {
	logger, _ := zap.NewDevelopment()
	var localStackContainer *localstack.LocalStackContainer
	var port nat.Port
	var dynamoClientLocal *dynamodb.Client
	var repo *EventRepository
	BeforeAll(func() {
		ctx := context.Background()
		var err error
		localStackContainer, err = localstack.RunContainer(
			ctx,
			testcontainers.WithImage(localstackImage),
		)
		Expect(err).ShouldNot(HaveOccurred())
		port, err = localStackContainer.MappedPort(context.Background(), portProtocol)
		Expect(err).ShouldNot(HaveOccurred())
		dynamoClientLocal, err = createDynamoLocalClient(port)
		Expect(err).ShouldNot(HaveOccurred())
		err = createEventTable(dynamoClientLocal, testTable)
		Expect(err).ShouldNot(HaveOccurred())
		repo = NewLocalEventRepository(dynamoClientLocal, testTable, logger)
	})
	event := booking.Event{
		EventId:      uuid.New().String(),
		BookingId:    uuid.New().String(),
		UserId:       uuid.New().String(),
		TripFrom:     "2006-01-02T15:04:05.999999999Z-0700",
		TripUntil:    "2006-01-02T15:04:05.999999999Z-0700",
		HotelName:    "mockHotel",
		HotelId:      uuid.New().String(),
		FlightId:     uuid.New().String(),
		AirlineName:  "cheap-airline",
		BookingState: booking.PaymentPending,
	}
	// this suit is ordered be carful!
	It("event is inserted into table", func() {
		actual, err := repo.Insert(&event)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(actual).Should(Equal(&event))
	})
	It("event is queried by gsi", func() {
		results, err := repo.GetBookingEventsByBID(event.BookingId)
		Expect(err).ShouldNot(HaveOccurred())
		resulsDeref := (*results)
		Expect(len(resulsDeref)).Should(Equal(1))
		actual := resulsDeref[0]
		Expect(actual).Should(Equal(event))
	})
	It("event is get by id successful", func() {
		actual, err := repo.GetByKey(event.BookingId, event.EventId)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(actual).Should(Equal(&event))
	})
	AfterAll(func() {
		if err := localStackContainer.Terminate(context.Background()); err != nil {
			Expect(err).ShouldNot(HaveOccurred())
		}
	})
})

func createEventTable(dynamoCli *dynamodb.Client, tableName string) error {
	input := dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		KeySchema: []types.KeySchemaElement{
			{AttributeName: aws.String(ItemHasKeyAttribute), KeyType: types.KeyTypeHash},
			{AttributeName: aws.String(ItemSortKeyAttribute), KeyType: types.KeyTypeRange},
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{AttributeName: aws.String(ItemHasKeyAttribute), AttributeType: types.ScalarAttributeTypeS},
			{AttributeName: aws.String(ItemSortKeyAttribute), AttributeType: types.ScalarAttributeTypeS},
			{AttributeName: aws.String(ItemGsi1KeyAttribute), AttributeType: types.ScalarAttributeTypeS},
		},
		GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
			{
				IndexName: aws.String(ItemGsi1IndexName),
				KeySchema: []types.KeySchemaElement{
					{AttributeName: aws.String(ItemGsi1KeyAttribute), KeyType: types.KeyTypeHash},
				},
				Projection: &types.Projection{ProjectionType: types.ProjectionTypeAll},
				ProvisionedThroughput: &types.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(10),
					WriteCapacityUnits: aws.Int64(10),
				},
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
	}
	_, err := dynamoCli.CreateTable(context.Background(), &input) // this is async operation maybe add a wait later?!
	Expect(err).ShouldNot(HaveOccurred())
	return nil
}

func createDynamoLocalClient(port nat.Port) (*dynamodb.Client, error) {
	var c, err = config.LoadDefaultConfig(context.Background(), config.WithRegion(regionLocal), config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
		Value: aws.Credentials{
			AccessKeyID: "test", SecretAccessKey: "test", SessionToken: "",
			Source: "Mock credentials used above for local instance",
		},
	}))
	Expect(err).ShouldNot(HaveOccurred())
	var t = dynamodb.NewFromConfig(c, func(options *dynamodb.Options) {
		options.BaseEndpoint = aws.String(fmt.Sprintf("http://localhost:%s", port.Port()))
	})
	return t, nil
}
