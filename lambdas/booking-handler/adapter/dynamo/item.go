package dynamo

import (
	"github.com/grantlerduck/go-aws-lambda-dynamo/lambdas/booking-handler/domain/booking"
)

const (
	ItemHasKeyAttribute  string = "pk"
	ItemSortKeyAttribute string = "sk"
	ItemGsi1KeyAttribute string = "gsi1_pk"
	ItemGsi1IndexName    string = "GSI1"
)

type Item struct {
	Pk          string        `json:"pk" dynamodbav:"pk"`
	Sk          string        `json:"sk" dynamodbav:"sk"`
	Gsi1Pk      string        `json:"gsi1_pk" dynamodbav:"gsi1_pk"`
	EventId     string        `json:"event_id" dynamodbav:"event_id"`
	BookingId   string        `json:"booking_id" dynamodbav:"booking_id"`
	UserId      string        `json:"user_id" dynamodbav:"user_id"`
	TripFrom    string        `json:"from" dynamodbav:"from"`
	TripUntil   string        `json:"until" dynamodbav:"until"`
	HotelName   string        `json:"hotel_name" dynamodbav:"hotel_name"`
	HotelId     string        `json:"hotel_id" dynamodbav:"hotel_id"`
	FlightId    string        `json:"flight_id" dynamodbav:"flight_id"`
	AirlineName string        `json:"airline_name" dynamodbav:"airline_name"`
	State       booking.State `json:"state" dynamodbav:"state"`
}

func (item *Item) fromDomainBooking(domain *booking.Event) *Item {
	item.Pk = domain.EventId
	item.Sk = domain.BookingId
	item.Gsi1Pk = domain.BookingId
	item.EventId = domain.EventId
	item.BookingId = domain.BookingId
	item.UserId = domain.UserId
	item.TripFrom = domain.TripFrom
	item.TripUntil = domain.TripUntil
	item.HotelName = domain.HotelName
	item.HotelId = domain.HotelId
	item.FlightId = domain.FlightId
	item.AirlineName = domain.AirlineName
	item.State = domain.BookingState
	return item
}

func (item *Item) toBookingDomain() *booking.Event {
	event := new(booking.Event)
	event.EventId = item.EventId
	event.BookingId = item.BookingId
	event.UserId = item.UserId
	event.TripFrom = item.TripFrom
	event.TripUntil = item.TripUntil
	event.HotelName = item.HotelName
	event.HotelId = item.HotelId
	event.FlightId = item.FlightId
	event.AirlineName = item.AirlineName
	event.BookingState = item.State
	return event
}
