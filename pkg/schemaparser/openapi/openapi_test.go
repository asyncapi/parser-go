package openapi

import (
	"encoding/json"
	"testing"

	"github.com/asyncapi/parser/pkg/models"
	is "gotest.tools/assert/cmp"

	"gotest.tools/assert"
)

func TestParse(t *testing.T) {
	schema := json.RawMessage(`{
		"payload": {
			"type": "object",
			"nullable": true,
			"example": {
				"test": "example1"
			},
			"examples": [{
				"test": "example0"
			}],
			"properties": {
				"test": {
					"type": "string",
					"nullable": true,
					"example": "test",
					"examples": [
						"fromTestArrayInProperties"
					]
				},
				"test2": {
					"type": ["number"],
					"nullable": true
				},
				"example": {
					"description": "this should not be affected"
				}
			},
			"additionalProperties": {
				"type": "number",
				"nullable": true,
				"example": "test2",
				"examples": [
					"fromTestArrayInAP"
				]
			}
		}
	}`)

	var m models.Message
	err := json.Unmarshal(schema, &m)
	assert.Assert(t, is.Nil(err))

	sp := OpenAPI{}
	err = sp.Parse(&m)
	assert.Assert(t, is.Nil(err))
	json.Marshal(m)
	assert.Equal(t, string(m.Payload), `{"additionalProperties":{"type":["number","null"],"examples":["fromTestArrayInAP","test2"]},"type":["object","null"],"properties":{"example":{"description":"this should not be affected"},"test":{"type":["string","null"],"examples":["fromTestArrayInProperties","test"]},"test2":{"type":["number","null"]}},"examples":[{"test":"example0"},{"test":"example1"}]}`)
}

func TestParseWithNullableAndWithoutType(t *testing.T) {
	schema := json.RawMessage(`{
		"payload": {
			"nullable": true,
			"example": {
				"test": "example1"
			},
			"examples": [{
				"test": "example0"
			}],
			"properties": {
				"test": {
					"nullable": true
				}
			},
			"additionalProperties": {
				"nullable": true
			}
		}
	}`)

	var m models.Message
	err := json.Unmarshal(schema, &m)
	assert.Assert(t, is.Nil(err))

	sp := OpenAPI{}
	err = sp.Parse(&m)
	assert.Assert(t, is.Nil(err))
	json.Marshal(m)
	assert.Equal(t, string(m.Payload), `{"additionalProperties":{"type":["null"]},"type":["null"],"properties":{"test":{"type":["null"]}},"examples":[{"test":"example0"},{"test":"example1"}]}`)
}

func TestParseWithoutNullableAndWithoutType(t *testing.T) {
	schema := json.RawMessage(`{
		"payload": {
			"example": {
				"test": "example1"
			},
			"examples": [{
				"test": "example0"
			}],
			"properties": {
				"test": {
					"description": "hello"
				}
			},
			"additionalProperties": {
				"description": "test"
			}
		}
	}`)

	var m models.Message
	err := json.Unmarshal(schema, &m)
	assert.Assert(t, is.Nil(err))

	sp := OpenAPI{}
	err = sp.Parse(&m)
	assert.Assert(t, is.Nil(err))
	json.Marshal(m)
	assert.Equal(t, string(m.Payload), `{"additionalProperties":{"description":"test"},"properties":{"test":{"description":"hello"}},"examples":[{"test":"example0"},{"test":"example1"}]}`)
}

func TestParseWithoutExamples(t *testing.T) {
	schema := json.RawMessage(`{
		"payload": {
			"example": {
				"test": "example1"
			},
			"properties": {
				"test": {
					"description": "hello",
					"example": "hello"
				}
			},
			"additionalProperties": {
				"description": "test",
				"example": "test"
			}
		}
	}`)

	var m models.Message
	err := json.Unmarshal(schema, &m)
	assert.Assert(t, is.Nil(err))

	sp := OpenAPI{}
	err = sp.Parse(&m)
	assert.Assert(t, is.Nil(err))
	json.Marshal(m)
	assert.Equal(t, string(m.Payload), `{"additionalProperties":{"description":"test","examples":["test"]},"properties":{"test":{"description":"hello","examples":["hello"]}},"examples":[{"test":"example1"}]}`)
}
