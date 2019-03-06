package models

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func EqualJSON(s1 []byte, s2 string) bool {
	var o1 interface{}
	var o2 interface{}

	json.Unmarshal([]byte(s1), &o1)
	json.Unmarshal([]byte(s2), &o2)

	return reflect.DeepEqual(o1, o2)
}

func TestServerUnmarshal(t *testing.T) {
	server := &Server{}
	err := server.UnmarshalJSON([]byte(`
	{
		"url":"api.streetlights.com", 
		"description": "my description",
		"scheme": "mqtt",
		"schemeVersion": "0.9.1",
		"baseChannel": "smartylighting/streetlights/1/0"
	}
	`))
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, server.Url, "api.streetlights.com")
	assert.Equal(t, server.Description, "my description")
	assert.Equal(t, server.Scheme, "mqtt")
	assert.Equal(t, server.SchemeVersion, "0.9.1")
	assert.Equal(t, server.BaseChannel, "smartylighting/streetlights/1/0")
}

func TestServerMarshal(t *testing.T) {
	server := Server{
		Url:           "api.streetlights.com",
		Scheme:        "mqtt",
		SchemeVersion: "0.9.1",
	}
	result, err := server.MarshalJSON()
	fmt.Println(string(result))
	assert.Assert(t, is.Nil(err))
	assert.Assert(t, EqualJSON(result, `
	{
		"scheme":"mqtt",
		"schemeVersion":"0.9.1",
		"url":"api.streetlights.com",
		"baseChannel":"",
		"description":"",
		"security":null,
		"variables":null
	}`))
}

func TestServerWithExtensionsMarshal(t *testing.T) {
	server := Server{
		Extensions: Extensions{
			"x-test": "test value",
		},
		Url:           "api.streetlights.com",
		Scheme:        "mqtt",
		SchemeVersion: "0.9.1",
	}
	result, err := server.MarshalJSON()
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result), `{"x-test":"test value","url":"api.streetlights.com","scheme":"mqtt","schemeVersion":"0.9.1"}`)
}

func TestServerVariablesUnmarshal(t *testing.T) {
	server := &Server{}
	err := server.UnmarshalJSON([]byte(`
	{
		"url":"api.streetlights.com", 
		"scheme": "mqtt",
		"variables": {
			"port": {
				"description": "Secure connection",
				"default": "1883",
				"enum": [
					"1883",
					"8883"
				]
			}
		}
	}
	`))
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, server.Url, "api.streetlights.com")
	assert.Equal(t, server.Variables.AdditionalProperties["port"].Description, "Secure connection")
	assert.Equal(t, server.Variables.AdditionalProperties["port"].Default, "1883")
	assert.DeepEqual(t, server.Variables.AdditionalProperties["port"].Enum, []string{"1883", "8883"})
}

func TestServerVariablesMarshal(t *testing.T) {
	additionalProperties := make(map[string]*ServerVariable)
	additionalProperties["port"] = &ServerVariable{
		Default:     "1883",
		Description: "Secure connection",
	}

	server := Server{
		Extensions: Extensions{
			"x-test": "test value",
		},
		Url:    "api.streetlights.com",
		Scheme: "mqtt",
		Variables: &ServerVariables{
			AdditionalProperties: additionalProperties,
		},
	}
	result, err := server.MarshalJSON()
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result), `{"x-test":"test value","url":"api.streetlights.com","scheme":"mqtt","variables":{"port":{"default":"1883","description":"Secure connection"}}}`)
}

func TestServerSecurityUnmarshal(t *testing.T) {
	server := &Server{}
	err := server.UnmarshalJSON([]byte(`
	{
		"url":"api.streetlights.com", 
		"scheme": "mqtt",
		"variables": {
			"port": {
				"description": "Secure connection",
				"default": "1883",
				"enum": [
					"1883",
					"8883"
				]
			}
		}
	}
	`))
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, server.Url, "api.streetlights.com")
	assert.Equal(t, server.Variables.AdditionalProperties["port"].Description, "Secure connection")
	assert.Equal(t, server.Variables.AdditionalProperties["port"].Default, "1883")
	assert.DeepEqual(t, server.Variables.AdditionalProperties["port"].Enum, []string{"1883", "8883"})
}
