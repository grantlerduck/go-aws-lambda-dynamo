package booking

import (
	bookingpb "github.com/grantlerduck/go-aws-lambda-dynamo/lambdas/booking-handler/proto"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Booking state string", func() {
	When("is mappable to Unconfirmed", func() {
		unconfirmed := "UNCONFIRMED"
		expectedResult := Unconfirmed
		It("returns correctly", func() {
			result := getState(unconfirmed)
			Expect(result).To(Equal(expectedResult))
		})
		It("gets string correctly", func() {
			result := Unconfirmed.String()
			Expect(result).To(Equal("unconfirmed"))
		})
	})
	When("is mappable to protobuf unconfirmed", func() {
		unconfirmed := bookingpb.State(1)
		expectedResult := Unconfirmed
		It("returns correctly", func() {
			result := getState(unconfirmed.String())
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to Confirmed", func() {
		confirmed := "Confirmed"
		expectedResult := Confirmed
		It("returns correctly", func() {
			result := getState(confirmed)
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to protobuf confirmed", func() {
		confirmed := bookingpb.State(2)
		expectedResult := Confirmed
		It("returns correctly", func() {
			result := getState(confirmed.String())
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to PaymentReceived", func() {
		payed := "payment_reCeived"
		expectedResult := PaymentReceived
		It("returns correctly", func() {
			result := getState(payed)
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to protobuf payment_received", func() {
		payed := bookingpb.State(3)
		expectedResult := PaymentReceived
		It("returns correctly", func() {
			result := getState(payed.String())
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to PaymentPending", func() {
		pending := "Payment_Pending"
		expectedResult := PaymentPending
		It("returns correctly", func() {
			result := getState(pending)
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to protobuf payment_pending", func() {
		pending := bookingpb.State(4)
		expectedResult := PaymentPending
		It("returns correctly", func() {
			result := getState(pending.String())
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to Planned", func() {
		planned := "Planned"
		expectedResult := Planned
		It("returns correctly", func() {
			result := getState(planned)
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to protobuf planned", func() {
		planned := bookingpb.State(5)
		expectedResult := Planned
		It("returns correctly", func() {
			result := getState(planned.String())
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to Canceled", func() {
		canceled := "canCeled"
		expectedResult := Canceled
		It("returns correctly", func() {
			result := getState(canceled)
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to protobuf canceled", func() {
		canceled := bookingpb.State(6)
		expectedResult := Canceled
		It("returns correctly", func() {
			result := getState(canceled.String())
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to CheckedIn", func() {
		checkedIn := "Checked_In"
		expectedResult := CheckedIn
		It("returns correctly", func() {
			result := getState(checkedIn)
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to protobuf checked_in", func() {
		checkedIn := bookingpb.State(7)
		expectedResult := CheckedIn
		It("returns correctly", func() {
			result := getState(checkedIn.String())
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to CheckedOut", func() {
		checkedOut := "checked_ouT"
		expectedResult := CheckedOut
		It("returns correctly", func() {
			result := getState(checkedOut)
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to protobuf checked_out", func() {
		checkedOut := bookingpb.State(8)
		expectedResult := CheckedOut
		It("returns correctly", func() {
			result := getState(checkedOut.String())
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to ReviewPending", func() {
		reviewPending := "review_pending"
		expectedResult := ReviewPending
		It("returns correctly", func() {
			result := getState(reviewPending)
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to protobuf review_pending", func() {
		reviewPending := bookingpb.State(9)
		expectedResult := ReviewPending
		It("returns correctly", func() {
			result := getState(reviewPending.String())
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to Reviewed", func() {
		reviewed := "reviewed"
		expectedResult := Reviewed
		It("returns correctly", func() {
			result := getState(reviewed)
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable to protobuf reviewed", func() {
		reviewed := bookingpb.State(10)
		expectedResult := Reviewed
		It("returns correctly", func() {
			result := getState(reviewed.String())
			Expect(result).To(Equal(expectedResult))
		})
	})
	When("is mappable not mappable to meaningful state", func() {
		someString := "qwelfnqälwrjf"
		expectedResult := Unknown
		It("returns Unknown", func() {
			result := getState(someString)
			Expect(result).To(Equal(expectedResult))
		})
	})
})
