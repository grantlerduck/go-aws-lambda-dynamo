package booking

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBookingDomain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Booking Domain Suite")
}
