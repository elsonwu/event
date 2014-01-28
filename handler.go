package event

import (
	"reflect"
)

type EventCallback func(owner interface{}, msgs Msgs) error

type eventHandler struct {
	events map[uintptr]*event
}

func (self *eventHandler) getEventOwnerKey(owner interface{}) uintptr {
	return reflect.ValueOf(owner).Pointer()
}

func (self *eventHandler) getEventByOwner(owner interface{}) (*event, uintptr) {
	ptr := self.getEventOwnerKey(owner)
	if e, ok := self.events[ptr]; ok {
		return e, ptr
	}

	return nil, ptr
}

func (self *eventHandler) Off(name string, owner interface{}) {
	evt, ptr := self.getEventByOwner(owner)
	if evt == nil {
		return
	}

	if name == "" {
		delete(self.events, ptr)
		return
	}

	evt.Detach(name)
	if evt.IsEmpty() {
		delete(self.events, ptr)
	}
}

func (self *eventHandler) On(name string, owner interface{}, callback EventCallback) {
	e, ptr := self.getEventByOwner(owner)
	if e != nil {
		e.Attach(name, callback)
		return
	}

	e = &event{owner: owner}
	e.Attach(name, callback)
	self.events[ptr] = e
}

func (self *eventHandler) Emit(name string, owner interface{}, msgs Msgs) error {
	evt, _ := self.getEventByOwner(owner)
	if evt == nil {
		return nil
	}

	if callbacks, ok := evt.callbacks[name]; ok {
		for _, callback := range callbacks {
			if err := callback(owner, msgs); err != nil {
				return err
			}
		}
	}

	return nil
}
