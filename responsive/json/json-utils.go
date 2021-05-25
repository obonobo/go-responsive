package json

import (
	"bytes"
	"encoding/json"
)

func Stringify(body Json) string {
	var buf bytes.Buffer
	marshalled, err := json.Marshal(body)
	if err != nil {
		return ""
	}
	json.HTMLEscape(&buf, marshalled)
	return buf.String()
}

func Parse(body string) Json {
	var res Json
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		return Json{}
	}
	return res
}
