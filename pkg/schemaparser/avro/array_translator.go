package avro

import (
	"encoding/json"
	"log"

	"github.com/asyncapi/parser/pkg/errs"
)

// ArrayItems Items of the array
type ArrayItems struct {
	Type                 string                   `json:"type"`
	AdditionalProperties AdditionalPropertiesItem `json:"additionalProperties,omitempty"`
}

// AdditionalPropertiesItem maps additional properties object
type AdditionalPropertiesItem struct {
	Type    string `json:"type"`
	Pattern string `json:"pattern,omitempty"`
}

// ArrayAvro maps avro array scheme
type ArrayAvro struct {
	Type  string     `json:"type,omitempty"`
	Items ArrayItems `json:"items,omitempty"`
}

// SimpleArrayAvro maps simple array scheme
type SimpleArrayAvro struct {
	Type  string `json:"type,omitempty"`
	Items string `json:"items,omitempty"`
}

// Convert transforms avro formatted message to JSONSchema
func (ra *ArrayAvro) Convert(message map[string]interface{}) (string, *errs.ParserError) {
	var aI ArrayItems
	var aA ArrayAvro
	var sAa SimpleArrayAvro
	var aAbytes []byte
	switch message["items"].(type) {
	// Simple objects
	case string:
		log.Printf("String")
		aI = ArrayItems{Type: message["items"].(string)}
		sAa = SimpleArrayAvro{Type: "array", Items: message["items"].(string)}
		aAbytes, err := json.Marshal(sAa)
		if err != nil {
			return "", errs.New(err.Error())
		}
		return string(aAbytes), nil
		// Complex objects
	case map[string]interface{}:
		log.Printf("Map")
		itemMap := message["items"].(map[string]interface{})
		aI = ArrayItems{Type: itemMap["type"].(string), AdditionalProperties: convertValues(itemMap["values"].(string))}
		aA = ArrayAvro{Type: "array", Items: aI}
		aAbytes, err := json.Marshal(aA)
		if err != nil {
			return "", errs.New(err.Error())
		}
		return string(aAbytes), nil
	default:
		log.Printf("I don't know about type %T!\n", message)
	}

	return string(aAbytes), nil
}

func convertValues(attrValues string) AdditionalPropertiesItem {
	log.Printf("Values %s", attrValues)
	switch attrValues {
	case "bytes":
		return AdditionalPropertiesItem{Type: "string", Pattern: "^[\u0000-\u00ff]*$"}
	default:
		return AdditionalPropertiesItem{}
	}
}
