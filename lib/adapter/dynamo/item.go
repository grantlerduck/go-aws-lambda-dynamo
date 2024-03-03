package dynamo

import (
	"github.com/google/uuid"
	"github.com/grantlerduck/go-was-lambda-dyanmo/lib/domain/booking"
	"strings"
)

type State string

const (
	Unconfirmed     State = "booking-unconfirmed"
	Confirmed       State = "booking-confirmed"
	PaymentReceived State = "booking-fee-payed"
	PaymentPending  State = "booking-fee-pending"
	Planned         State = "booking-planned"
	Canceled        State = "booking-canceled"
	CheckedIn       State = "checked-in"
	CheckedOut      State = "checked-out"
	ReviewPending   State = "review-pending"
	Reviewed        State = "customer-reviewed"
	Unknown         State = "unknown"
)

func (s State) String() string {
	return string(s)
}

var states = map[State]struct{}{
	Unconfirmed:     {},
	Confirmed:       {},
	PaymentReceived: {},
	PaymentPending:  {},
	Planned:         {},
	Canceled:        {},
	CheckedIn:       {},
	CheckedOut:      {},
	Reviewed:        {},
	Unknown:         {},
}

func getState(str string) State {
	state := State(strings.ToLower(str))
	_, ok := states[state]
	if !ok {
		return Unknown
	}
	return state
}

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
