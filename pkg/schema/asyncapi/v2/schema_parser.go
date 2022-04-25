package v2

import (
	"fmt"

	"github.com/pkg/errors"

	parseSchema "github.com/asyncapi/parser-go/pkg/schema"
	schemas "github.com/asyncapi/spec-json-schemas/v2"
)

var (
	Labels  = []string{"asyncapi"}
	parsers = make(map[string]*parseSchema.Parser)
)

// Parse parsers a document.
func Parse(v interface{}) error {
	version, err := extractVersion(v)
	if err != nil {
		return errors.Wrap(err, "error extracting AsyncAPI Spec version from provided document")
	}

	if parsers[version] == nil {
		s, err := schemas.Get(version)
		if err != nil {
			return err
		}

		if s == nil {
			return fmt.Errorf("version %q is not supported", version)
		}

		parsers[version] = parseSchema.NewParser(s)
	}

	return parsers[version].Parse(v)
}

func extractVersion(v interface{}) (string, error) {
	switch doc := v.(type) {
	case map[string]interface{}:
		if doc["asyncapi"] == nil {
			return "", errors.New("the `asyncapi` field is missing")
		}

		return doc["asyncapi"].(string), nil
	default:
		return "", errors.New("only map[string]interface{} type is supported")
	}
}
