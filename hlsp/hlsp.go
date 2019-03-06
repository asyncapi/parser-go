package hlsp

import (
	"encoding/json"
	"fmt"

	"github.com/ghodss/yaml"
	"github.com/xeipuuv/gojsonschema"
)

// ParserError is custom struct to hold different error types of the parser
type ParserError struct {
	errorMessage  string
	parsingErrors []gojsonschema.ResultError
}

// Error returns the error message.
func (v *ParserError) Error() string {
	return v.errorMessage
}

// ParsingErrors returns the errors that occurred while parsing the AsyncAPI document.
func (v *ParserError) ParsingErrors() []gojsonschema.ResultError {
	return v.parsingErrors
}

// Parse receives either a YAML or JSON AsyncAPI document.
// It parses the document and checks if it's valid AsyncAPI.
// Skips specification extensions and schemas validation.
// If validation fails, the Parser/Validator should trigger an error.
// Produces a beautified version of the document in JSON Schema Draft 07.
func Parse(AsyncAPI []byte) (json.RawMessage, *ParserError) {
	schemaLoader := gojsonschema.NewReferenceLoader("file://../asyncapi/2.0.0/schema.json")
	AsyncAPI, err := convertFromYAMLtoJSON(AsyncAPI)
	if err != nil {
		return nil, &ParserError{
			errorMessage: err.Error(),
		}
	}

	documentLoader := gojsonschema.NewBytesLoader(AsyncAPI)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return nil, &ParserError{
			errorMessage: err.Error(),
		}
	}

	if result.Valid() {
		return AsyncAPI, nil
	}

	return AsyncAPI, &ParserError{
		errorMessage:  "[Invalid AsyncAPI document] Check out .ParsingErrors() for more information.",
		parsingErrors: result.Errors(),
	}
}

func convertToJSON(doc []byte) ([]byte, error) {
	convertedDoc, err := convertFromYAMLtoJSON(doc)
	if err != nil {
		return convertedDoc, nil
	}

	return nil, fmt.Errorf("[Unsupported document format] Supported formats are: JSON or YAML")
}

func isJSON(doc string) bool {
	return false
}

func convertFromYAMLtoJSON(doc []byte) ([]byte, error) {
	return yaml.YAMLToJSON(doc)
}
