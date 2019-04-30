package parser

import (
	"encoding/json"

	"github.com/asyncapi/parser/pkg/errs"
	"github.com/asyncapi/parser/pkg/hlsp"
	"github.com/asyncapi/parser/pkg/schemaparser"
)

// Parse receives either a YAML or JSON AsyncAPI document, and tries to parse it.
func Parse(yamlOrJSONDocument []byte, circularReferences bool) (json.RawMessage, *errs.ParserError) {
	doc, err := hlsp.Parse(yamlOrJSONDocument, circularReferences)
	if err != nil {
		return nil, err
	}

	err = schemaparser.Parse(doc)
	if err != nil {
		return nil, err
	}

	docBytes, e := json.Marshal(doc)
	if e != nil {
		return nil, &errs.ParserError{
			ErrorMessage: e.Error(),
		}
	}

	return docBytes, nil
}
