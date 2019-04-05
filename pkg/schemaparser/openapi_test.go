package schemaparser

import (
	"encoding/json"
	"testing"

	is "gotest.tools/assert/cmp"

	"gotest.tools/assert"
)

func TestOpenAPIParse(t *testing.T) {
	schema := json.RawMessage(`{
		"type": "object",
		"properties": {
			"test": {
				"type": "string",
				"nullable": true
			}
		}
	}`)
	sp := OpenAPI{}

	err := sp.Parse(schema)
	assert.Assert(t, is.Nil(err))
}
