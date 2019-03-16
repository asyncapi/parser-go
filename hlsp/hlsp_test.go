package hlsp

import (
	"encoding/json"
	"testing"

	"github.com/asyncapi/parser/models"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestParse(t *testing.T) {
	asyncapi := []byte(`
asyncapi: '2.0.0'
id: myapi
info:
  title: My API
  version: '1.0.0'
channels: {}`)

	doc, err := Parse(asyncapi)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, doc.Asyncapi, "2.0.0")
	assert.Equal(t, doc.Id, "myapi")
	assert.Equal(t, doc.Info.Title, "My API")
	assert.Equal(t, doc.Info.Version, "1.0.0")
}
func TestParseWithEmptyYAML(t *testing.T) {
	asyncapi := []byte(``)

	_, err := Parse(asyncapi)
	assert.Equal(t, err.Error(), "[Invalid AsyncAPI document] Document is empty or null.")
	assert.Equal(t, len(err.ParsingErrors()), 0)
}
func TestParseWithInvalidDocument(t *testing.T) {
	asyncapi := []byte(`
asyncapi: '2.0.0'
info:
  title: My API
  version: '1.0.0'
channels: {}`)

	_, err := Parse(asyncapi)
	assert.Equal(t, err.Error(), "[Invalid AsyncAPI document] Check out err.ParsingErrors() for more information.")
	assert.Equal(t, len(err.ParsingErrors()), 1)
	assert.Equal(t, err.ParsingErrors()[0].Details()["property"], "id")
}

func TestParseJSON(t *testing.T) {
	asyncapi := []byte(`{
		"asyncapi": "2.0.0",
		"id": "myapi",
		"info": {
			"title": "My API",
			"version": "1.0.0"
		},
		"channels": {}
	}`)

	jsonDocument, err := ParseJSON(asyncapi)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(jsonDocument), `{"x-parser-messages":null,
		"asyncapi": "2.0.0",
		"id": "myapi",
		"info": {
			"title": "My API",
			"version": "1.0.0"
		},
		"channels": {}
	}`)
}

func TestParseJSONWithInvalidJSON(t *testing.T) {
	asyncapi := []byte(`"`)

	jsonDocument, err := ParseJSON(asyncapi)
	assert.Equal(t, err.Error(), "unexpected EOF")
	assert.Equal(t, len(err.ParsingErrors()), 0)
	assert.Equal(t, string(jsonDocument), ``)
}

func TestParseJSONWithEmptyJSON(t *testing.T) {
	asyncapi := []byte(``)

	jsonDocument, err := ParseJSON(asyncapi)
	assert.Equal(t, err.Error(), "EOF")
	assert.Equal(t, len(err.ParsingErrors()), 0)
	assert.Equal(t, string(jsonDocument), ``)
}

func TestParseJSONWithInvalidDocument(t *testing.T) {
	asyncapi := []byte(`{
		"asyncapi": "2.0.0",
		"info": {
			"title": "My API",
			"version": "1.0.0"
		},
		"channels": {}
	}`)

	jsonDocument, err := ParseJSON(asyncapi)
	assert.Equal(t, err.Error(), "[Invalid AsyncAPI document] Check out err.ParsingErrors() for more information.")
	assert.Equal(t, len(err.ParsingErrors()), 1)
	assert.Equal(t, err.ParsingErrors()[0].Details()["property"], "id")
	assert.Equal(t, string(jsonDocument), `{
		"asyncapi": "2.0.0",
		"info": {
			"title": "My API",
			"version": "1.0.0"
		},
		"channels": {}
	}`)
}

func TestParseBeautify(t *testing.T) {
	asyncapi := []byte(`{
		"asyncapi": "2.0.0",
		"id": "myapi",
		"info": {
			"title": "My API",
			"version": "1.0.0"
		},
		"channels": {
			"event/lighting/measured": {
				"subscribe": {
					"operationId": "receiveLightMeasurement",
					"message": {
						"name": "lightMeasured",
						"title": "Light measured",
						"contentType": "application/json",
						"payload": {
							"type": "object",
							"properties": {
								"lumens": {
									"type": "integer",
									"minimum": 0,
									"description": "Light intensity measured in lumens."
								}
							}
						}
					}
				}
			}
		}
	}`)

	beautifiedDoc, pErr := ParseJSON(asyncapi)
	assert.Check(t, is.Nil(pErr))

	var asyncAPI models.AsyncapiDocument
	err := json.Unmarshal(beautifiedDoc, &asyncAPI)
	assert.Check(t, is.Nil(err))

	xParserMessages := asyncAPI.Extensions["x-parser-messages"]
	var messageList models.ParserMessages
	json.Unmarshal(xParserMessages, &messageList)

	assert.Equal(t, len(messageList), 1)
	assert.Equal(t, messageList[0].ChannelName, "event/lighting/measured")
	assert.Equal(t, messageList[0].OperationName, "subscribe")
	assert.Equal(t, messageList[0].OperationId, "receiveLightMeasurement")
}
