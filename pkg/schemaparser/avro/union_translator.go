package avro

import (
	"encoding/json"
)

// UnionAvro maps union avro scheme
type UnionAvro struct {
	OneOf       []json.RawMessage `json:"oneOf"`
	Definitions json.RawMessage   `json:"definitions,omitempty"`
}
