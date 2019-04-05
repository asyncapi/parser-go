package schemaparser

import (
	"encoding/json"

	"github.com/asyncapi/parser/pkg/errs"
	"github.com/xeipuuv/gojsonschema"
)

type OpenAPI struct {
	schema *gojsonschema.Schema
}

// Parse parses and validates an OpenAPI/AsyncAPI 1.x schema.
func (v OpenAPI) Parse(schema json.RawMessage) *errs.ParserError {
	if v.schema == nil {
		schemaLoader := gojsonschema.NewBytesLoader(v.getSchema())
		sch, err := gojsonschema.NewSchema(schemaLoader)
		if err != nil {
			return &errs.ParserError{
				ErrorMessage: err.Error(),
			}
		}
		v.schema = sch
	}

	documentLoader := gojsonschema.NewBytesLoader(schema)

	result, err := v.schema.Validate(documentLoader)
	if err != nil {
		return &errs.ParserError{
			ErrorMessage: err.Error(),
		}
	}

	if result.Valid() {
		return nil
	}

	return &errs.ParserError{
		ErrorMessage:  "[Invalid OpenAPI/AsyncAPI schema] Check out err.ParsingErrors() for more information.",
		ParsingErrors: result.Errors(),
	}
}
