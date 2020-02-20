package cli

import (
	"github.com/asyncapi/parser-go/pkg/parser"
	"github.com/asyncapi/parser-go/pkg/schema"
	asyncapi "github.com/asyncapi/parser-go/pkg/schema/asyncapi/v2"
	openapi "github.com/asyncapi/parser-go/pkg/schema/openapi/v2"

	"github.com/docopt/docopt-go"
	"github.com/pkg/errors"

	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

var (
	errFileDoesNotExist = errors.New("file does not exist")
)

// Parser parses an AsyncAPI document.
type Parser = func(io.Reader, io.Writer) error

var _ parser.Parse = Parser(nil)

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

func isURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func (h Cli) reader() (io.Reader, error) {
	fileOption := h.Opts["<PATH>"]
	path := fmt.Sprintf("%v", fileOption)
	if isURL(path) {
		resp, err := http.Get(path)
		if err != nil {
			return nil, err
		}
		return resp.Body, nil
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.Wrap(errFileDoesNotExist, path)
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// NewParserAndReader creates both a parser and a reader of the AsyncAPI document.
func (h Cli) NewParserAndReader() (Parser, io.Reader, error) {
	var messageProcessor parser.MessageProcessor
	reader, err := h.reader()
	if err != nil {
		return nil, nil, err
	}
	d := asyncapi.Dispatcher{}
	schemaParsers := []struct {
		parse  schema.ParseMessage
		labels []string
	}{
		{
			openapi.Parse,
			openapi.Labels,
		},
		{
			asyncapi.Parse,
			asyncapi.Labels,
		},
	}
	for _, schemaParser := range schemaParsers {
		err = d.Add(schemaParser.parse, schemaParser.labels...)
		if err != nil {
			return nil, nil, err
		}
	}
	messageProcessor = asyncapi.BuildMessageProcessor(d)
	parse := messageProcessor.BuildParse(func(encoder *json.Encoder) error {
		encoder.SetIndent("", "  ")
		return nil
	})
	return parse, reader, err
}
