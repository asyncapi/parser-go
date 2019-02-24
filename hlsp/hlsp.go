package hlsp

import (
    "fmt"
)

// Parse receives either a YAML or JSON AsyncAPI document.
// It parses the document and checks if it's valid AsyncAPI.
// Skips specification extensions and schemas validation.
// If validation fails, the Parser/Validator should trigger an error.
// Produces a beautified version of the document in JSON Schema Draft 07.
func Parse(document string) (string, error) {

	return "", nil
}

func convertToJSON(doc string) (string, error) {
	if isJSON(doc) {
		return doc, nil
	}
	
	convertedDoc, err := convertFromYAMLtoJSON(doc)
	if err != nil {
		return convertedDoc, nil
	}

	return "", fmt.Errorf("Unsupported document format. Supported formats are: JSON or YAML.")
}

func isJSON(doc string) bool {
	return false
}

func convertFromYAMLtoJSON(doc string) (string, error) {
	return "", nil
}
