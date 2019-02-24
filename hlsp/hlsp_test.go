package hlsp

import (
	"testing"

	"gotest.tools/assert"
    is "gotest.tools/assert/cmp"
)

func TestParse(t *testing.T) {
	result, err := Parse(`{
		"asyncapi":"2.0.0",
		"id": "my-id",
		"info": {
			"title": "Test API",
			"version": "1.0.0"
		},
		"channels": {}
	}`)

	assert.Assert(t, is.Nil(err))
	assert.Equal(t, result, true)
}