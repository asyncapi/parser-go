package cli

import (
	"io"

	"github.com/asyncapi/parser-go/pkg/parser"
	"github.com/docopt/docopt-go"
)

// Parser parses an AsyncAPI document.
type Parser = func(io.Reader, io.Writer) error

// Cli is a helper type that allows you to instantiate the AsyncAPI Converter and io.Reader of
// the converted document with arguments passed from the terminal.
type Cli struct {
	docopt.Opts
}

// New returns a new Cli instance.
func New(opts docopt.Opts) Cli {
	return Cli{
		Opts: opts,
	}
}

// NewParserAndReader creates both a parser and a reader of the AsyncAPI document.
func (h Cli) NewParserAndReader() (Parser, io.Reader, error) {
	fileOption := h.Opts["<PATH>"]
	r, err := parser.NewReader(fileOption.(string))
	if err != nil {
		return nil, nil, err
	}

	p, err := parser.New()
	if err != nil {
		return nil, nil, err
	}

	return p, r, nil
}
