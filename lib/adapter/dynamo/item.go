package dynamo

import (
	"github.com/google/uuid"
	"github.com/grantlerduck/go-was-lambda-dyanmo/lib/domain/booking"
)

const (
	HasKeyAttribute  string = "event_id"
	SortKeyAttribute string = "booking_id"
	Gsi1KeyAttribute string = "booking_id"
)

type Item struct {
	EventId     string `json:"event_id" dynamodbav:"event_id"`
	BookingId   string `json:"booking_id" dynamodbav:"booking_id"`
	UserId      string `json:"user_id" dynamodbav:"user_id"`
	TripFrom    string `json:"from" dynamodbav:"from"`
	TripUntil   string `json:"until" dynamodbav:"until"`
	HotelName   string `json:"hotel_name" dynamodbav:"hotel_name"`
	HotelId     string `json:"hotel_id" dynamodbav:"hotel_id"`
	FlightId    string `json:"flight_id" dynamodbav:"flight_id"`
	AirlineName string `json:"airline_name" dynamodbav:"airline_name"`
	State       State  `json:"state" dynamodbav:"state"`
}

func FromDomainBooking(domain *booking.Event) *Item {
	item := new(Item)
	item.EventId = uuid.New().String()
	item.BookingId = domain.BookingId
	item.UserId = domain.UserId
	item.TripFrom = domain.TripFrom
	item.TripUntil = domain.TripUntil
	item.HotelName = domain.HotelName
	item.HotelId = domain.HotelId
	item.FlightId = domain.FlightId
	item.AirlineName = domain.AirlineName
	item.State = getState(domain.BookingState)
	return item
}

func ToBookingDomain(item Item) *booking.Event {
	event := new(booking.Event)
	event.BookingId = item.BookingId
	event.UserId = item.UserId
	event.TripFrom = item.TripFrom
	event.TripUntil = item.TripUntil
	event.HotelName = item.HotelName
	event.HotelId = item.HotelId
	event.FlightId = item.FlightId
	event.AirlineName = item.AirlineName
	event.BookingState = item.State.String()
	return event
}
