package avro

import (
	"encoding/json"
	"fmt"

	"github.com/asyncapi/parser/pkg/errs"
)

// SimpleAvro maps a simple type in avro
type SimpleAvro struct {
	Type    string  `json:"type,omitempty"`
	Pattern string  `json:"pattern,omitempty"`
	Min     float64 `json:"minimum,omitempty"`
	Max     float64 `json:"maximum,omitempty"`
}

const (
	minLong    = -9223372036854775808
	maxLong    = 9223372036854775807
	minInt     = -2147483648
	maxInt     = 2147483647
	byteString = "^[\u0000-\u00ff]*$"
)

// NewSimpleAvro creates a simple avro depending on the type
func NewSimpleAvro(attrType string) SimpleAvro {
	var rAvro SimpleAvro
	switch attrType {
	case "null", "boolean", "float", "double", "string":
		rAvro = SimpleAvro{Type: convertType(attrType)}
	case "long":
		rAvro = SimpleAvro{Type: convertType(attrType), Min: minLong, Max: maxLong}
	case "int":
		rAvro = SimpleAvro{Type: convertType(attrType), Min: minInt, Max: maxInt}
	case "bytes":
		rAvro = SimpleAvro{Type: convertType(attrType), Pattern: byteString}
	default:
		rAvro = SimpleAvro{Type: attrType}
	}
	return rAvro
}

// Convert transforms avro formatted message to JSONSchema
func (ra *SimpleAvro) Convert(message map[string]interface{}) (string, *errs.ParserError) {
	rAvro := NewSimpleAvro(message["type"].(string))
	jRAvro, _ := json.Marshal(rAvro)
	translated := fmt.Sprintf("%s", string(jRAvro))
	return translated, nil
}

func convertType(attrType string) string {
	switch attrType {
	case "long", "int":
		return "integer"
	case "float", "double":
		return "number"
	case "bytes":
		return "string"
	default:
		return attrType
	}
}
