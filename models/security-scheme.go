package models

import "encoding/json"

// SecurityScheme maps AsyncAPI "SecurityScheme" object
type SecurityScheme struct {
	Extensions       map[string]json.RawMessage `json:"-"`
	Type             string                     `json:"type,omitempty"`
	Description      string                     `json:"description,omitempty"`
	In               string                     `json:"in,omitempty"`
	Flows            json.RawMessage            `json:"flows,omitempty"`
	OpenIdConnectUrl string                     `json:"openIdConnectUrl,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *SecurityScheme) UnmarshalJSON(data []byte) error {
	type SecuritySchemeAlias SecurityScheme
	jsonMap := SecuritySchemeAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value SecurityScheme) MarshalJSON() ([]byte, error) {
	type SecuritySchemeAlias SecurityScheme
	jsonByteArray, err := json.Marshal(&SecuritySchemeAlias{
		Type:             value.Type,
		Description:      value.Description,
		In:               value.In,
		Flows:            value.Flows,
		OpenIdConnectUrl: value.OpenIdConnectUrl,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// HttpSecurityScheme maps AsyncAPI "HTTPSecurityScheme" object
type HttpSecurityScheme struct {
	Extensions   map[string]json.RawMessage `json:"-"`
	Scheme       string                     `json:"scheme,omitempty"`
	Description  string                     `json:"description,omitempty"`
	Type         string                     `json:"type,omitempty"`
	BearerFormat string                     `json:"bearerFormat,omitempty"`
	Name         string                     `json:"name,omitempty"`
	In           string                     `json:"in,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *HttpSecurityScheme) UnmarshalJSON(data []byte) error {
	type HttpSecuritySchemeAlias HttpSecurityScheme
	jsonMap := HttpSecuritySchemeAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value HttpSecurityScheme) MarshalJSON() ([]byte, error) {
	type HttpSecuritySchemeAlias HttpSecurityScheme
	jsonByteArray, err := json.Marshal(&HttpSecuritySchemeAlias{
		Scheme:       value.Scheme,
		Description:  value.Description,
		Type:         value.Type,
		BearerFormat: value.BearerFormat,
		Name:         value.Name,
		In:           value.In,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}
