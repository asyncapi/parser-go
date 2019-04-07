package openapi

import (
	"encoding/json"

	"github.com/asyncapi/parser/pkg/models"
)

// Schema maps AsyncAPI "schema" object
type Schema struct {
	Extensions           map[string]json.RawMessage `json:"-"`
	Ref                  json.RawMessage            `json:"$ref,omitempty"`
	Nullable             *bool                      `json:"nullable,omitempty"`
	Format               string                     `json:"format,omitempty"`
	Title                string                     `json:"title,omitempty"`
	Description          string                     `json:"description,omitempty"`
	Default              json.RawMessage            `json:"default,omitempty"`
	MultipleOf           float64                    `json:"multipleOf,omitempty"`
	Maximum              float64                    `json:"maximum,omitempty"`
	ExclusiveMaximum     bool                       `json:"exclusiveMaximum,omitempty"`
	Minimum              float64                    `json:"minimum,omitempty"`
	ExclusiveMinimum     bool                       `json:"exclusiveMinimum,omitempty"`
	MaxLength            int                        `json:"maxLength,omitempty"`
	MinLength            int                        `json:"minLength,omitempty"`
	Pattern              string                     `json:"pattern,omitempty"`
	MaxItems             int                        `json:"maxItems,omitempty"`
	MinItems             int                        `json:"minItems,omitempty"`
	UniqueItems          bool                       `json:"uniqueItems,omitempty"`
	MaxProperties        int                        `json:"maxProperties,omitempty"`
	MinProperties        int                        `json:"minProperties,omitempty"`
	Required             []string                   `json:"required,omitempty"`
	Enum                 json.RawMessage            `json:"enum,omitempty"`
	Deprecated           bool                       `json:"deprecated,omitempty"`
	AdditionalProperties interface{}                `json:"additionalProperties,omitempty"`
	Type                 interface{}                `json:"type,omitempty"`
	Items                json.RawMessage            `json:"items,omitempty"`
	AllOf                []*Schema                  `json:"allOf,omitempty"`
	OneOf                []*Schema                  `json:"oneOf,omitempty"`
	AnyOf                []*Schema                  `json:"anyOf,omitempty"`
	Not                  *Schema                    `json:"not,omitempty"`
	Properties           map[string]*Schema         `json:"properties,omitempty"`
	Discriminator        string                     `json:"discriminator,omitempty"`
	ReadOnly             bool                       `json:"readOnly,omitempty"`
	Xml                  json.RawMessage            `json:"xml,omitempty"`
	ExternalDocs         json.RawMessage            `json:"externalDocs,omitempty"`
	Example              json.RawMessage            `json:"example,omitempty"`
	Examples             []json.RawMessage          `json:"examples,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Schema) UnmarshalJSON(data []byte) error {
	type SchemaAlias Schema
	jsonMap := SchemaAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Ref = jsonMap.Ref
	value.Nullable = jsonMap.Nullable
	value.Format = jsonMap.Format
	value.Title = jsonMap.Title
	value.Description = jsonMap.Description
	value.Default = jsonMap.Default
	value.MultipleOf = jsonMap.MultipleOf
	value.Maximum = jsonMap.Maximum
	value.ExclusiveMaximum = jsonMap.ExclusiveMaximum
	value.Minimum = jsonMap.Minimum
	value.ExclusiveMinimum = jsonMap.ExclusiveMinimum
	value.MaxLength = jsonMap.MaxLength
	value.MinLength = jsonMap.MinLength
	value.Pattern = jsonMap.Pattern
	value.MaxItems = jsonMap.MaxItems
	value.MinItems = jsonMap.MinItems
	value.UniqueItems = jsonMap.UniqueItems
	value.MaxProperties = jsonMap.MaxProperties
	value.MinProperties = jsonMap.MinProperties
	value.Required = jsonMap.Required
	value.Enum = jsonMap.Enum
	value.Deprecated = jsonMap.Deprecated

	if jsonMap.AdditionalProperties != nil {
		sBytes, err := json.Marshal(jsonMap.AdditionalProperties)
		if err != nil {
			return err
		}

		var schBool bool
		if err := json.Unmarshal(sBytes, &schBool); err != nil {
			var schSchema Schema
			if err := json.Unmarshal(sBytes, &schSchema); err != nil {
				return err
			}

			value.AdditionalProperties = schSchema
		} else {
			value.AdditionalProperties = schBool
		}
	} else {
		value.AdditionalProperties = nil
	}

	if _, ok := jsonMap.Type.([]interface{}); ok {
		arrInterfaces := jsonMap.Type.([]interface{})
		arrStrings := make([]string, len(arrInterfaces))
		for i, v := range arrInterfaces {
			arrStrings[i] = v.(string)
		}
		value.Type = arrStrings
	}
	if _, ok := jsonMap.Type.(string); ok {
		value.Type = jsonMap.Type.(string)
	}
	value.Items = jsonMap.Items
	value.AllOf = jsonMap.AllOf
	value.OneOf = jsonMap.OneOf
	value.AnyOf = jsonMap.AnyOf
	value.Not = jsonMap.Not
	value.Properties = jsonMap.Properties
	value.Discriminator = jsonMap.Discriminator
	value.ReadOnly = jsonMap.ReadOnly
	value.Xml = jsonMap.Xml
	value.ExternalDocs = jsonMap.ExternalDocs
	value.Example = jsonMap.Example
	value.Examples = jsonMap.Examples

	exts, err := models.ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value Schema) MarshalJSON() ([]byte, error) {
	type SchemaAlias Schema
	jsonByteArray, err := json.Marshal(&SchemaAlias{
		Ref:                  value.Ref,
		Nullable:             value.Nullable,
		Format:               value.Format,
		Title:                value.Title,
		Description:          value.Description,
		Default:              value.Default,
		MultipleOf:           value.MultipleOf,
		Maximum:              value.Maximum,
		ExclusiveMaximum:     value.ExclusiveMaximum,
		Minimum:              value.Minimum,
		ExclusiveMinimum:     value.ExclusiveMinimum,
		MaxLength:            value.MaxLength,
		MinLength:            value.MinLength,
		Pattern:              value.Pattern,
		MaxItems:             value.MaxItems,
		MinItems:             value.MinItems,
		UniqueItems:          value.UniqueItems,
		MaxProperties:        value.MaxProperties,
		MinProperties:        value.MinProperties,
		Required:             value.Required,
		Enum:                 value.Enum,
		Deprecated:           value.Deprecated,
		AdditionalProperties: value.AdditionalProperties,
		Type:                 value.Type,
		Items:                value.Items,
		AllOf:                value.AllOf,
		OneOf:                value.OneOf,
		AnyOf:                value.AnyOf,
		Not:                  value.Not,
		Properties:           value.Properties,
		Discriminator:        value.Discriminator,
		ReadOnly:             value.ReadOnly,
		Xml:                  value.Xml,
		ExternalDocs:         value.ExternalDocs,
		Example:              value.Example,
		Examples:             value.Examples,
	})
	if err != nil {
		return nil, err
	}
	return models.MergeExtensions(jsonByteArray, value.Extensions)
}
