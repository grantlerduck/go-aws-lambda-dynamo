package booking

import (
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	bookingpb "github.com/grantlerduck/go-aws-lambda-dynamo/lambdas/booking-handler/proto"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"time"
)

const timeFormat = "2006-01-02T15:04:05.999999999Z-0700"
const minAllowedEpochMillis = 946724400000

type InvalidEventMessageError struct {
	msg string
}

func (e *InvalidEventMessageError) Error() string {
	return e.msg
}

type Service interface {
	Process(event *EventMessage) (*Event, error)
}

type EventProcessor struct {
	repo   Repository
	logger *zap.Logger
}

func (ep *EventProcessor) Process(eventmsg *EventMessage) (*Event, error) {
	pb, decodeErr := ep.decodePayload(eventmsg.Payload)
	if decodeErr != nil {
		ep.logger.Error("failed to process eventmsg",
			zap.String("key", eventmsg.Key),
			zap.Error(decodeErr),
		)
		return nil, decodeErr
	}
	validationErr := ep.validate(pb)
	if validationErr != nil {
		ep.logger.Error("failed to process eventmsg",
			zap.String("key", eventmsg.Key),
			zap.Error(validationErr),
		)
		return nil, validationErr
	}
	evnt := ep.mapToEvent(pb)
	result, repoErr := ep.repo.Insert(evnt)
	if repoErr != nil {
		ep.logger.Error("failed to process eventmsg",
			zap.String("bookingId", evnt.BookingId),
			zap.String("state", evnt.BookingState.String()),
			zap.Error(repoErr),
		)
		return nil, repoErr
	}
	return result, nil
}

func (ep *EventProcessor) decodePayload(payload string) (*bookingpb.Event, error) {
	bytes, decodeErr := base64.StdEncoding.DecodeString(payload)
	if decodeErr != nil {
		ep.logger.Error("failed to decode event payload",
			zap.Error(decodeErr),
		)
		return nil, decodeErr
	}
	var evntpb bookingpb.Event
	marshallErr := proto.Unmarshal(bytes, &evntpb)
	if marshallErr != nil {
		ep.logger.Error("failed to unmarshal payload",
			zap.Error(marshallErr),
		)
		return nil, marshallErr
	}
	initErr := proto.CheckInitialized(&evntpb)
	if initErr != nil {
		ep.logger.Error("failed to properly initialize payload",
			zap.Error(initErr),
		)
		return nil, initErr
	}
	return &evntpb, nil
}

func (ep *EventProcessor) validate(evntpb *bookingpb.Event) error {
	if evntpb.ToEpochMillis <= minAllowedEpochMillis || evntpb.FromEpochMillis <= minAllowedEpochMillis {
		return &InvalidEventMessageError{"invalid epoch milliseconds for 'to' and 'from' of message"}
	}
	if !ep.isUUID(evntpb.HotelId) || len(evntpb.HotelName) < 5 {
		return &InvalidEventMessageError{
			fmt.Sprintf("invalid hotel information hotelId=%s, hotelName=%s", evntpb.HotelId, evntpb.HotelName),
		}
	}
	if len(evntpb.AirlineName) < 3 || !ep.isUUID(evntpb.FlightId) {
		return &InvalidEventMessageError{
			fmt.Sprintf("invalid flight information flightId=%s, airlineName=%s", evntpb.FlightId, evntpb.AirlineName),
		}
	}
	if !ep.isUUID(evntpb.UserId) || !ep.isUUID(evntpb.BookingId) {
		return &InvalidEventMessageError{
			fmt.Sprintf("invalid booking information userID=%s", evntpb.UserId),
		}
	}
	return nil
}

func (ep *EventProcessor) isUUID(str string) bool {
	_, err := uuid.Parse(str)
	return err == nil
}

func (ep *EventProcessor) mapToEvent(evntpb *bookingpb.Event) *Event {
	event := Event{}
	event.EventId = uuid.New().String()
	event.BookingId = evntpb.BookingId
	event.UserId = evntpb.UserId
	event.TripFrom = time.UnixMilli(evntpb.FromEpochMillis).Format(timeFormat)
	event.TripUntil = time.UnixMilli(evntpb.ToEpochMillis).Format(timeFormat)
	event.HotelName = evntpb.HotelName
	event.HotelId = evntpb.HotelId
	event.FlightId = evntpb.FlightId
	event.AirlineName = evntpb.AirlineName
	event.BookingState = getState(evntpb.BookingState.String())
	return &event
}

func NewBookingService(repository Repository, logger *zap.Logger) *EventProcessor {
	return &EventProcessor{repo: repository, logger: logger}
}
