package errs

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

// New creates a new ParserError with the given message.
func New(msg string) *ParserError {
	return &ParserError{
		ErrorMessage: msg,
	}
}

// NewWithParsingErrors creates a new ParserError with the given message and parsing errors.
func NewWithParsingErrors(msg string, parsingErrors []gojsonschema.ResultError) *ParserError {
	return &ParserError{
		ErrorMessage:  msg,
		ParsingErrors: parsingErrors,
	}
}
