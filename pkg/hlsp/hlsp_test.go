package hlsp

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestParse(t *testing.T) {
	asyncapi := []byte(`
asyncapi: '2.0.0-rc1'
id: 'urn:myapi'
info:
  title: My API
  version: '1.0.0'
channels: {}`)

	doc, err := Parse(asyncapi, true)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, doc.Asyncapi, "2.0.0-rc1")
	assert.Equal(t, doc.Id, "urn:myapi")
}
func TestParseWithEmptyYAML(t *testing.T) {
	asyncapi := []byte(``)

	_, err := Parse(asyncapi, true)
	assert.Equal(t, err.Error(), "[Invalid AsyncAPI document] Document is empty or null.")
	assert.Equal(t, len(err.ParsingErrors), 0)
}
func TestParseWithInvalidDocument(t *testing.T) {
	asyncapi := []byte(`
asyncapi: '2.0.0-rc1'
info:
  title: My API
  version: '1.0.0'
channels: {}`)

	_, err := Parse(asyncapi, true)
	assert.Equal(t, err.Error(), "[Invalid AsyncAPI document] Check out err.ParsingErrors() for more information.")
	assert.Equal(t, len(err.ParsingErrors), 1)
	assert.Equal(t, err.ParsingErrors[0].Details()["property"], "id")
}

func TestParseJSON(t *testing.T) {
	asyncapi := []byte(`{
		"asyncapi": "2.0.0-rc1",
		"id": "urn:myapi",
		"info": {
			"title": "My API",
			"version": "1.0.0"
		},
		"channels": {}
	}`)

	jsonDocument, err := ParseJSON(asyncapi, true)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(jsonDocument), `{
		"asyncapi": "2.0.0-rc1",
		"id": "urn:myapi",
		"info": {
			"title": "My API",
			"version": "1.0.0"
		},
		"channels": {}
	}`)
}

func TestParseJSONWithInvalidJSON(t *testing.T) {
	asyncapi := []byte(`"`)

	jsonDocument, err := ParseJSON(asyncapi, true)
	assert.Equal(t, err.Error(), "unexpected EOF")
	assert.Equal(t, len(err.ParsingErrors), 0)
	assert.Equal(t, string(jsonDocument), ``)
}

func TestParseJSONWithEmptyJSON(t *testing.T) {
	asyncapi := []byte(``)

	jsonDocument, err := ParseJSON(asyncapi, true)
	assert.Equal(t, err.Error(), "EOF")
	assert.Equal(t, len(err.ParsingErrors), 0)
	assert.Equal(t, string(jsonDocument), ``)
}

func TestParseJSONWithInvalidDocument(t *testing.T) {
	asyncapi := []byte(`{
		"asyncapi": "2.0.0-rc1",
		"info": {
			"title": "My API",
			"version": "1.0.0"
		},
		"channels": {}
	}`)

	jsonDocument, err := ParseJSON(asyncapi, true)
	assert.Equal(t, err.Error(), "[Invalid AsyncAPI document] Check out err.ParsingErrors() for more information.")
	assert.Equal(t, len(err.ParsingErrors), 1)
	assert.Equal(t, err.ParsingErrors[0].Details()["property"], "id")
	assert.Equal(t, string(jsonDocument), `{
		"asyncapi": "2.0.0-rc1",
		"info": {
			"title": "My API",
			"version": "1.0.0"
		},
		"channels": {}
	}`)
}
