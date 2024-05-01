package dynamo

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/google/uuid"
	"github.com/grantlerduck/go-was-lambda-dyanmo/lib/domain/booking"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"time"
)

var _ = Describe("Given booking event", func() {
	When("has valid state", func() {
		event := booking.Event{
			BookingId:    uuid.New().String(),
			UserId:       uuid.New().String(),
			TripFrom:     time.RFC3339,
			TripUntil:    time.RFC3339,
			HotelName:    "mockHotel",
			HotelId:      uuid.New().String(),
			FlightId:     uuid.New().String(),
			AirlineName:  "cheap-airline",
			BookingState: booking.PaymentPending,
		}
		expectedResult := Item{
			Pk:          uuid.New().String(),
			Sk:          event.BookingId,
			Gsi1Pk:      event.BookingId,
			EventId:     uuid.New().String(),
			BookingId:   event.BookingId,
			UserId:      event.UserId,
			TripFrom:    event.TripFrom,
			TripUntil:   event.TripUntil,
			HotelName:   event.HotelName,
			HotelId:     event.HotelId,
			FlightId:    event.FlightId,
			AirlineName: event.AirlineName,
			State:       booking.PaymentPending,
		}
		It("maps to dyanmo item with meaningful state", func() {
			result := new(Item).fromDomainBooking(&event)
			Expect(result.EventId).ShouldNot(BeNil())
			Expect(*result).To(MatchFields(IgnoreExtras, Fields{
				"EventId":     Ignore(),
				"Pk":          Ignore(),
				"BookingId":   Equal(expectedResult.BookingId),
				"Sk":          Equal(expectedResult.BookingId),
				"Gsi1Pk":      Equal(expectedResult.BookingId),
				"UserId":      Equal(expectedResult.UserId),
				"TripFrom":    Equal(expectedResult.TripFrom),
				"TripUntil":   Equal(expectedResult.TripUntil),
				"HotelName":   Equal(expectedResult.HotelName),
				"HotelId":     Equal(expectedResult.HotelId),
				"FlightId":    Equal(expectedResult.FlightId),
				"AirlineName": Equal(expectedResult.AirlineName),
				"State":       Equal(expectedResult.State),
			}))
		})
	})
})

var _ = Describe("Given dynamo item", func() {
	item := Item{
		Pk:          uuid.New().String(),
		Sk:          uuid.New().String(),
		Gsi1Pk:      uuid.New().String(),
		EventId:     uuid.New().String(),
		BookingId:   uuid.New().String(),
		UserId:      uuid.New().String(),
		TripFrom:    time.RFC3339,
		TripUntil:   time.RFC3339,
		HotelName:   "mockHotel",
		HotelId:     uuid.New().String(),
		FlightId:    uuid.New().String(),
		AirlineName: "cheap-airline",
		State:       booking.PaymentPending,
	}
	When("marshaled to dynamo json", func() {
		itemJson, marshallErr := attributevalue.MarshalMap(item)
		It("does not return an error", func() {
			Expect(marshallErr).ShouldNot(HaveOccurred())
		})
		It("can be marshalled back to item", func() {
			var itemUnmarshalled Item
			unmarshalErr := attributevalue.UnmarshalMap(itemJson, &itemUnmarshalled)
			Expect(unmarshalErr).ShouldNot(HaveOccurred())
			Expect(itemUnmarshalled).To(Equal(item))
		})
	})
	When("mapped to domain", func() {
		expectedDomainEvent := booking.Event{
			BookingId:    item.BookingId,
			UserId:       item.UserId,
			TripFrom:     item.TripFrom,
			TripUntil:    item.TripUntil,
			HotelName:    item.HotelName,
			HotelId:      item.HotelId,
			FlightId:     item.FlightId,
			AirlineName:  item.AirlineName,
			BookingState: item.State,
		}
		It("equals expected", func() {
			actual := item.toBookingDomain()
			Expect(actual).To(Equal(&expectedDomainEvent))
		})
	})
})
