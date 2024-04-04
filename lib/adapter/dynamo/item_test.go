package dynamo

import (
	"github.com/google/uuid"
	"github.com/grantlerduck/go-was-lambda-dyanmo/lib/domain/booking"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"time"
)

var _ = Describe("Booking state string", func() {
	When("is mappable to Unconfirmed", func() {
		unconfirmed := "BOOKING-UNCONFIRMED"
		expectedResult := Unconfirmed
		It("returns correctly", func() {
			result := getState(unconfirmed)
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to Confirmed", func() {
		confirmed := "BOOKING-Confirmed"
		expectedResult := Confirmed
		It("returns correctly", func() {
			result := getState(confirmed)
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to PaymentReceived", func() {
		payed := "Booking-FeE-PayeD"
		expectedResult := PaymentReceived
		It("returns correctly", func() {
			result := getState(payed)
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to PaymentPending", func() {
		pending := "booking-fee-pending"
		expectedResult := PaymentPending
		It("returns correctly", func() {
			result := getState(pending)
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to Planned", func() {
		planned := "booking-planned"
		expectedResult := Planned
		It("returns correctly", func() {
			result := getState(planned)
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to Canceled", func() {
		canceled := "booking-canceled"
		expectedResult := Canceled
		It("returns correctly", func() {
			result := getState(canceled)
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to CheckedIn", func() {
		checkedIn := "checked-in"
		expectedResult := CheckedIn
		It("returns correctly", func() {
			result := getState(checkedIn)
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to CheckedIn", func() {
		checkedOut := "checked-out"
		expectedResult := CheckedOut
		It("returns correctly", func() {
			result := getState(checkedOut)
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to ReviewPending", func() {
		reviewPending := "review-pending"
		expectedResult := ReviewPending
		It("returns correctly", func() {
			result := getState(reviewPending)
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to Reviewed", func() {
		reviewed := "customer-reviewed"
		expectedResult := Reviewed
		It("returns correctly", func() {
			result := getState(reviewed)
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable not mappable to meaningful state", func() {
		someString := "qwelfnqälwrjf"
		expectedResult := Unknown
		It("returns Unknown", func() {
			result := getState(someString)
			Expect(result).To(Equal(expectedResult))
		})
	})
})

var _ = Describe("Given booking event", func() {
	When("is has valid state", func() {
		event := booking.Event{
			BookingId:    uuid.New().String(),
			UserId:       uuid.New().String(),
			TripFrom:     time.RFC3339,
			TripUntil:    time.RFC3339,
			HotelName:    "mockHotel",
			HotelId:      uuid.New().String(),
			FlightId:     uuid.New().String(),
			AirlineName:  "cheap-airline",
			BookingState: "booking-fee-pending",
		}
		expectedResult := Item{
			EventId:     uuid.New().String(),
			BookingId:   event.BookingId,
			UserId:      event.UserId,
			TripFrom:    event.TripFrom,
			TripUntil:   event.TripUntil,
			HotelName:   event.HotelName,
			HotelId:     event.HotelId,
			FlightId:    event.FlightId,
			AirlineName: event.AirlineName,
			State:       PaymentPending,
		}
		It("maps to booking event with meaningful state", func() {
			result := FromDomainBooking(&event)
			Expect(result.EventId).ShouldNot(BeNil())
			Expect(*result).To(MatchFields(IgnoreExtras, Fields{
				"EventId":     Ignore(),
				"BookingId":   Equal(expectedResult.BookingId),
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
