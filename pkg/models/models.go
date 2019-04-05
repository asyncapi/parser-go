package models

import "encoding/json"

// AsyncapiDocument maps AsyncAPI "AsyncapiDocument" object
type AsyncapiDocument struct {
	Extensions         map[string]json.RawMessage `json:"-"`
	Asyncapi           string                     `json:"asyncapi,omitempty"`
	Id                 string                     `json:"id,omitempty"`
	Info               Info                       `json:"info,omitempty"`
	Servers            []*Server                  `json:"servers,omitempty"`
	DefaultContentType string                     `json:"defaultContentType,omitempty"`
	Channels           Channels                   `json:"channels,omitempty"`
	Components         *Components                `json:"components,omitempty"`
	Tags               []*Tag                     `json:"tags,omitempty"`
	ExternalDocs       *ExternalDocs              `json:"externalDocs,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *AsyncapiDocument) UnmarshalJSON(data []byte) error {
	type AsyncapiDocumentAlias AsyncapiDocument
	jsonMap := AsyncapiDocumentAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Asyncapi = jsonMap.Asyncapi
	value.Id = jsonMap.Id
	value.Info = jsonMap.Info
	value.Servers = jsonMap.Servers
	value.DefaultContentType = jsonMap.DefaultContentType
	value.Channels = jsonMap.Channels
	value.Components = jsonMap.Components
	value.Tags = jsonMap.Tags
	value.ExternalDocs = jsonMap.ExternalDocs

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value AsyncapiDocument) MarshalJSON() ([]byte, error) {
	type AsyncapiDocumentAlias AsyncapiDocument
	jsonByteArray, err := json.Marshal(&AsyncapiDocumentAlias{
		Asyncapi:           value.Asyncapi,
		Id:                 value.Id,
		Info:               value.Info,
		Servers:            value.Servers,
		DefaultContentType: value.DefaultContentType,
		Channels:           value.Channels,
		Components:         value.Components,
		Tags:               value.Tags,
		ExternalDocs:       value.ExternalDocs,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// Reference maps AsyncAPI "Reference" object
type Reference struct {
	Ref ReferenceObject `json:"$ref,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Reference) UnmarshalJSON(data []byte) error {
	type ReferenceAlias Reference
	jsonMap := ReferenceAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Ref = jsonMap.Ref

	return nil
}

// MarshalJSON marshals JSON
func (value Reference) MarshalJSON() ([]byte, error) {
	type ReferenceAlias Reference
	jsonByteArray, err := json.Marshal(&ReferenceAlias{
		Ref: value.Ref,
	})
	if err != nil {
		return nil, err
	}
	return jsonByteArray, nil
}

// ReferenceObject maps AsyncAPI "ReferenceObject" object
type ReferenceObject struct {
}

// UnmarshalJSON unmarshals JSON
func (value *ReferenceObject) UnmarshalJSON(data []byte) error {
	type ReferenceObjectAlias ReferenceObject
	jsonMap := ReferenceObjectAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	return nil
}

// MarshalJSON marshals JSON
func (value ReferenceObject) MarshalJSON() ([]byte, error) {
	type ReferenceObjectAlias ReferenceObject
	jsonByteArray, err := json.Marshal(&ReferenceObjectAlias{})
	if err != nil {
		return nil, err
	}
	return jsonByteArray, nil
}

