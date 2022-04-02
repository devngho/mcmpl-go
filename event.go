package mcmplgo

import "errors"

type Event struct {
	UUID       string `json:"uuid"`
	Type       string `json:"type"`
	NumberType EventType
	Data       map[string]interface{} `json:"data"`
}

type EventType int

const (
	BlockBreakEvent EventType = 0 + iota
	BlockBurnEvent
	BlockCanBuildEvent

	PlayerJoinEvent
	PlayerQuitEvent
)

func (t EventType) Create(uuid string, data map[string]interface{}) (Event, error) {
	name, err := t.Name()

	if err != nil {
		return Event{}, err
	}

	return Event{
		UUID:       uuid,
		Data:       data,
		Type:       name,
		NumberType: t,
	}, nil
}

func (t EventType) Name() (string, error) {
	result := ""
	for k, v := range events {
		if t == v {
			result = k
		}
	}
	if result == "" {
		return "", errors.New("cant find task in registry")
	}
	return result, nil
}

func (e Event) Fire() {
	if v, ok := listeners[e.NumberType]; ok {
		for _, i := range v {
			i(e)
		}
	}
}
