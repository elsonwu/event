event
=====

a simple model event

	model := new(Model)

	evt := event.NewEventHandler()
	evt.On("save", model, func(m interface{}) {
		if m2, ok := m.(*Model); ok {
			fmt.Println("name:", m2.Name)
			m2.Name = m2.Name + " changed"
		}
		fmt.Println("saving...1")
	})

	evt.On("save", model, func(m interface{}) {
		if m2, ok := m.(*Model); ok {
			fmt.Println("name:", m2.Name)
		}
		fmt.Println("saving...2")
	})
	
	evt.Off("save", model)