package avro

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/asyncapi/parser/pkg/errs"
	"github.com/linkedin/goavro"
)

// Reference maps a ref object with its definitions
type Reference struct {
	Ref         json.RawMessage `json:"$ref"`
	Definitions json.RawMessage `json:"definitions,omitempty"`
}

//SchemaConverter is the basic interface that schema converters needs to implement
type SchemaConverter interface {
	Convert(map[string]interface{}) (string, *errs.ParserError)
}

// Parse parses and validates an Avro schema.
func Parse(message *json.RawMessage) *errs.ParserError {
	codec, err := goavro.NewCodec(string(*message))

	codec.Schema()
	if err != nil {
		return errs.New(err.Error())
	}

	var f interface{}
	err = json.Unmarshal([]byte(codec.Schema()), &f)
	if err != nil {
		log.Println("Error parsing JSON: ", err)
	}

	var convertedMessage string
	// JSON object parses into a map with string keys
	itemsMap, ok := f.(map[string]interface{})
	if ok {
		for k, v := range itemsMap {
			if k == "type" {
				switch v {
				case "enum":
					ra := &EnumAvro{}
					r, err := ra.Convert(itemsMap)
					if err != nil {
						return errs.New(err.Error())
					}
					convertedMessage = fmt.Sprintf("%s", r)
				case "record":
					ra := &RecordAvro{}
					r, err := ra.Convert(itemsMap)
					if err != nil {
						return errs.New(err.Error())
					}
					convertedMessage = fmt.Sprintf("%s", r)
				case "array":
					ra := &ArrayAvro{}
					r, err := ra.Convert(itemsMap)
					if err != nil {
						return errs.New(err.Error())
					}
					convertedMessage = fmt.Sprintf("%s", r)
				case "fixed":
					ra := &FixedAvro{}
					r, err := ra.Convert(itemsMap)
					if err != nil {
						return errs.New(err.Error())
					}
					convertedMessage = fmt.Sprintf("%s", r)
				case "null", "int", "string", "long", "boolean", "float", "double", "bytes":
					ra := &SimpleAvro{}
					r, err := ra.Convert(itemsMap)
					if err != nil {
						return errs.New(err.Error())
					}
					convertedMessage = fmt.Sprintf("%s", r)
				case "map":
					ra := &MapAvro{}
					r, err := ra.Convert(itemsMap)
					if err != nil {
						return errs.New(err.Error())
					}
					convertedMessage = fmt.Sprintf("%s", r)
				default:
					log.Println("Unknown type. Please create a Feature request")
					return errs.New("Unknown type. Please create a Feature request")
				}
			}
		}
	} else {
		log.Print("Union type")
		objectArray := f.([]interface{})
		for _, o := range objectArray {
			var itemMap interface{}
			switch o.(type) {
			// Complex objects
			case map[string]interface{}:
				log.Printf("Map")
				itemMap = o.(map[string]interface{})
				for k, v := range itemMap.(map[string]interface{}) {
					log.Printf("key %s,value %s \n", k, v)
					if k == "type" {
						switch v {
						case "record":
							log.Printf("Record type %s \n", v.(string))
							// ra := &RecordAvro{}
							// ra.Convert(v.([]byte))
						default:
							log.Println("Unknown type. Please create a Feature request")
						}
					}
				}
			// Simple objects
			case string:
				log.Printf("String")
				itemMap = o.(string)
				log.Printf("String %s", itemMap)
			default:
				log.Printf("I don't know about type %T!\n", o)
			}

		}
	}

	*message = []byte(convertedMessage)

	return nil
}

// TODO: int, long, float
func convertValues(attrValues string) AdditionalPropertiesItem {
	log.Printf("Values %s", attrValues)
	switch attrValues {
	case "bytes":
		return AdditionalPropertiesItem{Type: convertType(attrValues), Pattern: "^[\u0000-\u00ff]*$"}
	case "int":
		return AdditionalPropertiesItem{Type: convertType(attrValues), Min: minInt, Max: maxInt}
	case "long":
		return AdditionalPropertiesItem{Type: convertType(attrValues), Min: minLong, Max: maxLong}
	case "float", "double":
		return AdditionalPropertiesItem{Type: convertType(attrValues)}
	case "map":
		return AdditionalPropertiesItem{Type: convertType(attrValues)}
	default:
		return AdditionalPropertiesItem{Type: attrValues}
	}
}

func convertType(attrType string) string {
	switch attrType {
	case "long", "int":
		return "integer"
	case "float", "double":
		return "number"
	case "bytes":
		return "string"
	case "map":
		return "object"
	default:
		return attrType
	}
}
