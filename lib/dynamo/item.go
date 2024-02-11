package dynamo

type Item struct {
	EventId     string `json:"event_id"`
	BookingId   string `json:"booking_id"`
	UserId      string `json:"user_id"`
	TripFrom    string `json:"from"`
	TripUntil   string `json:"until"`
	HotelName   string `json:"hotel_name"`
	HotelId     string `json:"hotel_id"`
	FlightId    string `json:"flight_id"`
	AirlineName string `json:"airline_name"`
	State       State  `json:"state"`
}

type State string

const (
	Unconfirmed     State = "booking-unconfirmed"
	Confirmed       State = "booking-confirmed"
	PaymentReceived State = "booking-payed"
	Planned         State = "booking-planned"
	Canceled        State = "booking-canceled"
	CheckedIn       State = "checked-in"
	CheckedOut      State = "checked-out"
	ReviewPending   State = "review-pending"
	Reviewed        State = "customer-reviewed"
	Unknown         State = "unknown"
)
