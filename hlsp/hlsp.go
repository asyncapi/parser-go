package hlsp

import (
	"encoding/json"
	"path"
	"runtime"

	"github.com/asyncapi/parser/models"
	"github.com/ghodss/yaml"
	"github.com/xeipuuv/gojsonschema"
)

// Parse receives either a YAML or JSON AsyncAPI document, and tries to parse it.
func Parse(yamlOrJSONDocument []byte) (*models.ParsedAsyncAPI, *ParserError) {
	jsonDocument, err := yaml.YAMLToJSON(yamlOrJSONDocument)
	if err != nil {
		return nil, &ParserError{
			errorMessage: err.Error(),
		}
	}
	if string(jsonDocument) == "null" {
		return nil, &ParserError{
			errorMessage: "[Invalid AsyncAPI document] Document is empty or null.",
		}
	}

	beautifiedDoc, e := ParseJSON(jsonDocument)
	if e != nil {
		return nil, e
	}

	var parsedAsyncAPI models.ParsedAsyncAPI
	err = json.Unmarshal(beautifiedDoc, &parsedAsyncAPI)
	if err != nil {
		return nil, &ParserError{
			errorMessage: err.Error(),
		}
	}

	return &parsedAsyncAPI, nil
}

// ParseJSON receives a JSON AsyncAPI document.
// It parses the document and checks if it's valid AsyncAPI.
// Skips specification extensions and schemas validation.
// If validation fails, the Parser/Validator should trigger an error.
// Produces a beautified version of the document in JSON Schema Draft 07.
func ParseJSON(jsonDocument []byte) (json.RawMessage, *ParserError) {
	_, filename, _, ok := runtime.Caller(1)
	if ok == false {
		return nil, &ParserError{
			errorMessage: "[Unexpected error] Could not resolve relative file path to AsyncAPI schema.",
		}
	}
	filepath := path.Join(path.Dir(filename), "../asyncapi/2.0.0/schema.json")
	schemaLoader := gojsonschema.NewReferenceLoader("file://" + filepath)
	documentLoader := gojsonschema.NewBytesLoader(jsonDocument)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return nil, &ParserError{
			errorMessage: err.Error(),
		}
	}

	if result.Valid() {
		beautifiedDoc, err := Beautify(jsonDocument)
		if err != nil {
			return nil, &ParserError{
				errorMessage: err.Error(),
			}
		}
		return beautifiedDoc, nil
	}

	return jsonDocument, &ParserError{
		errorMessage:  "[Invalid AsyncAPI document] Check out err.ParsingErrors() for more information.",
		parsingErrors: result.Errors(),
	}
}
