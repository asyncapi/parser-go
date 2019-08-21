package decode

import (
	"github.com/asyncapi/converter-go/pkg/decode"
	"github.com/pkg/errors"

	"io"
)

var ErrUnableToDecodeDocument = errors.New("unable to decode document")

func ToMap(reader io.Reader) (map[string]interface{}, error) {
	var jsonData interface{}
	if err := decode.FromJSONWithYamlFallback(&jsonData, reader); err != nil {
		return nil, err
	}
	var (
		result map[string]interface{}
		ok     bool
	)
	if result, ok = jsonData.(map[string]interface{}); !ok {
		return nil, ErrUnableToDecodeDocument
	}
	return result, nil
}
