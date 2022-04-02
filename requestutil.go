package mcmplgo

func Info(msg string) {
	data := make(map[string]interface{})
	data["message"] = msg
	if v, err := Log.Create(data); err == nil {
		v.Fire()
	}
}
