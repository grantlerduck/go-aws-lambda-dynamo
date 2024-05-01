package main

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/grantlerduck/go-aws-lambda-dynamo/lib/domain/booking"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

var _ = Describe("Main handler function", func() {
	ctx := context.Background()
	loggerDev, _ := zap.NewDevelopment()
	When("event not nil", func() {
		var event = booking.EventMessage{
			Key:     uuid.New().String(),
			Tenant:  "eu",
			Origin:  "marketplace",
			Payload: "somepayload",
		}
		expectedKey := event.Key
		When("no error in processor", func() {
			It("handles event correctly", func() {
				processor := MockEventProcessor{}
				handler := BookingHandler{&processor, loggerDev}
				result, err := handler.HandleRequest(ctx, &event)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result).To(Equal(&expectedKey))
			})
		})
		When("processing error", func() {
			It("returns FailedToProcessError", func() {
				processor := MockEventFailProcessor{}
				handler := BookingHandler{&processor, loggerDev}
				_, err := handler.HandleRequest(ctx, &event)
				expectedProcessingError := FailedToProcessError{event}
				Expect(err).Should(MatchError(&expectedProcessingError))
				Expect(err.Error()).Should(Equal(expectedProcessingError.Error()))
			})
		})
	})
	When("event nil", func() {
		It("returns error EventNilError", func() {
			processor := MockEventProcessor{}
			handler := BookingHandler{&processor, loggerDev}
			_, err := handler.HandleRequest(ctx, nil)
			expectedNilError := EventNilError{}
			Expect(err).Should(MatchError(&expectedNilError))
			Expect(err.Error()).Should(Equal(expectedNilError.Error()))
		})
	})
})

type MockEventProcessor struct {
}

func (ep *MockEventProcessor) Process(event *booking.EventMessage) (*booking.Event, error) {
	return &booking.Event{BookingId: event.Key}, nil
}

type MockEventFailProcessor struct {
}

func (ep *MockEventFailProcessor) Process(event *booking.EventMessage) (*booking.Event, error) {
	return nil, errors.New("boom")
}
