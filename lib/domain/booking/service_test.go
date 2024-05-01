package booking

import (
	"encoding/base64"
	"errors"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

var _ = Describe("Given booking service", func() {
	When("payload valid message", func() {
		When("repository does not returns error on insert", func() {
			println("TODO")
		})
		When("repository returns error on insert", func() {
			println("TODO")

		})
	})
	When("payload is invalid", func() {
		logger, _ := zap.NewDevelopment()
		service := EventProcessor{&MockBookingRepo{}, logger}
		When("and no bas64 byte string", func() {
			It("returns decoding error", func() {
				evntMsg := EventMessage{Key: uuid.New().String(), Tenant: "eu", Origin: "marketplace", Payload: "üöäü?12qwd"}
				_, decodeErr := service.Process(&evntMsg)
				Expect(decodeErr).Should(Equal(base64.CorruptInputError(0)))
			})
		})
		When("and no valid protobuf bytes", func() {
			It("return protobuf marshall error", func() {
				randomBase64String := "MMI7NwW5SWfdx9N/swQqHGozVHC+2XxMf/cf7h6ih8BpSIhHwgIx2WpebkPmP3QMymZUgEVksENKhqXpU32nmsisYTQidszf3fBDNC8oo7N3VN/k8F+4UfHYuxCVWoc9XLYcxfdX1A+RdrtE8rnr+ZaHMlOpv1S5/2381A=="
				evntMsg := EventMessage{Key: uuid.New().String(), Tenant: "eu", Origin: "marketplace", Payload: randomBase64String}
				_, marshalErr := service.Process(&evntMsg)
				Expect(marshalErr).Should(HaveOccurred())
			})
		})
		When("and empty bytes", func() {
			It("return init error", func() {
				randomBase64String := ""
				evntMsg := EventMessage{Key: uuid.New().String(), Tenant: "eu", Origin: "marketplace", Payload: randomBase64String}
				_, validationErr := service.Process(&evntMsg)
				// TODO assert for explicit err
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
