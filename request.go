package mcmplgo

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Request struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type RequestType int

const (
	Log RequestType = 0 + iota
)

func (t RequestType) Create(data map[string]interface{}) (Request, error) {
	name, err := t.Name()

	if err != nil {
		return Request{}, err
	}

	return Request{
		Data: data,
		Type: name,
	}, nil
}

func (t RequestType) Name() (string, error) {
	result := ""
	for k, v := range requests {
		if t == v {
			result = k
		}
	}
	if result == "" {
		return "", errors.New("cant find task in registry")
	}
	return result, nil
}

func (t Request) Fire() {
	if v, ok := json.Marshal(t); ok == nil {
		fmt.Println(string(v))
	}
}
