package schema

import (
	parserErrors "github.com/asyncapi/parser-go/pkg/error"

	"github.com/pkg/errors"
	"github.com/xeipuuv/gojsonschema"
)

type ParseMessage func(interface{}) error

var _ ParseMessage = (&Parser{}).Parse

// Parser parses a given structure and validates it against a JSON Schema.
type Parser struct {
	schemaLoader gojsonschema.JSONLoader
}

// NewParser creates a new Parser.
func NewParser(schema []byte) *Parser {
	return &Parser{
		schemaLoader: gojsonschema.NewBytesLoader(schema),
	}
}

// Parse is the main function to parse a given structure and validate it against a JSON Schema.
// Implements ParseMessage.
func (p *Parser) Parse(data interface{}) error {
	documentLoader := gojsonschema.NewGoLoader(data)
	result, err := gojsonschema.Validate(p.schemaLoader, documentLoader)
	if err != nil {
		return err
	}
	if !result.Valid() {
		var errs []error
		for _, err := range result.Errors() {
			errs = append(errs, errors.New(err.String()))
		}
		return parserErrors.New(errs...)
	}
	return nil
}
