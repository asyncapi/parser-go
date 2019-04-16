package hlsp

import (
	"github.com/asyncapi/parser/pkg/dereferencer"
	"encoding/json"

	"github.com/asyncapi/parser/pkg/models"

	"github.com/asyncapi/parser/pkg/errs"
	"github.com/ghodss/yaml"
	"github.com/xeipuuv/gojsonschema"
)

// Parse receives either a YAML or JSON AsyncAPI document, and tries to parse it.
func Parse(yamlOrJSONDocument []byte) (*models.AsyncapiDocument, *errs.ParserError) {
	jsonDocument, err := ParseForC(yamlOrJSONDocument)
	if err != nil {
		return nil, err
	}

	var doc models.AsyncapiDocument
	if err := json.Unmarshal(jsonDocument, &doc); err != nil {
		return nil, &errs.ParserError{
			ErrorMessage: err.Error(),
		}
	}

	return &doc, nil
}

// ParseForC receives either a YAML or JSON AsyncAPI document, and tries to parse it.
// Returns the result as a json.RawMessage.
func ParseForC(yamlOrJSONDocument []byte) (json.RawMessage, *errs.ParserError) {
	jsonDocument, err := yaml.YAMLToJSON(yamlOrJSONDocument)
	if err != nil {
		return nil, &errs.ParserError{
			ErrorMessage: err.Error(),
		}
	}
	if string(jsonDocument) == "null" {
		return nil, &errs.ParserError{
			ErrorMessage: "[Invalid AsyncAPI document] Document is empty or null.",
		}
	}

	jsonBytes, e := ParseJSON(jsonDocument)
	if e != nil {
		return nil, e
	}
	return jsonBytes, nil
}

// ParseJSON receives a JSON AsyncAPI document.
// It parses the document and checks if it's valid AsyncAPI.
// Skips specification extensions and schemas validation.
// If validation fails, the Parser/Validator should trigger an error.
// Produces a beautified version of the document in JSON Schema Draft 07.
func ParseJSON(jsonDocument []byte) (json.RawMessage, *errs.ParserError) {
	schemaLoader := gojsonschema.NewBytesLoader(getSchema())
	documentLoader := gojsonschema.NewBytesLoader(jsonDocument)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return nil, &errs.ParserError{
			ErrorMessage: err.Error(),
		}
	}

	if result.Valid() {
		jsonDocument, err = dereferencer.Dereference(jsonDocument)
		return jsonDocument, nil
	}

	return jsonDocument, &errs.ParserError{
		ErrorMessage:  "[Invalid AsyncAPI document] Check out err.ParsingErrors() for more information.",
		ParsingErrors: result.Errors(),
	}
}
