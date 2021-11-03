package parser

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/asyncapi/parser-go/pkg/decode"
	"github.com/asyncapi/parser-go/pkg/jsonpath"
	hlsp "github.com/asyncapi/parser-go/pkg/parser/v2"
	"github.com/asyncapi/parser-go/pkg/schema"
	asyncapi "github.com/asyncapi/parser-go/pkg/schema/asyncapi/v2"
	openapi "github.com/asyncapi/parser-go/pkg/schema/openapi/v2"
)

// Parser parses an AsyncAPI document.
type Parser = func(io.Reader, io.Writer) error

type MessageProcessor func(map[string]interface{}) error

type EncoderOpts func(*json.Encoder) error

func (mp MessageProcessor) BuildParser(encoderOpts ...EncoderOpts) Parser {
	return func(reader io.Reader, writer io.Writer) error {
		// fetch document from reader
		var err error
		var jsonData map[string]interface{}
		jsonData, err = decode.ToMap(reader)
		if err != nil {
			return err
		}
		// parse asyncapi schema
		refLoader := jsonpath.NewRefLoader(http.DefaultClient)
		hlsParser := hlsp.NewParser(refLoader, "#/components/schemas")
		err = hlsParser.Parse(jsonData)
		if err != nil {
			return err
		}
		// parse supported schemas
		err = mp(jsonData)
		if err != nil {
			return err
		}
		if writer == nil {
			return nil
		}
		encoder := json.NewEncoder(writer)
		for _, opt := range encoderOpts {
			if err := opt(encoder); err != nil {
				return err
			}
		}
		return encoder.Encode(jsonData)
	}
}

func isURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func isLocalFile(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}

func NewReader(doc string) (io.Reader, error) {
	r := bytes.NewBuffer(nil)
	if isURL(doc) {
		// TODO create a HTTP client with a timeout
		resp, err := http.Get(doc)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		_, err = io.Copy(r, resp.Body)
		return r, err
	}

	if isLocalFile(doc) {
		f, err := os.Open(doc)
		if err != nil {
			return nil, err
		}

		defer f.Close()

		_, err = io.Copy(r, f)
		return r, err
	}

	// Otherwise, doc is considered as the file content in plain text.
	r.WriteString(doc)

	return r, nil
}

func New() (Parser, error) {
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

	d := asyncapi.Dispatcher{}
	for _, schemaParser := range schemaParsers {
		if err := d.Add(schemaParser.parse, schemaParser.labels...); err != nil {
			return nil, err
		}
	}

	messageProcessor := MessageProcessor(asyncapi.BuildMessageProcessor(d))
	parser := messageProcessor.BuildParser(func(encoder *json.Encoder) error {
		encoder.SetIndent("", "  ")
		return nil
	})

	return parser, nil
}
