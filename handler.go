package mcmplgo

import (
	"bufio"
	"encoding/json"
	"os"
)

func Run() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			var inp map[string]interface{}
			json.Unmarshal(scanner.Bytes(), &inp)
			data := inp["data"].(map[string]interface{})
			if v, err := tasks[inp["type"].(string)].Create(inp["uuid"].(string), data); err != nil {
				if v.NumberType == EventTask {
					Event{
						UUID:       v.UUID,
						Type:       data["type"].(string),
						NumberType: events[data["type"].(string)],
						Data:       data,
					}.Fire()
				} else {
					v.Fire()
				}
			}
		}
	}
}
