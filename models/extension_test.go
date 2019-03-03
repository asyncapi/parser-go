package models

import (
	"encoding/json"
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestExtensionsFromJSON(t *testing.T) {
	result, err := ExtensionsFromJSON([]byte(`{
		"not-extension": "should not be included",
		"x-test": "test value"
	}`))
	assert.Assert(t, is.Nil(err))
	assert.Assert(t, is.Nil(result["not-extension"]))
	assert.Equal(t, string(result["x-test"]), `"test value"`)
}

func TestExtensionsFromJSONWithObject(t *testing.T) {
	result, err := ExtensionsFromJSON([]byte(`{
		"x-test": {"one": {"two": 2}}
	}`))
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result["x-test"]), `{"one": {"two": 2}}`)
}

func TestExtensionsFromJSONWithArray(t *testing.T) {
	result, err := ExtensionsFromJSON([]byte(`{
		"x-test": [1, "one", true]
	}`))
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result["x-test"]), `[1, "one", true]`)
}

func TestExtensionsFromJSONWithNull(t *testing.T) {
	result, err := ExtensionsFromJSON([]byte(`{
		"x-test": null
	}`))
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result["x-test"]), `null`)
}

func TestMergeExtensions(t *testing.T) {
	info := Info{
		Title:   "Test",
		Version: "1.0.0",
	}
	jsonByteArray, err := json.Marshal(info)
	result, err := MergeExtensions(jsonByteArray, map[string]json.RawMessage{
		"x-test": json.RawMessage(`"test value"`),
	})
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result), `{"x-test":"test value","title":"Test","version":"1.0.0"}`)
}

func TestMergeExtensionsOnEmptyInfo(t *testing.T) {
	result, err := MergeExtensions([]byte(""), map[string]json.RawMessage{
		"x-test": json.RawMessage(`"test value"`),
	})
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result), `{"x-test":"test value"}`)
}

func TestMergeExtensionsOnNilInfo(t *testing.T) {
	result, err := MergeExtensions(nil, map[string]json.RawMessage{
		"x-test": json.RawMessage(`"test value"`),
	})
	assert.Equal(t, err.Error(), `jsonByteArray can't be nil`)
	assert.Assert(t, is.Nil(result))
}

func TestMergeExtensionsOnEmptyExtensions(t *testing.T) {
	result, err := MergeExtensions([]byte(`{"title":"Test"}`), map[string]json.RawMessage{})
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result), `{"title":"Test"}`)
}

func TestMergeExtensionsOnNilExtensions(t *testing.T) {
	result, err := MergeExtensions([]byte(`{"title":"Test"}`), nil)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result), `{"title":"Test"}`)
}

func TestMergeExtensionsOnEmptyValues(t *testing.T) {
	result, err := MergeExtensions([]byte(""), map[string]json.RawMessage{})
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result), `{}`)
}

func TestMergeExtensionsOnNilValues(t *testing.T) {
	result, err := MergeExtensions(nil, nil)
	assert.Equal(t, err.Error(), `jsonByteArray can't be nil`)
	assert.Assert(t, is.Nil(result))
}
