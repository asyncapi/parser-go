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

	// JSON object parses into a map with string keys
	itemsMap, ok := f.(map[string]interface{})
	if ok {
		convertedMessage, perr := translate(itemsMap)
		if perr != nil {
			return errs.New(perr.Error())
		}
		*message = []byte(convertedMessage)
	} else {
		log.Print("Union type")
		objectArray := f.([]interface{})
		var uAvro UnionAvro
		for _, o := range objectArray {
			var itemMap interface{}
			switch o.(type) {
			// Complex objects
			case map[string]interface{}:
				log.Printf("Map")
				convertedMessage, perr := translate(o.(map[string]interface{}))
				if perr != nil {
					return errs.New(perr.Error())
				}
				// remove definitions amb move them to the root object
				moveDefintionsToRoot(&uAvro, o.(map[string]interface{})["type"].(string), &convertedMessage)
				uAvro.OneOf = append(uAvro.OneOf, []byte(convertedMessage))
			// Simple objects
			case string:
				log.Printf("String")
				itemMap = o.(string)
				if itemMap == "null" {
					itemsMap = make(map[string]interface{})
					itemsMap["type"] = "null"
					convertedMessage, perr := translate(itemsMap)
					if perr != nil {
						return errs.New(perr.Error())
					}
					uAvro.OneOf = append(uAvro.OneOf, []byte(convertedMessage))
				}
				log.Printf("String %s", itemMap)
			default:
				log.Printf("I don't know about type %T!\n", o)
			}

		}
		bUAvro, err := json.Marshal(uAvro)
		if err != nil {
			return errs.New(err.Error())
		}
		*message = bUAvro
	}

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

func translate(itemsMap map[string]interface{}) (string, *errs.ParserError) {
	var convertedMessage string
	for k, v := range itemsMap {
		if k == "type" {
			switch v {
			case "enum":
				ra := &EnumAvro{}
				r, err := ra.Convert(itemsMap)
				if err != nil {
					return "", errs.New(err.Error())
				}
				convertedMessage = fmt.Sprintf("%s", r)
			case "record":
				ra := &RecordAvro{}
				r, err := ra.Convert(itemsMap)
				if err != nil {
					return "", errs.New(err.Error())
				}
				convertedMessage = fmt.Sprintf("%s", r)
			case "array":
				ra := &ArrayAvro{}
				r, err := ra.Convert(itemsMap)
				if err != nil {
					return "", errs.New(err.Error())
				}
				convertedMessage = fmt.Sprintf("%s", r)
			case "fixed":
				ra := &FixedAvro{}
				r, err := ra.Convert(itemsMap)
				if err != nil {
					return "", errs.New(err.Error())
				}
				convertedMessage = fmt.Sprintf("%s", r)
			case "null", "int", "string", "long", "boolean", "float", "double", "bytes":
				ra := &SimpleAvro{}
				r, err := ra.Convert(itemsMap)
				if err != nil {
					return "", errs.New(err.Error())
				}
				convertedMessage = fmt.Sprintf("%s", r)
			case "map":
				ra := &MapAvro{}
				r, err := ra.Convert(itemsMap)
				if err != nil {
					return "", errs.New(err.Error())
				}
				convertedMessage = fmt.Sprintf("%s", r)
			default:
				log.Println("Unknown type. Please create a Feature request")
				return "", errs.New("Unknown type. Please create a Feature request")
			}
		}
	}
	return convertedMessage, nil
}

func moveDefintionsToRoot(uAvro *UnionAvro, typeObj string, convertedMessage *string) {
	switch typeObj {
	case "enum":
		var enumAvro EnumAvro
		json.Unmarshal([]byte(*convertedMessage), &enumAvro)
		uAvro.Definitions = enumAvro.Definitions
		enumAvro.Definitions = nil
		noDefinitions, _ := json.Marshal(enumAvro)
		*convertedMessage = string(noDefinitions)
	case "record":
		var rAvro RecordAvro
		json.Unmarshal([]byte(*convertedMessage), &rAvro)
		uAvro.Definitions = rAvro.Definitions
		rAvro.Definitions = nil
		noDefinitions, _ := json.Marshal(rAvro)
		*convertedMessage = string(noDefinitions)
	case "array":
		var aAvro ArrayAvro
		json.Unmarshal([]byte(*convertedMessage), &aAvro)
		uAvro.Definitions = aAvro.Definitions
		aAvro.Definitions = nil
		noDefinitions, _ := json.Marshal(aAvro)
		*convertedMessage = string(noDefinitions)
	case "fixed":
		var fAvro FixedAvro
		json.Unmarshal([]byte(*convertedMessage), &fAvro)
		uAvro.Definitions = fAvro.Definitions
		fAvro.Definitions = nil
		noDefinitions, _ := json.Marshal(fAvro)
		*convertedMessage = string(noDefinitions)
	// case "null", "int", "string", "long", "boolean", "float", "double", "bytes":
	case "map":
		var mapAvro ComposeMapAvro
		json.Unmarshal([]byte(*convertedMessage), &mapAvro)
		uAvro.Definitions = mapAvro.Definitions
		mapAvro.Definitions = nil
		noDefinitions, _ := json.Marshal(mapAvro)
		*convertedMessage = string(noDefinitions)
	default:
		log.Println("Unknown type. Please create a Feature request")
	}
}
