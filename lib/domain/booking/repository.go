package booking

type Repository interface {
	Insert(event *Event) (*Event, error)
	GetByKey(bookingId string, eventId string) (*Event, error)
	GetBookingEventsByBID(bookingId string) (*Event, error)
}
