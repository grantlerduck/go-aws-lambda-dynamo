package booking

import (
	"encoding/base64"
	bookingpb "github.com/grantlerduck/go-aws-lambda-dynamo/proto"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"time"
)

const timeFormat = "2006-01-02T15:04:05.999999999Z-0700"

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
	evnt := ep.mapToEvent(pb, eventmsg.Key)
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
	bytes, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		ep.logger.Error("failed to decode event payload",
			zap.Error(err),
		)
		return nil, err
	}
	var evntpb bookingpb.Event
	marshallErr := proto.Unmarshal(bytes, &evntpb)
	if marshallErr != nil {
		ep.logger.Error("failed to unmarshal payload",
			zap.Error(err),
		)
	}
	return &evntpb, nil
}

func (ep *EventProcessor) mapToEvent(evntpb *bookingpb.Event, key string) *Event {
	event := Event{}
	event.BookingId = key
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
