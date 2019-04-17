package dereferencer

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

const (
refInFile = `{
"event/{streetlightId}/lighting/measured": {
	"parameters": [
			{
				"$ref": "#/components/parameters/streetlightId"
			}
		]
	},
	"components": {
		"parameters": {
			"streetlightId": {
				"name": "streetlightId",
				"description": "The ID of the streetlight.",
				"schema": {
					"type": "string"
				}
			}
		}
	}
}
	`
expectedResolved = `
"event/{streetlightId}/lighting/measured": {
	"parameters": [
			{
				"name": "streetlightId",
				"description": "The ID of the streetlight.",
				"schema": {
				"type": "string"
			}
		}
	]
}
	`
)

func TestDereferenceInFile(t *testing.T) {
	jsonDocument, err := Dereference([]byte(refInFile))
	assert.NoError(t, err, "error dereferencing")
	assert.Equal(t, expectedResolved, string(jsonDocument), "No equal")
}
