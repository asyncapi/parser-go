package parser

import (
	"encoding/json"

	"github.com/asyncapi/parser/pkg/errs"
	"github.com/asyncapi/parser/pkg/hlsp"
)

// Parse receives either a YAML or JSON AsyncAPI document, and tries to parse it.
func Parse(yamlOrJSONDocument []byte) (json.RawMessage, *errs.ParserError) {
	doc, err := hlsp.Parse(yamlOrJSONDocument)

	if err != nil {
		return nil, err
	}
	return doc, nil
}
