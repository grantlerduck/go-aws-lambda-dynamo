package main

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/grantlerduck/go-was-lambda-dyanmo/lib/domain/booking"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Main handler function", func() {
	ctx := context.Background()
	When("event not nil", func() {
		var event = booking.Event{
			BookingId:    uuid.New().String(),
			UserId:       uuid.New().String(),
			TripFrom:     time.RFC3339,
			TripUntil:    time.RFC3339,
			HotelName:    "mockHotel",
			HotelId:      uuid.New().String(),
			FlightId:     uuid.New().String(),
			AirlineName:  "cheap-airline",
			BookingState: "PAYMENT_PENDING",
		}
		expectedBookingId := event.BookingId
		When("no error in processor", func() {
			It("handles event correctly", func() {
				processor := MockEventProcessor{}
				handler := BookingHandler{&processor, logger}
				result, err := handler.HandleRequest(ctx, &event)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result).To(Equal(&expectedBookingId))
			})
		})
		When("processing error", func() {
			It("returns FailedToProcessError", func() {
				processor := MockEventFailProcessor{}
				handler := BookingHandler{&processor, logger}
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
			handler := BookingHandler{&processor, logger}
			_, err := handler.HandleRequest(ctx, nil)
			expectedNilError := EventNilError{}
			Expect(err).Should(MatchError(&expectedNilError))
			Expect(err.Error()).Should(Equal(expectedNilError.Error()))
		})
	})
})

type MockEventProcessor struct {
}

func (ep *MockEventProcessor) Process(event *booking.Event) (*booking.Event, error) {
	return event, nil
}

type MockEventFailProcessor struct {
}

func (ep *MockEventFailProcessor) Process(event *booking.Event) (*booking.Event, error) {
	return nil, errors.New("boom")
}
