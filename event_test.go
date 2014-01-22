package event

import (
	"testing"
)

type TestEvent struct {
	Name string
}

func newTestEvent() *event {
	te := new(TestEvent)
	te.Name = "test"
	e := &event{}
	e.owner = te
	return e
}

func TestEvent1(t *testing.T) {
	e := newTestEvent()
	e.Attach("t1", func(m interface{}, msgs Msgs) error {
		if m2, ok := m.(TestEvent); !ok {
			t.Error("the callback model is not TestEvent")
		} else {
			if m2 != e.owner {
				t.Error("the callback model does not the same as the owner")
			}

			if e2, ok := e.owner.(TestEvent); ok {
				if m2.Name != e2.Name {
					t.Error("the name value has been changed")
				}
			} else {
				t.Error("the owenr is not TestEvent")
			}
		}

		return nil
	})

	if 1 != len(e.callbacks) {
		t.Error("the callbacks len is wrong")
	}

	e.Detach("t1")

	if 0 != len(e.callbacks) {
		t.Error("the callbacks is still there, Detach does not work well")
	}

	e.Attach("t2", func(m interface{}, msgs Msgs) error { return nil })
	e.Attach("t2", func(m interface{}, msgs Msgs) error { return nil })
	e.Attach("t2", func(m interface{}, msgs Msgs) error { return nil })

	if 3 != len(e.Callbacks("t2")) {
		t.Error("the callbacks does not bind correctly")
	}

	e.DetachAll()

	if 0 != len(e.Callbacks("t2")) {
		t.Error("the Detachall does not work well")
	}
}
