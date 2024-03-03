package booking

type Event struct {
	BookingId    string `json:"booking_id"`
	UserId       string `json:"user_id"`
	TripFrom     string `json:"from"`
	TripUntil    string `json:"until"`
	HotelName    string `json:"hotel_name"`
	HotelId      string `json:"hotel_id"`
	FlightId     string `json:"flight_id"`
	AirlineName  string `json:"airline_name"`
	BookingState string `json:"booking_state"`
}
