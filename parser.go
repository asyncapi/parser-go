package parser

import (
	"encoding/json"

	"github.com/asyncapi/parser/hlsp"
)

// Parse receives either a YAML or JSON AsyncAPI document, and tries to parse it.
func Parse(yamlOrJSONDocument []byte) (json.RawMessage, *hlsp.ParserError) {
	doc, err := hlsp.Parse(yamlOrJSONDocument)

	if err != nil {
		return nil, err
	}

	return doc, nil
}

// Parse receives either a YAML or JSON AsyncAPI document, and tries to parse it.
// func Parse(yamlOrJSONDocument []byte) (*models.ParsedAsyncAPI, *hlsp.ParserError) {
// 	doc, err := hlsp.Parse(yamlOrJSONDocument)

// 	if err != nil {
// 		return nil, err
// 	}

// 	var parsedAsyncAPI models.ParsedAsyncAPI
// 	e := json.Unmarshal(doc, &parsedAsyncAPI)
// 	if e != nil {
// 		return nil, &hlsp.ParserError{
// 			ErrorMessage: e.Error(),
// 		}
// 	}

// 	return &parsedAsyncAPI, nil
// }
