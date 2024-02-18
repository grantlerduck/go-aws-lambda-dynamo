package app

import (
	"errors"
	"os"
)

type BookingEnv struct {
	TableName string
	Region    string
}

func NewBookingEnv() *BookingEnv {
	tableName := os.Getenv("DYNAMO_BOOKING_TABLE_NAME")
	if tableName == "" {
		panic(errors.New("DYNAMO_BOOKING_TABLE_NAME is not set"))
	}
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = "eu-west-1"
	}
	return &BookingEnv{TableName: tableName, Region: region}
}
