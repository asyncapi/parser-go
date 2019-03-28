package models

import (
	"encoding/json"
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestOperationMessageUnmarshal(t *testing.T) {
	om := OperationMessage{}
	err := json.Unmarshal([]byte(`{
		"x-test": {"nested": "object"},
		"name": "lightMeasured",
		"title": "Light measured",
		"summary": "Inform about environmental lighting conditions for a particular streetlight.",
		"contentType": "application/json",
		"payload": {
			"type": "object",
			"properties": {
				"lumens": {
					"type": "integer",
					"minimum": 0,
					"description": "Light intensity measured in lumens."
				},
				"sentAt": {
					"type": "string",
					"format": "date-time",
					"description": "Date and time when the message was sent."
				}
			}
		}
	}`), &om)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, om.Name, "lightMeasured")
	assert.Equal(t, om.Title, "Light measured")
	assert.Equal(t, om.Summary, "Inform about environmental lighting conditions for a particular streetlight.")
	assert.Equal(t, om.ContentType, "application/json")
	assert.Equal(t, string(om.Payload), string(json.RawMessage(`{
			"type": "object",
			"properties": {
				"lumens": {
					"type": "integer",
					"minimum": 0,
					"description": "Light intensity measured in lumens."
				},
				"sentAt": {
					"type": "string",
					"format": "date-time",
					"description": "Date and time when the message was sent."
				}
			}
		}`)))
	assert.Equal(t, string(om.Extensions["x-test"]), `{"nested": "object"}`)
	assert.Assert(t, is.Nil(om.Extensions["invalid"]))
}

func TestOperationMessageMarshal(t *testing.T) {
	om := OperationMessage{
		Extensions: map[string]json.RawMessage{
			"x-test": json.RawMessage(`"test value"`),
		},
		Title:       "Light measured",
		Name:        "lightMeasured",
		Summary:     "Inform about environmental lighting conditions for a particular streetlight.",
		ContentType: "application/json",
		Payload:     json.RawMessage("{}"),
	}
	result, err := json.Marshal(om)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result), `{"x-test":"test value","contentType":"application/json","payload":{},"summary":"Inform about environmental lighting conditions for a particular streetlight.","name":"lightMeasured","title":"Light measured"}`)
}
