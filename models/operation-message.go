package models

import "encoding/json"

// OperationMessage maps AsyncAPI "operationMessage" object
type OperationMessage struct {
	Extensions   map[string]json.RawMessage `json:"-"`
	SchemaFormat string                     `json:"schemaFormat,omitempty"`
	ContentType  string                     `json:"contentType,omitempty"`
	Headers      *Schema                    `json:"headers,omitempty"`
	Payload      json.RawMessage            `json:"payload,omitempty"`
	Tags         []*Tag                     `json:"tags,omitempty"`
	Summary      string                     `json:"summary,omitempty"`
	Name         string                     `json:"name,omitempty"`
	Title        string                     `json:"title,omitempty"`
	Description  string                     `json:"description,omitempty"`
	ExternalDocs *ExternalDocs              `json:"externalDocs,omitempty"`
	Deprecated   bool                       `json:"deprecated,omitempty"`
	Example      json.RawMessage            `json:"example,omitempty"`
	OneOf        []*Message                 `json:"oneOf,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *OperationMessage) UnmarshalJSON(data []byte) error {
	type OperationMessageAlias OperationMessage
	jsonMap := OperationMessageAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.SchemaFormat = jsonMap.SchemaFormat
	value.ContentType = jsonMap.ContentType
	value.Headers = jsonMap.Headers
	value.Payload = jsonMap.Payload
	value.Tags = jsonMap.Tags
	value.Summary = jsonMap.Summary
	value.Name = jsonMap.Name
	value.Title = jsonMap.Title
	value.Description = jsonMap.Description
	value.ExternalDocs = jsonMap.ExternalDocs
	value.Deprecated = jsonMap.Deprecated
	value.Example = jsonMap.Example
	value.OneOf = jsonMap.OneOf

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value OperationMessage) MarshalJSON() ([]byte, error) {
	type OperationMessageAlias OperationMessage
	if value.OneOf != nil {
		jsonByteArray, err := json.Marshal(value.OneOf)
		if err != nil {
			return nil, err
		}
		return MergeExtensions(jsonByteArray, value.Extensions)
	} else {
		jsonByteArray, err := json.Marshal(&OperationMessageAlias{
			SchemaFormat: value.SchemaFormat,
			ContentType:  value.ContentType,
			Headers:      value.Headers,
			Payload:      value.Payload,
			Tags:         value.Tags,
			Summary:      value.Summary,
			Name:         value.Name,
			Title:        value.Title,
			Description:  value.Description,
			ExternalDocs: value.ExternalDocs,
			Deprecated:   value.Deprecated,
			Example:      value.Example,
		})
		if err != nil {
			return nil, err
		}
		return MergeExtensions(jsonByteArray, value.Extensions)
	}
}
