package event

var globalEventHandler *eventHandler

func Global() *eventHandler {
	if globalEventHandler == nil {
		globalEventHandler = New()
	}

	return globalEventHandler
}

func New() *eventHandler {
	e := new(eventHandler)
	e.events = map[uintptr]*event{}
	return e
}
