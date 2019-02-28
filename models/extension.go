package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// Extension holds a collection of specification extensions
type Extensions = map[string]interface{}

// UnmarshalExtensions unmarshals only specification extensions.
func UnmarshalExtensions(data []byte) (Extensions, error) {
	var everything map[string]interface{}

	err := json.Unmarshal(data, &everything)
	if err != nil {
		return nil, err
	}

	result := make(Extensions)
	for k, v := range everything {
		if strings.HasPrefix(k, "x-") {
			result[k] = v
		}
	}

	return result, err
}

// MarshalWithExtensions marshals only specification extensions.
func MarshalWithExtensions(object interface{}, value Extensions) ([]byte, error) {
	if value == nil {
		return nil, errors.New("Extensions can't be nil.")
	}
	if object == nil {
		return nil, errors.New("Object can't be nil.")
	}

	e, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	j, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}

	if len(e) <= 2 && len(j) <= 2 {
		return []byte(`{}`), nil
	} else if len(e) <= 2 && len(j) > 2 {
		return j, nil
	} else if len(e) > 2 && len(j) <= 2 {
		return e, nil
	}

	var jsonString = string(j)
	var extensionsString = string(e)
	result := jsonString[:1] + extensionsString[1:len(extensionsString)-1] + `,` + jsonString[1:]

	return []byte(result), nil
}
