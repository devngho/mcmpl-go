package mcmplgo

func (t TaskType) Listen(listener TaskListener) {
	if v, ok := taskListeners[t]; ok {
		taskListeners[t] = append(v, listener)
	} else {
		taskListeners[t] = []TaskListener{listener}
	}
}

func (t EventType) Listen(listener EventListener) {
	if v, ok := listeners[t]; ok {
		listeners[t] = append(v, listener)
	} else {
		listeners[t] = []EventListener{listener}
	}
}
