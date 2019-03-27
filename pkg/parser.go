package parser

import (
	"encoding/json"

	"github.com/asyncapi/parser/pkg/hlsp"
)

var (
	DefaultDefinitionFile = "asyncapi/2.0.0/example.yaml"
)

// Parse receives either a YAML or JSON AsyncAPI document, and tries to parse it.
func Parse(yamlOrJSONDocument []byte) (json.RawMessage, *hlsp.ParserError) {
	doc, err := hlsp.Parse(yamlOrJSONDocument)

	if err != nil {
		return nil, err
	}

	return doc, nil
}
