package utils

import (
	"bytes"
	"encoding/json"
)

// 不转码的jsonMarshal
func JsonMarshal(data interface{}) string {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	_ = jsonEncoder.Encode(data)
	return bf.String()
}
