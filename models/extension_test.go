package models

import (
	"testing"

	"gotest.tools/assert"
    is "gotest.tools/assert/cmp"
)

func TestUnmarshalExtensions(t *testing.T) {
	info := Info{
		Title: "My API",
	}
	err := UnmarshalExtensions(&info.ExtensionProps, []byte(`{"x-test": "test value"}`))
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, len(info.ExtensionProps.Extensions), 1)
	assert.Equal(t, info.ExtensionProps.Extensions["x-test"], "test value")
}