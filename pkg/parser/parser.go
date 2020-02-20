package parser

import (
	"github.com/asyncapi/parser-go/pkg/decode"
	"github.com/asyncapi/parser-go/pkg/jsonpath"
	hlsp "github.com/asyncapi/parser-go/pkg/parser/v2"

	"encoding/json"
	"io"
	"net/http"
)

type Parse func(io.Reader, io.Writer) error

type MessageProcessor func(*map[string]interface{}) error

type EncoderOpts func(*json.Encoder) error

func (mp MessageProcessor) BuildParse(encoderOpts ...EncoderOpts) Parse {
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
		err = hlsParser.Parse(&jsonData)
		if err != nil {
			return err
		}
		// parse supported schemas
		err = mp(&jsonData)
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
