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

	// t.Log(messageList)
}

func TestBeautifyOneChannel(t *testing.T) {
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
	var beautifiedDoc json.RawMessage
	var err error

	if beautifiedDoc, err = Beautify(asyncapi); err != nil {
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

	assert.Equal(t, len(messageList), 1)
	assert.Equal(t, messageList[0].ChannelName, "event/lighting/measured")
	assert.Equal(t, messageList[0].OperationName, "subscribe")
	assert.Equal(t, messageList[0].OperationId, "receiveLightMeasurement")
}
