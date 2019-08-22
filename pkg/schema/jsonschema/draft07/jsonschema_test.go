package draft07

import (
	. "github.com/onsi/gomega"

	"encoding/json"
	"testing"
)

func TestParseInvalidSchemaObject(t *testing.T) {
	g := NewWithT(t)
	v := "this will not work"
	err := Parse(v)
	g.Expect(err).Should(HaveOccurred())
}

func TestParse(t *testing.T) {
	tests := []struct {
		name string
		v    interface{}
	}{
		{
			name: "boolean value",
			v:    true,
		},
		{
			name: "object value",
			v:    validTestPayload(t),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := Parse(&test.v)
			NewWithT(t).Expect(err).ShouldNot(HaveOccurred())
		})
	}
}

func validTestPayload(t *testing.T) interface{} {
	var v interface{}
	payload := []byte(`{
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
	err := json.Unmarshal(payload, &v)
	if err != nil {
		t.Fatal("data unmarshalling error")
	}
	return v
}
