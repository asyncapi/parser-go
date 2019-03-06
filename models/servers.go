package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

// Server maps AsyncAPI serves.item object
type Server struct {
	Extensions    Extensions             `json:"-"`
	Url           string                 `json:"url"`
	Description   string                 `json:"description,omitempty"`
	Scheme        string                 `json:"scheme"`
	SchemeVersion string                 `json:"schemeVersion,omitempty"`
	Variables     *ServerVariables       `json:"variables,omitempty"`
	BaseChannel   string                 `json:"baseChannel,omitempty"`
	Security      []*SecurityRequirement `json:"security,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Server) UnmarshalJSON(b []byte) error {
	// schemeReceived := false
	// urlReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "baseChannel":
			if err := json.Unmarshal([]byte(v), &value.BaseChannel); err != nil {
				return err
			}
		case "description":
			if err := json.Unmarshal([]byte(v), &value.Description); err != nil {
				return err
			}
		case "scheme":
			if err := json.Unmarshal([]byte(v), &value.Scheme); err != nil {
				return err
			}
			// schemeReceived = true
		case "schemeVersion":
			if err := json.Unmarshal([]byte(v), &value.SchemeVersion); err != nil {
				return err
			}
		case "security":
			if err := json.Unmarshal([]byte(v), &value.Security); err != nil {
				return err
			}
		case "url":
			if err := json.Unmarshal([]byte(v), &value.Url); err != nil {
				return err
			}
			// urlReceived = true
		case "variables":
			if err := json.Unmarshal([]byte(v), &value.Variables); err != nil {
				return err
			}
		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}
	return nil
}

// func (value *Server) Marshal() ([]byte, error) {
// 	if value.Extensions == nil {
// 		return json.Marshal(value)
// 	}
// 	return MarshalWithExtensions(value, value.Extensions)
// }
// IsZeroOfUnderlyingType
func IsZeroOfUnderlyingType(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

// MarshalJSON marshals JSON
func (value *Server) MarshalJSON() ([]byte, error) {
	// type serverAlias Server
	// if value.Extensions != nil {
	// 	type serverAlias Server
	// 	return MarshalWithExtensions(serverAlias, value.Extensions)
	// }

	return json.Marshal(map[string]interface{}{
		"baseChannel":   value.BaseChannel,
		"description":   value.Description,
		"scheme":        value.Scheme,
		"schemeVersion": value.SchemeVersion,
		"security":      value.Security,
		"url":           value.Url,
		"variables":     value.Variables,
	})
}

// ServerVariables object
type ServerVariables struct {
	AdditionalProperties map[string]*ServerVariable `json:"-,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (strct *ServerVariables) UnmarshalJSON(b []byte) error {
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		// an additional "*ServerVariable" value
		var additionalValue *ServerVariable
		if err := json.Unmarshal([]byte(v), &additionalValue); err != nil {
			return err // invalid additionalProperty
		}
		if strct.AdditionalProperties == nil {
			strct.AdditionalProperties = make(map[string]*ServerVariable, 0)
		}
		strct.AdditionalProperties[k] = additionalValue
	}
	return nil
}

// MarshalJSON marshals JSON
func (strct *ServerVariables) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal any additional Properties
	for k, v := range strct.AdditionalProperties {
		if comma {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf("\"%s\":", k))
		if tmp, err := json.Marshal(v); err != nil {
			return nil, err
		} else {
			buf.Write(tmp)
		}
		comma = true
	}

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

// ServerVariable An object representing a Server Variable for server URL template substitution.
type ServerVariable struct {
	Default     string   `json:"default,omitempty"`
	Description string   `json:"description,omitempty"`
	Enum        []string `json:"enum,omitempty"`
	Examples    []string `json:"examples,omitempty"`
}

// Unmarshal unmarshals JSON
// func (value *ServerVariable) Unmarshal(data []byte) error {
// 	return json.Unmarshal(data, value)
// }

// // Marshal marshals JSON
// func (value *ServerVariable) Marshal() ([]byte, error) {
// 	return json.Marshal(value)
// }

// SecurityRequirement an Object representing a map
type SecurityRequirement struct {
	AdditionalProperties map[string][]string `json:"-,omitempty"`
}
