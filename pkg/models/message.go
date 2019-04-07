package models

import "encoding/json"

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
