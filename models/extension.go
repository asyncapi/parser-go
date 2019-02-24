package models

import (
	"encoding/json"
	"strings"
)

type ExtensionProps struct {
	Extensions map[string]interface{}
}

// Unmarshals only specification extensions.
func UnmarshalExtensions(value *ExtensionProps, data []byte) error {
	var everything map[string]interface{}
	
	err := json.Unmarshal(data, &everything)
	if (err != nil) {
		return err
	}

	value.Extensions = make(map[string]interface{})
	for k, v := range everything {
		if strings.HasPrefix(k, "x-") {
			value.Extensions[k] = v
		}
	}

	return err
}