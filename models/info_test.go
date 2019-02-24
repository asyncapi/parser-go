package models

import (
	"testing"

	"gotest.tools/assert"
    is "gotest.tools/assert/cmp"
)

func TestInfoUnmarshal(t *testing.T) {
	info := Info{}
	err := info.Unmarshal([]byte(`{"title":"my API", "x-test": "test value"}`))
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, len(info.ExtensionProps.Extensions), 0)
}