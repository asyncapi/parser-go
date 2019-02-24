package models

import (
	"encoding/json"
)

// Info maps AsyncAPI info object
type Info struct {
	ExtensionProps ExtensionProps `json:"-"`
	Title          string         `json:"title,omitempty"`
	Description    string         `json:"description,omitempty"`
	TermsOfService string         `json:"termsOfService,omitempty"`
	Contact        *Contact       `json:"contact,omitempty"`
	License        *License       `json:"license,omitempty"`
	Version        string         `json:"version,omitempty"`
}

// Unmarshal unmarshals JSON
func (value *Info) Unmarshal(data []byte) error {
	return json.Unmarshal(data, value)
}

// Marshal marshals JSON
func (value *Info) Marshal() ([]byte, error) {
	return json.Marshal(value)
}

// Contact maps AsyncAPI info.contact object
type Contact struct {
	ExtensionProps ExtensionProps `json:"-"`
	Name           string         `json:"name,omitempty"`
	URL            string         `json:"url,omitempty"`
	Email          string         `json:"email,omitempty"`
}

// Unmarshal unmarshals JSON
func (value *Contact) Unmarshal(data []byte) error {
	return json.Unmarshal(data, value)
}

// Marshal marshals JSON
func (value *Contact) Marshal() ([]byte, error) {
	return json.Marshal(value)
}

// License maps AsyncAPI info.license object
type License struct {
	ExtensionProps ExtensionProps `json:"-"`
	Name           string         `json:"name,omitempty"`
	URL            string         `json:"url,omitempty"`
}

// Unmarshal unmarshals JSON
func (value *License) Unmarshal(data []byte) error {
	return json.Unmarshal(data, value)
}

// Marshal marshals JSON
func (value *License) Marshal() ([]byte, error) {
	return json.Marshal(value)
}
