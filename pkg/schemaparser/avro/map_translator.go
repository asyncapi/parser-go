package avro

import (
	"encoding/json"
	"log"

	"github.com/asyncapi/parser/pkg/errs"
)

// Reference maps a ref object with its definitions
type Reference struct {
	Ref         json.RawMessage `json:"$ref"`
	Definitions json.RawMessage `json:"definitions,omitempty"`
}

// ComposeMapAvro maps a avro map object
type ComposeMapAvro struct {
	Type        string          `json:"type,omitempty"`
	Map         json.RawMessage `json:"additionalProperties,omitempty"`
	Definitions json.RawMessage `json:"definitions,omitempty"`
}

// MapAvro maps a avro map object
type MapAvro struct {
	Type                 string                   `json:"type,omitempty"`
	AdditionalProperties AdditionalPropertiesItem `json:"additionalProperties,omitempty"`
}

// NewMapAvro creates a fixed avro depending on the type
func NewMapAvro(attrType string) MapAvro {
	aP := AdditionalPropertiesItem{Type: convertType(attrType)}
	mAvro := MapAvro{Type: convertType("map"), AdditionalProperties: aP}
	return mAvro
}

// NewComposeMapAvro creates a fixed avro depending on the type
func NewComposeMapAvro(itemMap map[string]interface{}) ComposeMapAvro {
	log.Printf("Map %v", itemMap)
	var bAdditionaProperties json.RawMessage
	bAdditionaProperties, err := json.Marshal(itemMap)
	if err != nil {
		log.Fatalf("Error marshalling additionalProperties: %v", itemMap)
	}
	Parse(&bAdditionaProperties)
	var ref Reference
	json.Unmarshal(bAdditionaProperties, &ref)
	log.Printf("Reference %s", ref)
	var cMAvro ComposeMapAvro
	if ref.Ref == nil {
		mAvro := MapAvro{Type: convertType(itemMap["type"].(string)), AdditionalProperties: convertValues(itemMap["values"].(string))}
		bMAvro, _ := json.Marshal(mAvro)
		cMAvro = ComposeMapAvro{Type: convertType("map"), Map: bMAvro}
	} else {
		cMAvro = ComposeMapAvro{Type: convertType("map"), Map: ref.Ref, Definitions: ref.Definitions}
	}
	return cMAvro
}

// Convert transforms avro formatted message to JSONSchema
func (ra *MapAvro) Convert(message map[string]interface{}) (string, *errs.ParserError) {
	switch message["values"].(type) {
	// Simple objects
	case string:
		mAvro := NewMapAvro(message["values"].(string))
		jMAvro, _ := json.Marshal(mAvro)
		return string(jMAvro), nil
		// Complex objects
	case map[string]interface{}:
		log.Printf("Map")
		itemMap := message["values"].(map[string]interface{})
		mAvro := NewComposeMapAvro(itemMap)
		jMAvro, _ := json.Marshal(mAvro)
		return string(jMAvro), nil
	}
	return "", errs.New("Can't convert Map object")
}
