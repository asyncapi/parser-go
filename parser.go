package parser

import (
	"github.com/asyncapi/parser/hlsp"
	"github.com/asyncapi/parser/models"
)

// Parse receives either a YAML or JSON AsyncAPI document, and tries to parse it.
func Parse(yamlOrJSONDocument []byte) (*models.ParsedAsyncAPI, *hlsp.ParserError) {
	doc, err := hlsp.Parse(yamlOrJSONDocument)

	return doc, err
}
