package models

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestUnmarshalExtensions(t *testing.T) {
	info := Info{}
	result, err := UnmarshalExtensions([]byte(`{"x-test": "test value"}`))
	info.Extensions = result
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, len(info.Extensions), 1)
	assert.Equal(t, info.Extensions["x-test"], "test value")
}

func TestMarshalWithExtensions(t *testing.T) {
	info := Info{
		Extensions: Extensions{
			"x-test": "test value",
		},
		Title:   "Test",
		Version: "1.0.0",
	}
	result, err := MarshalWithExtensions(info, info.Extensions)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result), `{"x-test":"test value","title":"Test","version":"1.0.0"}`)
}

func TestMarshalWithExtensionsOnEmptyInfo(t *testing.T) {
	info := Info{
		Extensions: Extensions{
			"x-test": "test value",
		},
	}
	result, err := MarshalWithExtensions(Info{}, info.Extensions)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result), `{"x-test":"test value"}`)
}

func TestMarshalWithExtensionsOnNilInfo(t *testing.T) {
	info := Info{
		Extensions: Extensions{
			"x-test": "test value",
		},
	}
	result, err := MarshalWithExtensions(nil, info.Extensions)
	assert.Equal(t, err.Error(), `Object can't be nil.`)
	assert.Assert(t, is.Nil(result))
}

func TestMarshalWithExtensionsOnEmptyExtensions(t *testing.T) {
	info := Info{
		Extensions: Extensions{},
		Title:      "Test",
	}
	result, err := MarshalWithExtensions(info, info.Extensions)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result), `{"title":"Test"}`)
}

func TestMarshalWithExtensionsOnNilExtensions(t *testing.T) {
	info := Info{
		Title: "Test",
	}
	result, err := MarshalWithExtensions(info, nil)
	assert.Equal(t, err.Error(), `Extensions can't be nil.`)
	assert.Assert(t, is.Nil(result))
}

func TestMarshalWithExtensionsOnEmptyValues(t *testing.T) {
	info := Info{
		Extensions: Extensions{},
	}
	result, err := MarshalWithExtensions(Info{}, info.Extensions)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result), `{}`)
}

func TestMarshalWithExtensionsOnNilValues(t *testing.T) {
	result, err := MarshalWithExtensions(nil, nil)
	assert.Equal(t, err.Error(), `Extensions can't be nil.`)
	assert.Assert(t, is.Nil(result))
}
