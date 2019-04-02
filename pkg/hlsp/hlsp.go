package hlsp

import (
	"encoding/json"

	"github.com/ghodss/yaml"
	"github.com/xeipuuv/gojsonschema"
)

// Parse receives either a YAML or JSON AsyncAPI document, and tries to parse it.
func Parse(yamlOrJSONDocument []byte) (json.RawMessage, *ParserError) {
	jsonDocument, err := yaml.YAMLToJSON(yamlOrJSONDocument)
	if err != nil {
		return nil, &ParserError{
			ErrorMessage: err.Error(),
		}
	}
	if string(jsonDocument) == "null" {
		return nil, &ParserError{
			ErrorMessage: "[Invalid AsyncAPI document] Document is empty or null.",
		}
	}

	beautifiedDoc, e := ParseJSON(jsonDocument)
	if e != nil {
		return nil, e
	}
	return beautifiedDoc, nil
}

// ParseJSON receives a JSON AsyncAPI document.
// It parses the document and checks if it's valid AsyncAPI.
// Skips specification extensions and schemas validation.
// If validation fails, the Parser/Validator should trigger an error.
// Produces a beautified version of the document in JSON Schema Draft 07.
func ParseJSON(jsonDocument []byte) (json.RawMessage, *ParserError) {
	schemaLoader := gojsonschema.NewBytesLoader(getSchema())
	documentLoader := gojsonschema.NewBytesLoader(jsonDocument)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return nil, &ParserError{
			ErrorMessage: err.Error(),
		}
	}

	if result.Valid() {
		beautifiedDoc, err := Beautify(jsonDocument)
		if err != nil {
			return nil, &ParserError{
				ErrorMessage: err.Error(),
			}
		}
		return beautifiedDoc, nil
	}

	return jsonDocument, &ParserError{
		ErrorMessage:  "[Invalid AsyncAPI document] Check out err.ParsingErrors() for more information.",
		ParsingErrors: result.Errors(),
	}
}
