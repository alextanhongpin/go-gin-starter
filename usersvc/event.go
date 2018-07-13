package usersvc

// Event represent the events emitted by the user service
type Event struct {
	// Client *webhook.Client
}

// Use the verb-noun naming convention, and only submit events that mutates the state

// func (e *Event) CreatedUser(name string) {
// Emit webhook event here
// }

// NewEvent returns a pointer of the Event
func NewEvent() *Event {
	return &Event{}
}
