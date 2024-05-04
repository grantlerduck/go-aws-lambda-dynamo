package booking

type Event struct {
	EventId      string `json:"event_id"`
	BookingId    string `json:"booking_id"`
	UserId       string `json:"user_id"`
	TripFrom     string `json:"from"`
	TripUntil    string `json:"until"`
	HotelName    string `json:"hotel_name"`
	HotelId      string `json:"hotel_id"`
	FlightId     string `json:"flight_id"`
	AirlineName  string `json:"airline_name"`
	BookingState State  `json:"booking_state"`
}

type EventMessage struct {
	Key     string `json:"key"`
	Tenant  string `json:"tenant"`
	Origin  string `json:"origin"`
	Payload string `json:"payload"` // base64 encoded bytestring, needs to be unmarshalled to bookingpb.Event
}
