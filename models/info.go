package models

import (
	"encoding/json"
)

// Info maps AsyncAPI info object
type Info struct {
	Extensions     map[string]json.RawMessage `json:"-"`
	Title          string                     `json:"title,omitempty"`
	Description    string                     `json:"description,omitempty"`
	TermsOfService string                     `json:"termsOfService,omitempty"`
	Contact        *Contact                   `json:"contact,omitempty"`
	License        *License                   `json:"license,omitempty"`
	Version        string                     `json:"version,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Info) UnmarshalJSON(data []byte) error {
	type Alias Info
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Title = jsonMap.Title
	value.Description = jsonMap.Description
	value.TermsOfService = jsonMap.TermsOfService
	value.Contact = jsonMap.Contact
	value.License = jsonMap.License
	value.Version = jsonMap.Version
	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts
	return nil
}

// MarshalJSON marshals JSON
func (value Info) MarshalJSON() ([]byte, error) {
	type Alias Info
	jsonByteArray, err := json.Marshal(&Alias{
		Title:          value.Title,
		Description:    value.Description,
		TermsOfService: value.TermsOfService,
		Contact:        value.Contact,
		License:        value.License,
		Version:        value.Version,
	})
	if err != nil {
		return nil, err
	}

	return MergeExtensions(jsonByteArray, value.Extensions)
}

// Contact maps AsyncAPI info.contact object
type Contact struct {
	Extensions map[string]json.RawMessage `json:"-"`
	Name       string                     `json:"name,omitempty"`
	URL        string                     `json:"url,omitempty"`
	Email      string                     `json:"email,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Contact) UnmarshalJSON(data []byte) error {
	type Alias Contact
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Name = jsonMap.Name
	value.URL = jsonMap.URL
	value.Email = jsonMap.Email
	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts
	return nil
}

// MarshalJSON marshals JSON
func (value *Contact) MarshalJSON() ([]byte, error) {
	type Alias Contact
	jsonByteArray, err := json.Marshal(&Alias{
		Name:  value.Name,
		URL:   value.URL,
		Email: value.Email,
	})
	if err != nil {
		return nil, err
	}

	return MergeExtensions(jsonByteArray, value.Extensions)
}

// License maps AsyncAPI info.license object
type License struct {
	Extensions map[string]json.RawMessage `json:"-"`
	Name       string                     `json:"name,omitempty"`
	URL        string                     `json:"url,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *License) UnmarshalJSON(data []byte) error {
	type Alias License
	jsonMap := Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Name = jsonMap.Name
	value.URL = jsonMap.URL
	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts
	return nil
}

// MarshalJSON marshals JSON
func (value *License) MarshalJSON() ([]byte, error) {
	type Alias License
	jsonByteArray, err := json.Marshal(&Alias{
		Name: value.Name,
		URL:  value.URL,
	})
	if err != nil {
		return nil, err
	}

	return MergeExtensions(jsonByteArray, value.Extensions)
}
