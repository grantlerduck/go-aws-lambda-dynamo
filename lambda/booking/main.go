package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/grantlerduck/go-was-lambda-dyanmo/lib/booking"
)

func HandleRequest(ctx context.Context, event *booking.Event) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}
	message := fmt.Sprintf("Received booking event %s!", event.BookingId)
	println(message)
	return &event.BookingId, nil
}

func main() {
	lambda.Start(HandleRequest)
}
