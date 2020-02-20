package v2

import (
	parserErrors "github.com/asyncapi/parser-go/pkg/error"
	parserSchema "github.com/asyncapi/parser-go/pkg/schema"
	"github.com/asyncapi/parser-go/pkg/schema/jsonschema/draft07"

	"github.com/pkg/errors"

	"fmt"
)

var (
	ErrInvalidSchema                              = errors.New("invalid OpenAPI/AsyncAPI schema")
	_                   parserSchema.ParseMessage = Parse
	openapiSchemaParser                           = parserSchema.NewParser(schema)
	null                                          = "null"
	Labels                                        = []string{
		"",
		"openapi",
		"application/vnd.oai.openapi",
		"application/vnd.asyncapi",
	}
)

func Parse(data interface{}) error {
	message, ok := data.(*map[string]interface{})
	if !ok {
		return nil
	}
	payload, found := (*message)["payload"]
	if !found {
		return nil
	}
	if err := openapiSchemaParser.Parse(payload); err != nil {
		// not a boolean value or object
		return err
	}
	schema, ok := payload.(map[string]interface{})
	if !ok {
		// a boolean value
		return nil
	}
	return reduceAndCorrectSchema(&schema)
}

func reduceAndCorrectSchema(schema *map[string]interface{}) error {
	var errs []error
	for _, fn := range []func(*map[string]interface{}) error{
		reduceExample,
		reduceAndCorrectProperties,
		reduceAndCorrectAdditionalProperties,
		correctType,
	} {
		err := fn(schema)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return parserErrors.New(errs...)
}

func reduceExample(schema *map[string]interface{}) error {
	example := mapValue("example", schema)
	if example == nil {
		return nil
	}
	examples := mapSliceValue("examples", schema)
	examples = append(examples, example)
	(*schema)["examples"] = examples
	delete(*schema, "example")
	return nil
}

func reduceAndCorrectProperties(schema *map[string]interface{}) error {
	properties := mapObjectValue("properties", schema)
	if properties == nil {
		return nil
	}
	var errs []error
	for k, v := range *properties {
		property, ok := v.(map[string]interface{})
		if !ok {
			errs = append(errs, errors.Wrap(ErrInvalidSchema, k))
			continue
		}
		err := draft07.Parse(property)
		if err != nil {
			// property is not a schema
			continue
		}
		err = reduceAndCorrectSchema(&property)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 1 {
		return parserErrors.New(errs...)
	}
	(*schema)["properties"] = properties
	return nil
}

func reduceAndCorrectAdditionalProperties(schema *map[string]interface{}) error {
	additionalProperties := mapObjectValue("additionalProperties", schema)
	if additionalProperties == nil {
		return nil
	}
	err := draft07.Parse(additionalProperties)
	if err != nil {
		// additional property is not a schema
		return nil
	}
	err = reduceAndCorrectSchema(additionalProperties)
	if err != nil {
		return err
	}
	(*schema)["additionalProperties"] = additionalProperties
	return nil
}

func mapValue(key string, v interface{}) interface{} {
	value, ok := v.(*map[string]interface{})
	if !ok {
		return nil
	}
	return (*value)[key]
}

func mapSliceValue(key string, v interface{}) []interface{} {
	value := mapValue(key, v)
	if value == nil {
		return nil
	}
	sliceValue, ok := value.([]interface{})
	if !ok {
		return nil
	}
	return sliceValue
}

func mapObjectValue(key string, v interface{}) *map[string]interface{} {
	value := mapValue(key, v)
	if value == nil {
		return nil
	}
	objectValue, ok := value.(map[string]interface{})
	if !ok {
		return nil
	}
	return &objectValue
}

func correctType(schema *map[string]interface{}) error {
	nullable, ok := (*schema)["nullable"].(bool)
	if !ok || !nullable {
		return nil
	}
	schemaType := mapValue("type", schema)
	schemaTypeSlice, ok := schemaType.([]interface{})
	if !ok {
		(*schema)["type"] = toStringSlice(schemaType)
		delete(*schema, "nullable")
		return nil
	}
	if containsNull(schemaTypeSlice) {
		return nil
	}
	(*schema)["type"] = append(schemaTypeSlice, null)
	delete(*schema, "nullable")
	return nil
}

func toStringSlice(v interface{}) []interface{} {
	if v == nil || fmt.Sprintf("%v", v) == null {
		return []interface{}{null}
	}
	return []interface{}{v, null}
}

func containsNull(typeSlice []interface{}) bool {
	for _, itemType := range typeSlice {
		typeStr := fmt.Sprintf("%v", itemType)
		if typeStr == null {
			return true
		}
	}
	return false
}
