package booking

import "strings"

type State string

// Just getting used to enums in go, would replace it with just protobuf
const (
	Unconfirmed     State = "unconfirmed"
	Confirmed       State = "confirmed"
	PaymentReceived State = "payment_received"
	PaymentPending  State = "payment_pending"
	Planned         State = "planned"
	Canceled        State = "canceled"
	CheckedIn       State = "checked_in"
	CheckedOut      State = "checked_out"
	ReviewPending   State = "review_pending"
	Reviewed        State = "reviewed"
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
