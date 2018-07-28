package main

var Params map[string]interface{}

func init() {
	Params = make(map[string]interface{})
}

func Apply(m map[string]interface{}) {
	for k, v := range Params {
		m[k] = v
	}
}

func Save(m map[string]interface{}) {
	for k, v := range m {
		Params[k] = v
	}
}

// go build -buildmode=plugin
