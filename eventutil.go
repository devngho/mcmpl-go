package mcmplgo

func (t TaskType) Listen(listener TaskListener) {
	if v, ok := taskListeners[t]; ok {
		taskListeners[t] = append(v, listener)
	} else {
		taskListeners[t] = []TaskListener{listener}
		if name, err := t.Name(); err == nil {
			data := make(map[string]interface{})
			data["type"] = name

			if v, err := RegisterTask.Create(data); err == nil {
				v.Fire()
			}
		}
	}
}

func (t EventType) Listen(listener EventListener) {
	if len(listeners) == 0 {
		data := make(map[string]interface{})
		data["type"] = "Event"

		if v, err := RegisterTask.Create(data); err == nil {
			v.Fire()
		}
	}
	if v, ok := listeners[t]; ok {
		listeners[t] = append(v, listener)
	} else {
		listeners[t] = []EventListener{listener}
		if name, err := t.Name(); err == nil {
			data := make(map[string]interface{})
			data["type"] = name

			if v, err := RegisterEvent.Create(data); err == nil {
				v.Fire()
			}
		}
	}
}
