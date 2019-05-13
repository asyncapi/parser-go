package avro

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/asyncapi/parser/pkg/errs"
)

// PropertyItem is a single property in json schema
type PropertyItem struct {
	Name    string `json:"name,omitempty"`
	Type    string `json:"type,omitempty"`
	Minimum int    `json:"minimum,omitempty"`
	Maximum int    `json:"maximum,omitempty"`
}

// RecordAvro maps record type object
type RecordAvro struct {
	Type                 string                  `json:"type"`
	AdditionalProperties bool                    `json:"additionalProperties"`
	Required             []string                `json:"required"`
	Props                map[string]PropertyItem `json:"properties"`
	Definitions          json.RawMessage         `json:"definitions,omitempty"`
}

// Convert transforms avro formatted message to JSONSchema
func (ra *RecordAvro) Convert(message map[string]interface{}) (string, *errs.ParserError) {
	convertedMessage := string(`{
		"definitions" : {
		  "%s": %s
		},
		"$ref" : "#/definitions/%s"
	}`)
	var refKey, translated string
	refKey = fmt.Sprintf("%s:%s:%s", message["type"].(string), message["namespace"].(string), message["name"].(string))
	var propSlice = make(map[string]PropertyItem)
	var required []string
	objectArray := message["fields"].([]interface{})
	for _, o := range objectArray {
		switch o.(type) {
		// Complex objects
		case map[string]interface{}:
			// log.Printf("Map")
			itemMap := o.(map[string]interface{})
			propertyItem := PropertyItem{Type: convertType(itemMap["type"].(string))}
			propSlice[itemMap["name"].(string)] = propertyItem
			required = append(required, itemMap["name"].(string))
		// Simple objects
		case string:
			// log.Printf("String")
			itemMap := o.(string)
			log.Printf("String %s", itemMap)
		default:
			log.Printf("I don't know about type %T!\n", o)
			return "", errs.New(fmt.Sprintf("I don't know about type %T!\n", o))
		}

	}
	rAvro := RecordAvro{Type: "object", Required: required, AdditionalProperties: false, Props: propSlice}
	jRAvro, _ := json.Marshal(rAvro)
	translated = fmt.Sprintf("%s", string(jRAvro))
	return fmt.Sprintf(convertedMessage, refKey, translated, refKey), nil
}
