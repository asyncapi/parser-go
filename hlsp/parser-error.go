package hlsp

import (
	"github.com/xeipuuv/gojsonschema"
)

// ParserError is custom struct to hold different error types of the parser
type ParserError struct {
	ErrorMessage  string
	ParsingErrors []gojsonschema.ResultError
}

// Error returns the error message.
func (v *ParserError) Error() string {
	return v.ErrorMessage
}
