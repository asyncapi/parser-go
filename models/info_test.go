package models

import (
	"encoding/json"
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestInfoUnmarshal(t *testing.T) {
	info := Info{}
	err := json.Unmarshal([]byte(`{
		"title":"my API",
		"contact":{
			"name":"Fran",
			"email":"fmvilas@gmail.com"
		},
		"license": {
			"name": "Apache 2.0"
		},
		"x-test": {"nested": "object"},
		"invalid": "invalid value"
	}`), &info)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, info.Title, "my API")
	assert.Equal(t, info.Contact.Name, "Fran")
	assert.Equal(t, info.Contact.Email, "fmvilas@gmail.com")
	assert.Equal(t, info.License.Name, "Apache 2.0")
	assert.Equal(t, string(info.Extensions["x-test"]), `{"nested": "object"}`)
	assert.Assert(t, is.Nil(info.Extensions["invalid"]))
}

func TestInfoMarshal(t *testing.T) {
	info := Info{
		Extensions: map[string]json.RawMessage{
			"x-test": json.RawMessage(`"test value"`),
		},
		Title:   "My API",
		Version: "1.0.0",
	}
	result, err := json.Marshal(info)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result), `{"x-test":"test value","title":"My API","version":"1.0.0"}`)
}

func TestInfoContactUnmarshal(t *testing.T) {
	info := Info{}
	err := json.Unmarshal([]byte(`{"title":"my API", "contact": { "name": "Fran" } }`), &info)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, len(info.Extensions), 0)
	assert.Equal(t, info.Contact.Name, "Fran")
}

func TestInfoContactMarshal(t *testing.T) {
	info := Info{
		Title:   "My API",
		Version: "1.0.0",
		Contact: &Contact{
			Name: "Fran",
		},
	}
	result, err := json.Marshal(info)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result), `{"title":"My API","contact":{"name":"Fran"},"version":"1.0.0"}`)
}

func TestInfoLicenseUnmarshal(t *testing.T) {
	info := Info{}
	err := json.Unmarshal([]byte(`{"title":"my API", "license": { "name": "Apache" } }`), &info)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, len(info.Extensions), 0)
	assert.Equal(t, info.License.Name, "Apache")
}

func TestInfoLicenseMarshal(t *testing.T) {
	info := Info{
		Title:   "My API",
		Version: "1.0.0",
		License: &License{
			Name: "Apache",
		},
	}
	result, err := json.Marshal(info)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result), `{"title":"My API","license":{"name":"Apache"},"version":"1.0.0"}`)
}
