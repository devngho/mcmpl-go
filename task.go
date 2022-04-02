package mcmplgo

import "errors"

type Task struct {
	UUID       string                 `json:"uuid"`
	Type       string                 `json:"type"`
	Data       map[string]interface{} `json:"data"`
	NumberType TaskType
}

type TaskType int

const (
	OnEnable TaskType = 0 + iota
	EventTask
)

func (t TaskType) Create(uuid string, data map[string]interface{}) (Task, error) {
	name, err := t.Name()

	if err != nil {
		return Task{}, err
	}

	return Task{
		UUID:       uuid,
		Data:       data,
		Type:       name,
		NumberType: t,
	}, nil
}

func (t TaskType) Name() (string, error) {
	result := ""
	for k, v := range tasks {
		if t == v {
			result = k
		}
	}
	if result == "" {
		return "", errors.New("cant find task in registry")
	}
	return result, nil
}

func (e Task) Fire() {
	if v, ok := taskListeners[e.NumberType]; ok {
		for _, i := range v {
			i(e)
		}
	}
}
