package openapi

import (
	"encoding/json"

	"github.com/asyncapi/parser/pkg/errs"
	"github.com/asyncapi/parser/pkg/models"
	"github.com/xeipuuv/gojsonschema"
)

type OpenAPI struct {
	schema *gojsonschema.Schema
}

// Parse parses and validates an OpenAPI/AsyncAPI 1.x schema.
func (v OpenAPI) Parse(message *models.Message) *errs.ParserError {
	if v.schema == nil {
		schemaLoader := gojsonschema.NewBytesLoader(getSchema())
		sch, err := gojsonschema.NewSchema(schemaLoader)
		if err != nil {
			return errs.New(err.Error())
		}
		v.schema = sch
	}

	payloadBytes, err := message.Payload.MarshalJSON()
	if err != nil {
		return errs.New(err.Error())
	}

	documentLoader := gojsonschema.NewBytesLoader(payloadBytes)
	result, err := v.schema.Validate(documentLoader)
	if err != nil {
		return errs.New(err.Error())
	}

	if result.Valid() {
		if err := convertSchema(message); err != nil {
			return err
		}

		return nil
	}

	return errs.NewWithParsingErrors(
		"[Invalid OpenAPI/AsyncAPI schema] Check out err.ParsingErrors() for more information.",
		result.Errors(),
	)
}

func convertSchema(message *models.Message) *errs.ParserError {
	var schema Schema
	err := json.Unmarshal(message.Payload, &schema)
	if err != nil {
		return errs.New(err.Error())
	}

	if err := walkSchema(&schema); err != nil {
		return err
	}

	j, err := json.Marshal(schema)
	if err != nil {
		return errs.New(err.Error())
	}

	message.Payload = j

	return nil
}

func walkSchema(schema *Schema) *errs.ParserError {
	if schema.Nullable != nil && *schema.Nullable == true {
		var typeArray []string
		typeArray, ok := schema.Type.([]string)
		if !ok {
			typeString, ok := schema.Type.(string)
			if typeString != "null" {
				if ok {
					typeArray = make([]string, 1)
					typeArray[0] = typeString
				} else {
					typeArray = make([]string, 0)
				}
				schema.Type = append(typeArray, "null")
			}
		} else {
			if !contains(typeArray, "null") {
				schema.Type = append(typeArray, "null")
			}
		}
		schema.Nullable = nil
	}

	if schema.Example != nil {
		if schema.Examples == nil {
			schema.Examples = make([]json.RawMessage, 0)
		}

		schema.Examples = append(schema.Examples, schema.Example)
		schema.Example = nil
	}

	if schema.Properties != nil {
		for _, prop := range schema.Properties {
			walkSchema(prop)
		}
	}

	if schema.AdditionalProperties != nil {
		_, isBool := schema.AdditionalProperties.(bool)
		if !isBool {
			apSchema, isSchema := schema.AdditionalProperties.(Schema)
			if isSchema {
				walkSchema(&apSchema)
				schema.AdditionalProperties = apSchema
			}
		}
	}

	return nil
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
