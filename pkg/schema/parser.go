package schema

import (
	parserErrors "github.com/asyncapi/parser/pkg/error"

	"github.com/pkg/errors"
	"github.com/xeipuuv/gojsonschema"
)

type ParseMessage func(interface{}) error

var _ ParseMessage = (&Parser{}).Parse

type Parser struct {
	schemaLoader gojsonschema.JSONLoader
}

func NewParser(schema []byte) Parser {
	return Parser{
		schemaLoader: gojsonschema.NewBytesLoader(schema),
	}
}

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
