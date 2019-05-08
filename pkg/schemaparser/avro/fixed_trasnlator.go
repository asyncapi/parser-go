package avro

import (
	"encoding/json"
	"fmt"

	"github.com/asyncapi/parser/pkg/errs"
)

// FixedAvro maps fixed scheme
type FixedAvro struct {
	Type    string  `json:"type,omitempty"`
	Pattern string  `json:"pattern,omitempty"`
	Min     float64 `json:"minLength,omitempty"`
	Max     float64 `json:"maxLength,omitempty"`
}

// NewFixedAvro creates a fixed avro depending on the type
func NewFixedAvro(attrType, refKey string, size float64) FixedAvro {
	rAvro := FixedAvro{Type: "string", Pattern: byteString, Min: size, Max: size}
	return rAvro
}

// Convert transforms avro formatted message to JSONSchema
func (ra *FixedAvro) Convert(message map[string]interface{}) (string, *errs.ParserError) {
	convertedMessage := string(`{
		"definitions" : {
		  "%s": %s
		},
		"$ref" : "#/definitions/%s"
	}`)

	refKey := fmt.Sprintf("%s:%s", message["type"].(string), message["name"].(string))
	fAvro := NewFixedAvro(message["type"].(string), refKey, message["size"].(float64))

	jFAvro, _ := json.Marshal(fAvro)
	translated := fmt.Sprintf("%s", string(jFAvro))
	return fmt.Sprintf(convertedMessage, refKey, translated, refKey), nil
}
