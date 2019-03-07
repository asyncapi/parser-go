package hlsp

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/asyncapi/parser/models"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

// TODO: Remove this test, as it's not a unit test.
func Test(t *testing.T) {
	yamlFile, err := os.Open("../asyncapi/2.0.0/example.yaml")
	if err != nil {
		t.Log(err)
		return
	}
	defer yamlFile.Close()

	fileBytes, err := ioutil.ReadAll(yamlFile)
	if err != nil {
		t.Log(err)
		return
	}

	jsonFile, e := Parse(fileBytes)
	if e != nil {
		t.Log(e.Error())
		t.Log(e.ParsingErrors())
		return
	}

	var AsyncAPI models.AsyncapiDocument
	err = json.Unmarshal(jsonFile, &AsyncAPI)
	if err != nil {
		t.Log(err)
		return
	}
	j, _ := AsyncAPI.Channels["event/{streetlightId}/lighting/measured"].Subscribe.Message.Payload.MarshalJSON()
	t.Log(string(j))
}

func TestParse(t *testing.T) {
	asyncapi := []byte(`
asyncapi: '2.0.0'
id: myapi
info:
  title: My API
  version: '1.0.0'
channels: {}`)

	jsonDocument, err := Parse(asyncapi)
	t.Log(err)
	// assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(jsonDocument), `{"asyncapi":"2.0.0","channels":{},"id":"myapi","info":{"title":"My API","version":"1.0.0"}}`)
}
func TestParseWithEmptyYAML(t *testing.T) {
	asyncapi := []byte(``)

	jsonDocument, err := ParseJSON(asyncapi)
	assert.Equal(t, err.Error(), "EOF")
	assert.Equal(t, len(err.ParsingErrors()), 0)
	assert.Equal(t, string(jsonDocument), ``)
}
func TestParseWithInvalidDocument(t *testing.T) {
	asyncapi := []byte(`
asyncapi: '2.0.0'
info:
  title: My API
  version: '1.0.0'
channels: {}`)

	jsonDocument, err := Parse(asyncapi)
	assert.Equal(t, err.Error(), "[Invalid AsyncAPI document] Check out err.ParsingErrors() for more information.")
	assert.Equal(t, len(err.ParsingErrors()), 1)
	assert.Equal(t, err.ParsingErrors()[0].Details()["property"], "id")
	assert.Equal(t, string(jsonDocument), `{"asyncapi":"2.0.0","channels":{},"info":{"title":"My API","version":"1.0.0"}}`)
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
	assert.Equal(t, string(jsonDocument), `{
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
