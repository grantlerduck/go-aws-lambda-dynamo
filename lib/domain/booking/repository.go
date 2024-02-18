package booking

type Repository interface {
	Insert(event *Event) (*Event, error)
}
