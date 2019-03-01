package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Server maps AsyncAPI serves.item object
type Server struct {
	Extensions    Extensions       `json:"-"`
	URL           string           `json:"url"`
	Description   string           `json:"description,omitempty"`
	Scheme        string           `json:"scheme"`
	SchemeVersion string           `json:"schemeVersion,omitempty"`
	Variables     *ServerVariables `json:"variables,omitempty"`
	BaseChannel   string           `json:"baseChannel,omitempty"`
	// Security      []*SecurityRequirement `json:"security,omitempty"`
}

// Unmarshal unmarshals JSON
func (value *Server) Unmarshal(data []byte) error {
	return json.Unmarshal(data, value)
}

// Marshal marshals JSON
func (value *Server) Marshal() ([]byte, error) {
	if value.Extensions == nil {
		return json.Marshal(value)
	}
	return MarshalWithExtensions(value, value.Extensions)
}

// ServerVariables object
type ServerVariables struct {
	AdditionalProperties map[string]*ServerVariable `json:",omitempty"`
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
func (value *ServerVariable) Unmarshal(data []byte) error {
	return json.Unmarshal(data, value)
}

// Marshal marshals JSON
func (value *ServerVariable) Marshal() ([]byte, error) {
	return json.Marshal(value)
}
