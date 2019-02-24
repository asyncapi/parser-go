package models

import (
	"encoding/json"
)

type Info struct {
	ExtensionProps  ExtensionProps
	Title           string   `json:"title,omitempty"`
	Description     string   `json:"description,omitempty"`
	TermsOfService  string   `json:"termsOfService,omitempty"`
	Contact         *Contact `json:"contact,omitempty"`
	License         *License `json:"license,omitempty"`
	Version         string   `json:"version,omitempty"`
}

func (value *Info) Unmarshal(data []byte) error {
	return json.Unmarshal(data, value)
}

type Contact struct {
	ExtensionProps  ExtensionProps
	Name             string `json:"name,omitempty"`
	URL              string `json:"url,omitempty"`
	Email            string `json:"email,omitempty"`
}

func (value *Contact) Unmarshal(data []byte) error {
	return json.Unmarshal(data, value)
}

type License struct {
	ExtensionProps   ExtensionProps
	Name             string `json:"name,omitempty"`
	URL              string `json:"url,omitempty"`
}

func (value *License) Unmarshal(data []byte) error {
	return json.Unmarshal(data, value)
}