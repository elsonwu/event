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
	
	evt.Off("save", model)