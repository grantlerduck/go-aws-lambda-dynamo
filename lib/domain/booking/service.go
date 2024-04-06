package booking

import (
	"go.uber.org/zap"
)

type Service interface {
	Process(event *Event) (*Event, error)
}

type EventProcessor struct {
	repo   Repository
	logger *zap.Logger
}

func (ep *EventProcessor) Process(event *Event) (*Event, error) {
	result, err := ep.repo.Insert(event)
	if err != nil {
		ep.logger.Error("failed to process event",
			zap.String("bookingId", event.BookingId),
			zap.String("state", event.BookingState),
			zap.Error(err),
		)
		return nil, err
	}
	return result, nil
}

func NewBookingService(repository Repository, logger *zap.Logger) *EventProcessor {
	return &EventProcessor{repo: repository, logger: logger}
}
