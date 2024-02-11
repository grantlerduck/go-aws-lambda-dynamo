package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/grantlerduck/go-was-lambda-dyanmo/lib/booking"
	"testing"
	"time"
)

func TestHandleRequest(t *testing.T) {
	var event = booking.Event{
		BookingId:    uuid.New().String(),
		UserId:       uuid.New().String(),
		TripFrom:     time.RFC3339,
		TripUntil:    time.RFC3339,
		HotelName:    "mockHotel",
		HotelId:      uuid.New().String(),
		FlightId:     uuid.New().String(),
		AirlineName:  "cheap-airline",
		BookingState: "PAYMENT_PENDING",
	}
	var bookingId = event.BookingId
	var ctx = context.Background()
	var result, err = HandleRequest(ctx, &event)
	if err != nil {
		t.Error(err)
	}
	if *result != bookingId {
		t.Errorf("Booking id was %s, expected %s", *result, bookingId)
	}
}

func TestHandleRequestNilEvent(t *testing.T) {
	var ctx = context.Background()
	var _, err = HandleRequest(ctx, nil)
	if err == nil {
		t.Error(err)
	}
}
