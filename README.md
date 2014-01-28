event
=====

a simple model event

	model := new(Model)

	evt := event.New()
	evt.On("save", model, func(m interface{}, msgs event.Msgs) error {
		if m2, ok := m.(*Model); ok {
			fmt.Println("name:", m2.Name)
			m2.Name = m2.Name + " changed"
		}
		fmt.Println("saving...1")
		return nil
	})

	evt.On("save", model, func(m interface{}, msgs event.Msgs) error {
		if m2, ok := m.(*Model); ok {
			fmt.Println("name:", m2.Name)
		}
		fmt.Println("saving...2")
		return nil
	})
	
	// off the "save" event on model
	evt.Off("save", model)

	// off everthing on model
	evt.Off("", model)

Note

	it's better to off any useless model, if not Go's gc can't clean the model this evt