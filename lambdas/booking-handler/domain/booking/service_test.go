package booking

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/google/uuid"
	bookingpb "github.com/grantlerduck/go-aws-lambda-dynamo/lambdas/booking-handler/proto"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

var _ = Describe("Given booking service", func() {
	logger, _ := zap.NewDevelopment()
	service := EventProcessor{&MockBookingRepo{}, logger}
	When("payload valid message", func() {
		When("and service unmarshalls and validates message", func() {
			It("repository does not return error on insert", func() {
				evntPb := bookingpb.Event{
					BookingId:       uuid.New().String(),
					UserId:          uuid.New().String(),
					FromEpochMillis: time.Now().UnixMilli(),
					ToEpochMillis:   time.Now().UnixMilli(),
					HotelName:       "mockHotel",
					HotelId:         uuid.New().String(),
					FlightId:        uuid.New().String(),
					AirlineName:     "cheap-airline",
					BookingState:    bookingpb.State_checked_in,
				}
				bytes, marshalErr := proto.Marshal(&evntPb)
				Expect(marshalErr).ShouldNot(HaveOccurred())
				eventPayload := base64.StdEncoding.EncodeToString(bytes)
				evntMsg := EventMessage{Key: uuid.New().String(), Tenant: "eu", Origin: "marketplace", Payload: eventPayload}
				// persist json for testin pruposes
				eventJson, _ := json.Marshal(evntMsg)
    			fileErr := os.WriteFile("../../../../event.json", eventJson, 0644)
				Expect(fileErr).ShouldNot(HaveOccurred())
				_, err := service.Process(&evntMsg) // the return value is mocked from the repo
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		When("and service unmarshalls and validates message", func() {
			logger, _ := zap.NewDevelopment()
			service := EventProcessor{&MockBookingFailRepo{}, logger}
			It("repository does return error on insert", func() {
				
				evntPb := bookingpb.Event{
					BookingId:       uuid.New().String(),
					UserId:          uuid.New().String(),
					FromEpochMillis: time.Now().UnixMilli(),
					ToEpochMillis:   time.Now().UnixMilli(),
					HotelName:       "mockHotel",
					HotelId:         uuid.New().String(),
					FlightId:        uuid.New().String(),
					AirlineName:     "cheap-airline",
					BookingState:    bookingpb.State_checked_in,
				}
				bytes, marshalErr := proto.Marshal(&evntPb)
				Expect(marshalErr).ShouldNot(HaveOccurred())
				eventPayload := base64.StdEncoding.EncodeToString(bytes)
				evntMsg := EventMessage{Key: uuid.New().String(), Tenant: "eu", Origin: "marketplace", Payload: eventPayload}
				_, err := service.Process(&evntMsg) // the return value is mocked from the repo
				Expect(err).Should(HaveOccurred())
			})
		})
	})
	When("payload is invalid and", func() {
		When("message no base64 byte string", func() {
			It("returns decoding error", func() {
				evntMsg := EventMessage{Key: uuid.New().String(), Tenant: "eu", Origin: "marketplace", Payload: "üöäü?12qwd"}
				_, decodeErr := service.Process(&evntMsg)
				Expect(decodeErr).Should(Equal(base64.CorruptInputError(0)))
			})
		})
		When("message no valid protobuf bytes", func() {
			It("return protobuf marshall error", func() {
				randomBase64String := "MMI7NwW5SWfdx9N/swQqHGozVHC+2XxMf/cf7h6ih8BpSIhHwgIx2WpebkPmP3QMymZUgEVksENKhqXpU32nmsisYTQidszf3fBDNC8oo7N3VN/k8F+4UfHYuxCVWoc9XLYcxfdX1A+RdrtE8rnr+ZaHMlOpv1S5/2381A=="
				evntMsg := EventMessage{Key: uuid.New().String(), Tenant: "eu", Origin: "marketplace", Payload: randomBase64String}
				_, marshalErr := service.Process(&evntMsg)
				Expect(marshalErr).Should(HaveOccurred())
			})
		})
		When("message is empty bytes", func() {
			It("return error", func() {
				randomBase64String := ""
				evntMsg := EventMessage{Key: uuid.New().String(), Tenant: "eu", Origin: "marketplace", Payload: randomBase64String}
				_, validationErr := service.Process(&evntMsg)
				Expect(validationErr).Should(HaveOccurred())
			})
		})
	})
})

type MockBookingRepo struct {
}

func (repo *MockBookingRepo) Insert(event *Event) (*Event, error) {
	return event, nil
}

func (repo *MockBookingRepo) GetByKey(bookingId string, eventId string) (*Event, error) {
	return &Event{BookingId: bookingId}, nil
}

func (repo *MockBookingRepo) GetBookingEventsByBID(bookingId string) (*[]Event, error) {
	return &[]Event{{BookingId: bookingId}}, nil
}

type MockBookingFailRepo struct {
}

func (repo *MockBookingFailRepo) Insert(event *Event) (*Event, error) {
	return nil, errors.New("boom")
}

func (repo *MockBookingFailRepo) GetByKey(bookingId string, eventId string) (*Event, error) {
	return nil, errors.New("boom")
}

func (repo *MockBookingFailRepo) GetBookingEventsByBID(bookingId string) (*[]Event, error) {
	return nil, errors.New("boom")
}
