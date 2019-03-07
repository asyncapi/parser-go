package hlsp

import (
	"github.com/xeipuuv/gojsonschema"
)

// ParserError is custom struct to hold different error types of the parser
type ParserError struct {
	errorMessage  string
	parsingErrors []gojsonschema.ResultError
}

// Error returns the error message.
func (v *ParserError) Error() string {
	return v.errorMessage
}

// ParsingErrors returns the errors that occurred while parsing the AsyncAPI document.
func (v *ParserError) ParsingErrors() []gojsonschema.ResultError {
	return v.parsingErrors
}
