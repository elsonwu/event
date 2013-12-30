package event

func NewEventHandler() *eventHandler {
	e := new(eventHandler)
	e.events = []*event{}
	return e
}
