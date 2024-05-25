package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/grantlerduck/go-aws-lambda-dynamo/lambdas/booking-handler/adapter/dynamo"
	"github.com/grantlerduck/go-aws-lambda-dynamo/lambdas/booking-handler/app"
	"github.com/grantlerduck/go-aws-lambda-dynamo/lambdas/booking-handler/domain/booking"
	"go.uber.org/zap"
	"log"
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer func() {
		_ = logger.Sync()
	}()
}

type FailedToProcessError struct {
	event booking.EventMessage
}

func (e *FailedToProcessError) Error() string {
	return fmt.Sprintf("failed to process event %s", e.event.Key)
}

type EventNilError struct{}

func (e *EventNilError) Error() string {
	return "event is nil"
}

type BookingHandler struct {
	service booking.Service
	logger  *zap.Logger
}

func (handler *BookingHandler) HandleRequest(ctx context.Context, event *booking.EventMessage) (*string, error) {
	if event == nil {
		return nil, &EventNilError{}
	}
	handler.logger.Info("received booking event", zap.Any("eventId", event.Key))
	ev, err := handler.service.Process(event)
	if err != nil {
		return nil, &FailedToProcessError{*event}
	}
	return &ev.BookingId, nil
}

func main() {
	defer func() {
		_ = logger.Sync()
	}()
	env := app.NewBookingEnv()
	repo := dynamo.NewEventRepository(env.Region, env.TableName, logger)
	service := booking.NewBookingService(repo, logger)
	handler := &BookingHandler{service: service, logger: logger}
	lambda.Start(handler.HandleRequest)
}
