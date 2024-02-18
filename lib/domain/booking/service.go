package booking

import (
	"go.uber.org/zap"
)

type Service interface {
	Process(event *Event) (*Event, error)
}

type EvenProcessor struct {
	repo   Repository
	logger *zap.Logger
}

func (ep *EvenProcessor) Process(event *Event) (*Event, error) {
	result, err := ep.repo.Insert(event)
	if err != nil {
		ep.logger.Error("Failed to process event",
			zap.Any("item", event),
			zap.Error(err),
		)
		return nil, err
	}
	return result, nil
}

func NewBookingService(repository Repository, logger *zap.Logger) *EvenProcessor {
	return &EvenProcessor{repo: repository, logger: logger}
}
