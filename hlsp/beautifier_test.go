package hlsp

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/asyncapi/parser/models"
	"gotest.tools/assert"
)

func TestBeautify(t *testing.T) {
	jsonFile, err := os.Open("../asyncapi/2.0.0/example.json")
	if err != nil {
		t.Log(err)
		return
	}
	defer jsonFile.Close()

	fileBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		t.Log(err)
		return
	}
	var beautifiedDoc json.RawMessage
	if beautifiedDoc, err = Beautify(fileBytes); err != nil {
		t.Log(err)
		return
	}

	var asyncAPI models.AsyncapiDocument
	err = json.Unmarshal(beautifiedDoc, &asyncAPI)
	if err != nil {
		t.Log(err)
		return
	}
	xParserMessages := asyncAPI.Extensions["x-parser-messages"]

	var messageList ParserMessages
	json.Unmarshal(xParserMessages, &messageList)

	assert.Equal(t, len(messageList), 4)

	t.Log(messageList)
}
