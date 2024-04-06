package dynamo

import "strings"

type State string

const (
	Unconfirmed     State = "booking-unconfirmed"
	Confirmed       State = "booking-confirmed"
	PaymentReceived State = "booking-fee-payed"
	PaymentPending  State = "booking-fee-pending"
	Planned         State = "booking-planned"
	Canceled        State = "booking-canceled"
	CheckedIn       State = "checked-in"
	CheckedOut      State = "checked-out"
	ReviewPending   State = "review-pending"
	Reviewed        State = "customer-reviewed"
	Unknown         State = "unknown"
)

func (s State) String() string {
	return string(s)
}

var states = map[State]struct{}{
	Unconfirmed:     {},
	Confirmed:       {},
	PaymentReceived: {},
	PaymentPending:  {},
	Planned:         {},
	Canceled:        {},
	CheckedIn:       {},
	CheckedOut:      {},
	ReviewPending:   {},
	Reviewed:        {},
	Unknown:         {},
}

func getState(str string) State {
	state := State(strings.ToLower(str))
	_, ok := states[state]
	if !ok {
		return Unknown
	}
	return state
}
