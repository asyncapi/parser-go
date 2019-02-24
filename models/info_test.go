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

func TestInfoMarshal(t *testing.T) {
	info := Info{
		Title:   "My API",
		Version: "1.0.0",
	}
	result, err := info.Marshal()
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result), `{"title":"My API","version":"1.0.0"}`)
}

func TestInfoContactUnmarshal(t *testing.T) {
	info := Info{}
	err := info.Unmarshal([]byte(`{"title":"my API", "contact": { "name": "Fran" } }`))
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, len(info.ExtensionProps.Extensions), 0)
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
	result, err := info.Marshal()
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result), `{"title":"My API","contact":{"name":"Fran"},"version":"1.0.0"}`)
}

func TestInfoLicenseUnmarshal(t *testing.T) {
	info := Info{}
	err := info.Unmarshal([]byte(`{"title":"my API", "license": { "name": "Apache" } }`))
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, len(info.ExtensionProps.Extensions), 0)
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
	result, err := info.Marshal()
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result), `{"title":"My API","license":{"name":"Apache"},"version":"1.0.0"}`)
}