// Info maps AsyncAPI "info" object
type Info struct {
	Extensions     map[string]json.RawMessage `json:"-"`
	Title          string                     `json:"title,omitempty"`
	Version        string                     `json:"version,omitempty"`
	Description    string                     `json:"description,omitempty"`
	TermsOfService string                     `json:"termsOfService,omitempty"`
	Contact        *Contact                   `json:"contact,omitempty"`
	License        *License                   `json:"license,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Info) UnmarshalJSON(data []byte) error {
	type InfoAlias Info
	jsonMap := InfoAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Title = jsonMap.Title
	value.Version = jsonMap.Version
	value.Description = jsonMap.Description
	value.TermsOfService = jsonMap.TermsOfService
	value.Contact = jsonMap.Contact
	value.License = jsonMap.License

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value Info) MarshalJSON() ([]byte, error) {
	type InfoAlias Info
	jsonByteArray, err := json.Marshal(&InfoAlias{
		Title:          value.Title,
		Version:        value.Version,
		Description:    value.Description,
		TermsOfService: value.TermsOfService,
		Contact:        value.Contact,
		License:        value.License,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// Contact maps AsyncAPI "contact" object
type Contact struct {
	Extensions map[string]json.RawMessage `json:"-"`
	Name       string                     `json:"name,omitempty"`
	Url        string                     `json:"url,omitempty"`
	Email      string                     `json:"email,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Contact) UnmarshalJSON(data []byte) error {
	type ContactAlias Contact
	jsonMap := ContactAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Name = jsonMap.Name
	value.Url = jsonMap.Url
	value.Email = jsonMap.Email

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value Contact) MarshalJSON() ([]byte, error) {
	type ContactAlias Contact
	jsonByteArray, err := json.Marshal(&ContactAlias{
		Name:  value.Name,
		Url:   value.Url,
		Email: value.Email,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// License maps AsyncAPI "license" object
type License struct {
	Extensions map[string]json.RawMessage `json:"-"`
	Name       string                     `json:"name,omitempty"`
	Url        string                     `json:"url,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *License) UnmarshalJSON(data []byte) error {
	type LicenseAlias License
	jsonMap := LicenseAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Name = jsonMap.Name
	value.Url = jsonMap.Url

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value License) MarshalJSON() ([]byte, error) {
	type LicenseAlias License
	jsonByteArray, err := json.Marshal(&LicenseAlias{
		Name: value.Name,
		Url:  value.Url,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// Server maps AsyncAPI "server" object
type Server struct {
	Extensions      map[string]json.RawMessage `json:"-"`
	Url             string                     `json:"url,omitempty"`
	Description     string                     `json:"description,omitempty"`
	Protocol        string                     `json:"protocol,omitempty"`
	ProtocolVersion string                     `json:"protocolVersion,omitempty"`
	Variables       *ServerVariables           `json:"variables,omitempty"`
	BaseChannel     string                     `json:"baseChannel,omitempty"`
	Security        []*SecurityRequirement     `json:"security,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Server) UnmarshalJSON(data []byte) error {
	type ServerAlias Server
	jsonMap := ServerAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Url = jsonMap.Url
	value.Description = jsonMap.Description
	value.Protocol = jsonMap.Protocol
	value.ProtocolVersion = jsonMap.ProtocolVersion
	value.Variables = jsonMap.Variables
	value.BaseChannel = jsonMap.BaseChannel
	value.Security = jsonMap.Security

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value Server) MarshalJSON() ([]byte, error) {
	type ServerAlias Server
	jsonByteArray, err := json.Marshal(&ServerAlias{
		Url:             value.Url,
		Description:     value.Description,
		Protocol:        value.Protocol,
		ProtocolVersion: value.ProtocolVersion,
		Variables:       value.Variables,
		BaseChannel:     value.BaseChannel,
		Security:        value.Security,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// ServerVariables maps AsyncAPI "serverVariables" object
type ServerVariables map[string]*ServerVariable

// ServerVariable maps AsyncAPI "serverVariable" object
type ServerVariable struct {
	Extensions  map[string]json.RawMessage `json:"-"`
	Enum        []string                   `json:"enum,omitempty"`
	Default     string                     `json:"default,omitempty"`
	Description string                     `json:"description,omitempty"`
	Examples    []string                   `json:"examples,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *ServerVariable) UnmarshalJSON(data []byte) error {
	type ServerVariableAlias ServerVariable
	jsonMap := ServerVariableAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Enum = jsonMap.Enum
	value.Default = jsonMap.Default
	value.Description = jsonMap.Description
	value.Examples = jsonMap.Examples

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value ServerVariable) MarshalJSON() ([]byte, error) {
	type ServerVariableAlias ServerVariable
	jsonByteArray, err := json.Marshal(&ServerVariableAlias{
		Enum:        value.Enum,
		Default:     value.Default,
		Description: value.Description,
		Examples:    value.Examples,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// Channels maps AsyncAPI "channels" object
type Channels map[string]*ChannelItem

// Components maps AsyncAPI "components" object
type Components struct {
	Schemas         *Schemas        `json:"schemas,omitempty"`
	Messages        *Messages       `json:"messages,omitempty"`
	SecuritySchemes json.RawMessage `json:"securitySchemes,omitempty"`
	Parameters      *Parameters     `json:"parameters,omitempty"`
	CorrelationIds  json.RawMessage `json:"correlationIds,omitempty"`
	Traits          *Traits         `json:"traits,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Components) UnmarshalJSON(data []byte) error {
	type ComponentsAlias Components
	jsonMap := ComponentsAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Schemas = jsonMap.Schemas
	value.Messages = jsonMap.Messages
	value.SecuritySchemes = jsonMap.SecuritySchemes
	value.Parameters = jsonMap.Parameters
	value.CorrelationIds = jsonMap.CorrelationIds
	value.Traits = jsonMap.Traits

	return nil
}

// MarshalJSON marshals JSON
func (value Components) MarshalJSON() ([]byte, error) {
	type ComponentsAlias Components
	jsonByteArray, err := json.Marshal(&ComponentsAlias{
		Schemas:         value.Schemas,
		Messages:        value.Messages,
		SecuritySchemes: value.SecuritySchemes,
		Parameters:      value.Parameters,
		CorrelationIds:  value.CorrelationIds,
		Traits:          value.Traits,
	})
	if err != nil {
		return nil, err
	}
	return jsonByteArray, nil
}

// Schemas maps AsyncAPI "schemas" object
type Schemas map[string]*Schema

// Messages maps AsyncAPI "messages" object
type Messages map[string]*Message

// Parameters maps AsyncAPI "parameters" object
type Parameters map[string]*Parameter

// Schema maps AsyncAPI "schema" object
type Schema struct {
	Extensions           map[string]json.RawMessage `json:"-"`
	Ref                  *ReferenceObject           `json:"$ref,omitempty"`
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
	AdditionalProperties json.RawMessage            `json:"additionalProperties,omitempty"`
	Type                 string                     `json:"type,omitempty"`
	Items                json.RawMessage            `json:"items,omitempty"`
	AllOf                []*Schema                  `json:"allOf,omitempty"`
	OneOf                []*Schema                  `json:"oneOf,omitempty"`
	AnyOf                []*Schema                  `json:"anyOf,omitempty"`
	Not                  *Schema                    `json:"not,omitempty"`
	Properties           json.RawMessage            `json:"properties,omitempty"`
	Discriminator        string                     `json:"discriminator,omitempty"`
	ReadOnly             bool                       `json:"readOnly,omitempty"`
	Xml                  *Xml                       `json:"xml,omitempty"`
	ExternalDocs         *ExternalDocs              `json:"externalDocs,omitempty"`
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
	value.AdditionalProperties = jsonMap.AdditionalProperties
	value.Type = jsonMap.Type
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

	exts, err := ExtensionsFromJSON(data)
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
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// Xml maps AsyncAPI "xml" object
type Xml struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Prefix    string `json:"prefix,omitempty"`
	Attribute bool   `json:"attribute,omitempty"`
	Wrapped   bool   `json:"wrapped,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Xml) UnmarshalJSON(data []byte) error {
	type XmlAlias Xml
	jsonMap := XmlAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Name = jsonMap.Name
	value.Namespace = jsonMap.Namespace
	value.Prefix = jsonMap.Prefix
	value.Attribute = jsonMap.Attribute
	value.Wrapped = jsonMap.Wrapped

	return nil
}

// MarshalJSON marshals JSON
func (value Xml) MarshalJSON() ([]byte, error) {
	type XmlAlias Xml
	jsonByteArray, err := json.Marshal(&XmlAlias{
		Name:      value.Name,
		Namespace: value.Namespace,
		Prefix:    value.Prefix,
		Attribute: value.Attribute,
		Wrapped:   value.Wrapped,
	})
	if err != nil {
		return nil, err
	}
	return jsonByteArray, nil
}

// ExternalDocs maps AsyncAPI "externalDocs" object
type ExternalDocs struct {
	Extensions  map[string]json.RawMessage `json:"-"`
	Description string                     `json:"description,omitempty"`
	Url         string                     `json:"url,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *ExternalDocs) UnmarshalJSON(data []byte) error {
	type ExternalDocsAlias ExternalDocs
	jsonMap := ExternalDocsAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Description = jsonMap.Description
	value.Url = jsonMap.Url

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value ExternalDocs) MarshalJSON() ([]byte, error) {
	type ExternalDocsAlias ExternalDocs
	jsonByteArray, err := json.Marshal(&ExternalDocsAlias{
		Description: value.Description,
		Url:         value.Url,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// ChannelItem maps AsyncAPI "channelItem" object
type ChannelItem struct {
	Extensions   map[string]json.RawMessage `json:"-"`
	Ref          *ReferenceObject           `json:"$ref,omitempty"`
	Parameters   []*Parameter               `json:"parameters,omitempty"`
	Publish      *Operation                 `json:"publish,omitempty"`
	Subscribe    *Operation                 `json:"subscribe,omitempty"`
	Deprecated   bool                       `json:"deprecated,omitempty"`
	ProtocolInfo json.RawMessage            `json:"protocolInfo,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *ChannelItem) UnmarshalJSON(data []byte) error {
	type ChannelItemAlias ChannelItem
	jsonMap := ChannelItemAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Ref = jsonMap.Ref
	value.Parameters = jsonMap.Parameters
	value.Publish = jsonMap.Publish
	value.Subscribe = jsonMap.Subscribe
	value.Deprecated = jsonMap.Deprecated
	value.ProtocolInfo = jsonMap.ProtocolInfo

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value ChannelItem) MarshalJSON() ([]byte, error) {
	type ChannelItemAlias ChannelItem
	jsonByteArray, err := json.Marshal(&ChannelItemAlias{
		Ref:          value.Ref,
		Parameters:   value.Parameters,
		Publish:      value.Publish,
		Subscribe:    value.Subscribe,
		Deprecated:   value.Deprecated,
		ProtocolInfo: value.ProtocolInfo,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// Parameter maps AsyncAPI "parameter" object
type Parameter struct {
	Extensions  map[string]json.RawMessage `json:"-"`
	Description string                     `json:"description,omitempty"`
	Name        string                     `json:"name,omitempty"`
	Schema      *Schema                    `json:"schema,omitempty"`
	Ref         *ReferenceObject           `json:"$ref,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Parameter) UnmarshalJSON(data []byte) error {
	type ParameterAlias Parameter
	jsonMap := ParameterAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Description = jsonMap.Description
	value.Name = jsonMap.Name
	value.Schema = jsonMap.Schema
	value.Ref = jsonMap.Ref

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value Parameter) MarshalJSON() ([]byte, error) {
	type ParameterAlias Parameter
	jsonByteArray, err := json.Marshal(&ParameterAlias{
		Description: value.Description,
		Name:        value.Name,
		Schema:      value.Schema,
		Ref:         value.Ref,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// Operation maps AsyncAPI "operation" object
type Operation struct {
	Extensions   map[string]json.RawMessage `json:"-"`
	Traits       []json.RawMessage          `json:"traits,omitempty"`
	Summary      string                     `json:"summary,omitempty"`
	Description  string                     `json:"description,omitempty"`
	Tags         []*Tag                     `json:"tags,omitempty"`
	ExternalDocs *ExternalDocs              `json:"externalDocs,omitempty"`
	OperationId  string                     `json:"operationId,omitempty"`
	ProtocolInfo json.RawMessage            `json:"protocolInfo,omitempty"`
	Message      *Message                   `json:"message,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Operation) UnmarshalJSON(data []byte) error {
	type OperationAlias Operation
	jsonMap := OperationAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Traits = jsonMap.Traits
	value.Summary = jsonMap.Summary
	value.Description = jsonMap.Description
	value.Tags = jsonMap.Tags
	value.ExternalDocs = jsonMap.ExternalDocs
	value.OperationId = jsonMap.OperationId
	value.ProtocolInfo = jsonMap.ProtocolInfo
	value.Message = jsonMap.Message

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value Operation) MarshalJSON() ([]byte, error) {
	type OperationAlias Operation
	jsonByteArray, err := json.Marshal(&OperationAlias{
		Traits:       value.Traits,
		Summary:      value.Summary,
		Description:  value.Description,
		Tags:         value.Tags,
		ExternalDocs: value.ExternalDocs,
		OperationId:  value.OperationId,
		ProtocolInfo: value.ProtocolInfo,
		Message:      value.Message,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// Message maps AsyncAPI "message" object
type Message struct {
	Extensions    map[string]json.RawMessage `json:"-"`
	OneOf         []*Message                 `json:"oneOf,omitempty"`
	SchemaFormat  string                     `json:"schemaFormat,omitempty"`
	ContentType   string                     `json:"contentType,omitempty"`
	Headers       json.RawMessage            `json:"headers,omitempty"`
	Payload       json.RawMessage            `json:"payload,omitempty"`
	CorrelationId json.RawMessage            `json:"correlationId,omitempty"`
	Tags          []*Tag                     `json:"tags,omitempty"`
	Summary       string                     `json:"summary,omitempty"`
	Name          string                     `json:"name,omitempty"`
	Title         string                     `json:"title,omitempty"`
	Description   string                     `json:"description,omitempty"`
	ExternalDocs  *ExternalDocs              `json:"externalDocs,omitempty"`
	Deprecated    bool                       `json:"deprecated,omitempty"`
	Examples      []json.RawMessage          `json:"examples,omitempty"`
	ProtocolInfo  json.RawMessage            `json:"protocolInfo,omitempty"`
	Traits        []json.RawMessage          `json:"traits,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Message) UnmarshalJSON(data []byte) error {
	type MessageAlias Message
	jsonMap := MessageAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.OneOf = jsonMap.OneOf
	value.SchemaFormat = jsonMap.SchemaFormat
	value.ContentType = jsonMap.ContentType
	value.Headers = jsonMap.Headers
	value.Payload = jsonMap.Payload
	value.CorrelationId = jsonMap.CorrelationId
	value.Tags = jsonMap.Tags
	value.Summary = jsonMap.Summary
	value.Name = jsonMap.Name
	value.Title = jsonMap.Title
	value.Description = jsonMap.Description
	value.ExternalDocs = jsonMap.ExternalDocs
	value.Deprecated = jsonMap.Deprecated
	value.Examples = jsonMap.Examples
	value.ProtocolInfo = jsonMap.ProtocolInfo
	value.Traits = jsonMap.Traits

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value Message) MarshalJSON() ([]byte, error) {
	type MessageAlias Message
	jsonByteArray, err := json.Marshal(&MessageAlias{
		OneOf:         value.OneOf,
		SchemaFormat:  value.SchemaFormat,
		ContentType:   value.ContentType,
		Headers:       value.Headers,
		Payload:       value.Payload,
		CorrelationId: value.CorrelationId,
		Tags:          value.Tags,
		Summary:       value.Summary,
		Name:          value.Name,
		Title:         value.Title,
		Description:   value.Description,
		ExternalDocs:  value.ExternalDocs,
		Deprecated:    value.Deprecated,
		Examples:      value.Examples,
		ProtocolInfo:  value.ProtocolInfo,
		Traits:        value.Traits,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// CorrelationId maps AsyncAPI "correlationId" object
type CorrelationId struct {
	Extensions  map[string]json.RawMessage `json:"-"`
	Description string                     `json:"description,omitempty"`
	Location    string                     `json:"location,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *CorrelationId) UnmarshalJSON(data []byte) error {
	type CorrelationIdAlias CorrelationId
	jsonMap := CorrelationIdAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Description = jsonMap.Description
	value.Location = jsonMap.Location

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value CorrelationId) MarshalJSON() ([]byte, error) {
	type CorrelationIdAlias CorrelationId
	jsonByteArray, err := json.Marshal(&CorrelationIdAlias{
		Description: value.Description,
		Location:    value.Location,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// SpecificationExtension maps AsyncAPI "specificationExtension" object
type SpecificationExtension map[string]json.RawMessage

// Tag maps AsyncAPI "tag" object
type Tag struct {
	Extensions   map[string]json.RawMessage `json:"-"`
	Name         string                     `json:"name,omitempty"`
	Description  string                     `json:"description,omitempty"`
	ExternalDocs *ExternalDocs              `json:"externalDocs,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Tag) UnmarshalJSON(data []byte) error {
	type TagAlias Tag
	jsonMap := TagAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Name = jsonMap.Name
	value.Description = jsonMap.Description
	value.ExternalDocs = jsonMap.ExternalDocs

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value Tag) MarshalJSON() ([]byte, error) {
	type TagAlias Tag
	jsonByteArray, err := json.Marshal(&TagAlias{
		Name:         value.Name,
		Description:  value.Description,
		ExternalDocs: value.ExternalDocs,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// Traits maps AsyncAPI "traits" object
type Traits map[string]json.RawMessage

// OperationTrait maps AsyncAPI "operationTrait" object
type OperationTrait struct {
	Extensions   map[string]json.RawMessage `json:"-"`
	Summary      string                     `json:"summary,omitempty"`
	Description  string                     `json:"description,omitempty"`
	Tags         []*Tag                     `json:"tags,omitempty"`
	ExternalDocs *ExternalDocs              `json:"externalDocs,omitempty"`
	OperationId  string                     `json:"operationId,omitempty"`
	ProtocolInfo json.RawMessage            `json:"protocolInfo,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *OperationTrait) UnmarshalJSON(data []byte) error {
	type OperationTraitAlias OperationTrait
	jsonMap := OperationTraitAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Summary = jsonMap.Summary
	value.Description = jsonMap.Description
	value.Tags = jsonMap.Tags
	value.ExternalDocs = jsonMap.ExternalDocs
	value.OperationId = jsonMap.OperationId
	value.ProtocolInfo = jsonMap.ProtocolInfo

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value OperationTrait) MarshalJSON() ([]byte, error) {
	type OperationTraitAlias OperationTrait
	jsonByteArray, err := json.Marshal(&OperationTraitAlias{
		Summary:      value.Summary,
		Description:  value.Description,
		Tags:         value.Tags,
		ExternalDocs: value.ExternalDocs,
		OperationId:  value.OperationId,
		ProtocolInfo: value.ProtocolInfo,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// MessageTrait maps AsyncAPI "messageTrait" object
type MessageTrait struct {
	Extensions    map[string]json.RawMessage `json:"-"`
	SchemaFormat  string                     `json:"schemaFormat,omitempty"`
	ContentType   string                     `json:"contentType,omitempty"`
	Headers       json.RawMessage            `json:"headers,omitempty"`
	CorrelationId json.RawMessage            `json:"correlationId,omitempty"`
	Tags          []*Tag                     `json:"tags,omitempty"`
	Summary       string                     `json:"summary,omitempty"`
	Name          string                     `json:"name,omitempty"`
	Title         string                     `json:"title,omitempty"`
	Description   string                     `json:"description,omitempty"`
	ExternalDocs  *ExternalDocs              `json:"externalDocs,omitempty"`
	Deprecated    bool                       `json:"deprecated,omitempty"`
	Examples      []json.RawMessage          `json:"examples,omitempty"`
	ProtocolInfo  json.RawMessage            `json:"protocolInfo,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *MessageTrait) UnmarshalJSON(data []byte) error {
	type MessageTraitAlias MessageTrait
	jsonMap := MessageTraitAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.SchemaFormat = jsonMap.SchemaFormat
	value.ContentType = jsonMap.ContentType
	value.Headers = jsonMap.Headers
	value.CorrelationId = jsonMap.CorrelationId
	value.Tags = jsonMap.Tags
	value.Summary = jsonMap.Summary
	value.Name = jsonMap.Name
	value.Title = jsonMap.Title
	value.Description = jsonMap.Description
	value.ExternalDocs = jsonMap.ExternalDocs
	value.Deprecated = jsonMap.Deprecated
	value.Examples = jsonMap.Examples
	value.ProtocolInfo = jsonMap.ProtocolInfo

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value MessageTrait) MarshalJSON() ([]byte, error) {
	type MessageTraitAlias MessageTrait
	jsonByteArray, err := json.Marshal(&MessageTraitAlias{
		SchemaFormat:  value.SchemaFormat,
		ContentType:   value.ContentType,
		Headers:       value.Headers,
		CorrelationId: value.CorrelationId,
		Tags:          value.Tags,
		Summary:       value.Summary,
		Name:          value.Name,
		Title:         value.Title,
		Description:   value.Description,
		ExternalDocs:  value.ExternalDocs,
		Deprecated:    value.Deprecated,
		Examples:      value.Examples,
		ProtocolInfo:  value.ProtocolInfo,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// UserPassword maps AsyncAPI "userPassword" object
type UserPassword struct {
	Extensions  map[string]json.RawMessage `json:"-"`
	Type        string                     `json:"type,omitempty"`
	Description string                     `json:"description,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *UserPassword) UnmarshalJSON(data []byte) error {
	type UserPasswordAlias UserPassword
	jsonMap := UserPasswordAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Type = jsonMap.Type
	value.Description = jsonMap.Description

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value UserPassword) MarshalJSON() ([]byte, error) {
	type UserPasswordAlias UserPassword
	jsonByteArray, err := json.Marshal(&UserPasswordAlias{
		Type:        value.Type,
		Description: value.Description,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// ApiKey maps AsyncAPI "apiKey" object
type ApiKey struct {
	Extensions  map[string]json.RawMessage `json:"-"`
	Type        string                     `json:"type,omitempty"`
	In          string                     `json:"in,omitempty"`
	Description string                     `json:"description,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *ApiKey) UnmarshalJSON(data []byte) error {
	type ApiKeyAlias ApiKey
	jsonMap := ApiKeyAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Type = jsonMap.Type
	value.In = jsonMap.In
	value.Description = jsonMap.Description

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value ApiKey) MarshalJSON() ([]byte, error) {
	type ApiKeyAlias ApiKey
	jsonByteArray, err := json.Marshal(&ApiKeyAlias{
		Type:        value.Type,
		In:          value.In,
		Description: value.Description,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// X509 maps AsyncAPI "X509" object
type X509 struct {
	Extensions  map[string]json.RawMessage `json:"-"`
	Type        string                     `json:"type,omitempty"`
	Description string                     `json:"description,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *X509) UnmarshalJSON(data []byte) error {
	type X509Alias X509
	jsonMap := X509Alias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Type = jsonMap.Type
	value.Description = jsonMap.Description

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value X509) MarshalJSON() ([]byte, error) {
	type X509Alias X509
	jsonByteArray, err := json.Marshal(&X509Alias{
		Type:        value.Type,
		Description: value.Description,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// SymmetricEncryption maps AsyncAPI "symmetricEncryption" object
type SymmetricEncryption struct {
	Extensions  map[string]json.RawMessage `json:"-"`
	Type        string                     `json:"type,omitempty"`
	Description string                     `json:"description,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *SymmetricEncryption) UnmarshalJSON(data []byte) error {
	type SymmetricEncryptionAlias SymmetricEncryption
	jsonMap := SymmetricEncryptionAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Type = jsonMap.Type
	value.Description = jsonMap.Description

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value SymmetricEncryption) MarshalJSON() ([]byte, error) {
	type SymmetricEncryptionAlias SymmetricEncryption
	jsonByteArray, err := json.Marshal(&SymmetricEncryptionAlias{
		Type:        value.Type,
		Description: value.Description,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// AsymmetricEncryption maps AsyncAPI "asymmetricEncryption" object
type AsymmetricEncryption struct {
	Extensions  map[string]json.RawMessage `json:"-"`
	Type        string                     `json:"type,omitempty"`
	Description string                     `json:"description,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *AsymmetricEncryption) UnmarshalJSON(data []byte) error {
	type AsymmetricEncryptionAlias AsymmetricEncryption
	jsonMap := AsymmetricEncryptionAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Type = jsonMap.Type
	value.Description = jsonMap.Description

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value AsymmetricEncryption) MarshalJSON() ([]byte, error) {
	type AsymmetricEncryptionAlias AsymmetricEncryption
	jsonByteArray, err := json.Marshal(&AsymmetricEncryptionAlias{
		Type:        value.Type,
		Description: value.Description,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// NonBearerHttpSecurityScheme maps AsyncAPI "NonBearerHTTPSecurityScheme" object
type NonBearerHttpSecurityScheme struct {
	Extensions  map[string]json.RawMessage `json:"-"`
	Scheme      string                     `json:"scheme,omitempty"`
	Description string                     `json:"description,omitempty"`
	Type        string                     `json:"type,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *NonBearerHttpSecurityScheme) UnmarshalJSON(data []byte) error {
	type NonBearerHttpSecuritySchemeAlias NonBearerHttpSecurityScheme
	jsonMap := NonBearerHttpSecuritySchemeAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Scheme = jsonMap.Scheme
	value.Description = jsonMap.Description
	value.Type = jsonMap.Type

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value NonBearerHttpSecurityScheme) MarshalJSON() ([]byte, error) {
	type NonBearerHttpSecuritySchemeAlias NonBearerHttpSecurityScheme
	jsonByteArray, err := json.Marshal(&NonBearerHttpSecuritySchemeAlias{
		Scheme:      value.Scheme,
		Description: value.Description,
		Type:        value.Type,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// BearerHttpSecurityScheme maps AsyncAPI "BearerHTTPSecurityScheme" object
type BearerHttpSecurityScheme struct {
	Extensions   map[string]json.RawMessage `json:"-"`
	Scheme       string                     `json:"scheme,omitempty"`
	BearerFormat string                     `json:"bearerFormat,omitempty"`
	Type         string                     `json:"type,omitempty"`
	Description  string                     `json:"description,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *BearerHttpSecurityScheme) UnmarshalJSON(data []byte) error {
	type BearerHttpSecuritySchemeAlias BearerHttpSecurityScheme
	jsonMap := BearerHttpSecuritySchemeAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Scheme = jsonMap.Scheme
	value.BearerFormat = jsonMap.BearerFormat
	value.Type = jsonMap.Type
	value.Description = jsonMap.Description

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value BearerHttpSecurityScheme) MarshalJSON() ([]byte, error) {
	type BearerHttpSecuritySchemeAlias BearerHttpSecurityScheme
	jsonByteArray, err := json.Marshal(&BearerHttpSecuritySchemeAlias{
		Scheme:       value.Scheme,
		BearerFormat: value.BearerFormat,
		Type:         value.Type,
		Description:  value.Description,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// ApiKeyHttpSecurityScheme maps AsyncAPI "APIKeyHTTPSecurityScheme" object
type ApiKeyHttpSecurityScheme struct {
	Extensions  map[string]json.RawMessage `json:"-"`
	Type        string                     `json:"type,omitempty"`
	Name        string                     `json:"name,omitempty"`
	In          string                     `json:"in,omitempty"`
	Description string                     `json:"description,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *ApiKeyHttpSecurityScheme) UnmarshalJSON(data []byte) error {
	type ApiKeyHttpSecuritySchemeAlias ApiKeyHttpSecurityScheme
	jsonMap := ApiKeyHttpSecuritySchemeAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Type = jsonMap.Type
	value.Name = jsonMap.Name
	value.In = jsonMap.In
	value.Description = jsonMap.Description

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value ApiKeyHttpSecurityScheme) MarshalJSON() ([]byte, error) {
	type ApiKeyHttpSecuritySchemeAlias ApiKeyHttpSecurityScheme
	jsonByteArray, err := json.Marshal(&ApiKeyHttpSecuritySchemeAlias{
		Type:        value.Type,
		Name:        value.Name,
		In:          value.In,
		Description: value.Description,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// Oauth2Flows maps AsyncAPI "oauth2Flows" object
type Oauth2Flows struct {
	Extensions  map[string]json.RawMessage `json:"-"`
	Type        string                     `json:"type,omitempty"`
	Description string                     `json:"description,omitempty"`
	Flows       json.RawMessage            `json:"flows,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Oauth2Flows) UnmarshalJSON(data []byte) error {
	type Oauth2FlowsAlias Oauth2Flows
	jsonMap := Oauth2FlowsAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Type = jsonMap.Type
	value.Description = jsonMap.Description
	value.Flows = jsonMap.Flows

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value Oauth2Flows) MarshalJSON() ([]byte, error) {
	type Oauth2FlowsAlias Oauth2Flows
	jsonByteArray, err := json.Marshal(&Oauth2FlowsAlias{
		Type:        value.Type,
		Description: value.Description,
		Flows:       value.Flows,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// Oauth2Flow maps AsyncAPI "oauth2Flow" object
type Oauth2Flow struct {
	Extensions       map[string]json.RawMessage `json:"-"`
	AuthorizationUrl string                     `json:"authorizationUrl,omitempty"`
	TokenUrl         string                     `json:"tokenUrl,omitempty"`
	RefreshUrl       string                     `json:"refreshUrl,omitempty"`
	Scopes           *Oauth2Scopes              `json:"scopes,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *Oauth2Flow) UnmarshalJSON(data []byte) error {
	type Oauth2FlowAlias Oauth2Flow
	jsonMap := Oauth2FlowAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.AuthorizationUrl = jsonMap.AuthorizationUrl
	value.TokenUrl = jsonMap.TokenUrl
	value.RefreshUrl = jsonMap.RefreshUrl
	value.Scopes = jsonMap.Scopes

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value Oauth2Flow) MarshalJSON() ([]byte, error) {
	type Oauth2FlowAlias Oauth2Flow
	jsonByteArray, err := json.Marshal(&Oauth2FlowAlias{
		AuthorizationUrl: value.AuthorizationUrl,
		TokenUrl:         value.TokenUrl,
		RefreshUrl:       value.RefreshUrl,
		Scopes:           value.Scopes,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// Oauth2Scopes maps AsyncAPI "oauth2Scopes" object
type Oauth2Scopes map[string]string

// OpenIdConnect maps AsyncAPI "openIdConnect" object
type OpenIdConnect struct {
	Extensions       map[string]json.RawMessage `json:"-"`
	Type             string                     `json:"type,omitempty"`
	Description      string                     `json:"description,omitempty"`
	OpenIdConnectUrl string                     `json:"openIdConnectUrl,omitempty"`
}

// UnmarshalJSON unmarshals JSON
func (value *OpenIdConnect) UnmarshalJSON(data []byte) error {
	type OpenIdConnectAlias OpenIdConnect
	jsonMap := OpenIdConnectAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Type = jsonMap.Type
	value.Description = jsonMap.Description
	value.OpenIdConnectUrl = jsonMap.OpenIdConnectUrl

	exts, err := ExtensionsFromJSON(data)
	if err != nil {
		return err
	}
	value.Extensions = exts

	return nil
}

// MarshalJSON marshals JSON
func (value OpenIdConnect) MarshalJSON() ([]byte, error) {
	type OpenIdConnectAlias OpenIdConnect
	jsonByteArray, err := json.Marshal(&OpenIdConnectAlias{
		Type:             value.Type,
		Description:      value.Description,
		OpenIdConnectUrl: value.OpenIdConnectUrl,
	})
	if err != nil {
		return nil, err
	}
	return MergeExtensions(jsonByteArray, value.Extensions)
}

// SecurityRequirement maps AsyncAPI "SecurityRequirement" object
type SecurityRequirement map[string][]string
