package main

var Params map[string]interface{}

func OnRequest(m map[string]interface{}) {
	for k, v := range Params {
		m[k] = v
	}
}

func SetParams(key string, value interface{}) {
	Params[key] = value
}

func OnResponse(m map[string]interface{}) {
	for k, v := range m {
		Params[k] = v
	}
}

func GetParams(key string) interface{} {
	return Params[key]
}