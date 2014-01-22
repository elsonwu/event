package event

type EventCallback func(owner interface{}, msgs Msgs) error

type eventHandler struct {
	events []*event
}

func (self *eventHandler) getEventByOwner(owner interface{}) *event {
	for _, e := range self.events {
		if e.owner == owner {
			return e
		}
	}

	return nil
}

func (self *eventHandler) Off(name string, owner interface{}) {
	evt := self.getEventByOwner(owner)
	if evt == nil {
		return
	}

	evt.Detach(name)
}

func (self *eventHandler) On(name string, owner interface{}, callback EventCallback) {

	find := false
	e := self.getEventByOwner(owner)
	if e != nil {
		find = true
		e.Attach(name, callback)
	}

	if !find {
		e := &event{owner: owner}
		e.Attach(name, callback)
		self.events = append(self.events, e)
	}
}

func (self *eventHandler) Emit(name string, owner interface{}, msgs Msgs) error {
	for _, event := range self.events {
		if event.owner != owner {
			continue
		}

		if callbacks, ok := event.callbacks[name]; ok {
			for _, callback := range callbacks {
				if err := callback(owner, msgs); err != nil {
					return err
				}
			}

			break
		}
	}

	return nil
}
