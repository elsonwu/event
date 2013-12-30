package event

import (
	"strconv"
	"testing"
)

type TestHandlerOwner struct {
	Name string
}

func TestHandler1(t *testing.T) {
	owner := new(TestHandlerOwner)
	h := NewEventHandler()

	i := 1
	h.On("do", owner, func(m interface{}) {
		if m2, ok := m.(*TestHandlerOwner); !ok {
			t.Error("the model does not a TestHandlerOwner")
		} else {
			m2.Name = "test" + strconv.Itoa(i)
			i++
		}
	})

	h.Emit("do", owner)

	if 2 != i {
		t.Error("the call back does not work")
	}

	ev := h.getEventByOwner(owner)
	if ev == nil {
		t.Error("event does not save with the owner")
	} else if ev.owner == nil {
		t.Error("event does not have a owner")
	} else if m2, ok := ev.owner.(*TestHandlerOwner); !ok {
		t.Error("event's owner is not a TestHandlerOwner")
	} else if m2.Name != "test1" {
		t.Error("event's owner has not been updated")
	}

	h.Emit("do", owner)

	if owner.Name != "test2" {
		t.Error("the callback not work at second time")
	}

	h.Off("do", owner)

	if owner.Name != "test2" {
		t.Error("the callback still run after set off")
	}
}
