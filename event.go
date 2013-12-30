package event

type event struct {
	owner     interface{}
	callbacks map[string][]EventCallback
}

func (self *event) Attach(name string, callback EventCallback) {
	if self.callbacks == nil {
		self.callbacks = map[string][]EventCallback{}
	}

	self.callbacks[name] = append(self.callbacks[name], callback)
}

func (self *event) DetachAll() {
	self.callbacks = nil
}

func (self *event) Detach(name string) {
	if self.callbacks == nil {
		return
	}

	if _, ok := self.callbacks[name]; ok {
		delete(self.callbacks, name)
	}
}

func (self *event) Callbacks(name string) []EventCallback {
	if _, ok := self.callbacks[name]; ok {
		return self.callbacks[name]
	}

	return nil
}
