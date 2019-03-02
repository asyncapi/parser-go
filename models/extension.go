package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// ExtensionsFromJSON obtains AsyncAPI extensions from a JSON object.
func ExtensionsFromJSON(data []byte) (map[string]json.RawMessage, error) {
	exts := make(map[string]json.RawMessage, 0)
	extensions := make(map[string]json.RawMessage, 0)
	// parse all the defined properties
	if err := json.Unmarshal(data, &exts); err != nil {
		return nil, err
	}
	for k, v := range exts {
		if strings.HasPrefix(k, "x-") {
			extensions[k] = v
		}
	}
	return extensions, nil
}

// MergeExtensions merges extensions with a JSON object.
func MergeExtensions(jsonByteArray []byte, value map[string]json.RawMessage) ([]byte, error) {
	if jsonByteArray == nil {
		return nil, errors.New("jsonByteArray can't be nil")
	}

	if value == nil {
		return jsonByteArray, nil
	}

	e, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	if len(e) <= 2 && len(jsonByteArray) <= 2 {
		return []byte(`{}`), nil
	} else if len(e) <= 2 && len(jsonByteArray) > 2 {
		return jsonByteArray, nil
	} else if len(e) > 2 && len(jsonByteArray) <= 2 {
		return e, nil
	}

	var jsonString = string(jsonByteArray)
	var extensionsString = string(e)
	result := jsonString[:1] + extensionsString[1:len(extensionsString)-1] + `,` + jsonString[1:]

	return []byte(result), nil
}
