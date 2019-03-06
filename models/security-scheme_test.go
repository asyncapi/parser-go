package models

import (
	"encoding/json"
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestSecuritySchemeUnmarshal(t *testing.T) {
	ss := SecurityScheme{}
	err := json.Unmarshal([]byte(`{
		"x-test": {"nested": "object"},
		"type": "oauth2",
		"description": "Flows to support OAuth 2.0",
		"flows": {
			"implicit": {
				"authorizationUrl": "https://authserver.example/auth",
				"scopes": {
					"streetlights:on": "Ability to switch lights on",
					"streetlights:off": "Ability to switch lights off",
					"streetlights:dim": "Ability to dim the lights"
				}
			},
			"password": {
				"tokenUrl": "https://authserver.example/token",
				"scopes": {
					"streetlights:on": "Ability to switch lights on",
					"streetlights:off": "Ability to switch lights off",
					"streetlights:dim": "Ability to dim the lights"
				}
			},
			"clientCredentials": {
				"tokenUrl": "https://authserver.example/token",
				"scopes": {
					"streetlights:on": "Ability to switch lights on",
					"streetlights:off": "Ability to switch lights off",
					"streetlights:dim": "Ability to dim the lights"
				}
			},
			"authorizationCode": {
				"authorizationUrl": "https://authserver.example/auth",
				"tokenUrl": "https://authserver.example/token",
				"refreshUrl": "https://authserver.example/refresh",
				"scopes": {
					"streetlights:on": "Ability to switch lights on",
					"streetlights:off": "Ability to switch lights off",
					"streetlights:dim": "Ability to dim the lights"
				}
			}
		}
	}`), &ss)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, ss.Type, "oauth2")
	assert.Equal(t, ss.Description, "Flows to support OAuth 2.0")
	assert.Equal(t, string(ss.Flows), string(json.RawMessage(`{
			"implicit": {
				"authorizationUrl": "https://authserver.example/auth",
				"scopes": {
					"streetlights:on": "Ability to switch lights on",
					"streetlights:off": "Ability to switch lights off",
					"streetlights:dim": "Ability to dim the lights"
				}
			},
			"password": {
				"tokenUrl": "https://authserver.example/token",
				"scopes": {
					"streetlights:on": "Ability to switch lights on",
					"streetlights:off": "Ability to switch lights off",
					"streetlights:dim": "Ability to dim the lights"
				}
			},
			"clientCredentials": {
				"tokenUrl": "https://authserver.example/token",
				"scopes": {
					"streetlights:on": "Ability to switch lights on",
					"streetlights:off": "Ability to switch lights off",
					"streetlights:dim": "Ability to dim the lights"
				}
			},
			"authorizationCode": {
				"authorizationUrl": "https://authserver.example/auth",
				"tokenUrl": "https://authserver.example/token",
				"refreshUrl": "https://authserver.example/refresh",
				"scopes": {
					"streetlights:on": "Ability to switch lights on",
					"streetlights:off": "Ability to switch lights off",
					"streetlights:dim": "Ability to dim the lights"
				}
			}
		}`)))
	assert.Equal(t, string(ss.Extensions["x-test"]), `{"nested": "object"}`)
	assert.Assert(t, is.Nil(ss.Extensions["invalid"]))
}

func TestSecuritySchemeMarshal(t *testing.T) {
	ss := SecurityScheme{
		Extensions: map[string]json.RawMessage{
			"x-test": json.RawMessage(`"test value"`),
		},
		Type:        "oauth2",
		Description: "Flows to support OAuth 2.0",
		Flows:       json.RawMessage("{}"),
		In:          "header",
	}
	result, err := json.Marshal(ss)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result), `{"x-test":"test value","type":"oauth2","description":"Flows to support OAuth 2.0","in":"header","flows":{}}`)
}

func TestHttpSecuritySchemeUnmarshal(t *testing.T) {
	hss := HttpSecurityScheme{}
	err := json.Unmarshal([]byte(`{
		"x-test": {"nested": "object"},
		"scheme": "a",
		"description": "b",
		"type": "c",
		"bearerFormat": "d",
		"name": "e",
		"in": "f"
	}`), &hss)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, hss.Scheme, "a")
	assert.Equal(t, hss.Description, "b")
	assert.Equal(t, hss.Type, "c")
	assert.Equal(t, hss.BearerFormat, "d")
	assert.Equal(t, hss.Name, "e")
	assert.Equal(t, hss.In, "f")
	assert.Equal(t, string(hss.Extensions["x-test"]), `{"nested": "object"}`)
	assert.Assert(t, is.Nil(hss.Extensions["invalid"]))
}

func TestHttpSecuritySchemeMarshal(t *testing.T) {
	ss := HttpSecurityScheme{
		Extensions: map[string]json.RawMessage{
			"x-test": json.RawMessage(`"test value"`),
		},
		Scheme:       "a",
		Description:  "b",
		Type:         "c",
		BearerFormat: "d",
		Name:         "e",
		In:           "f",
	}
	result, err := json.Marshal(ss)
	assert.Assert(t, is.Nil(err))
	assert.Equal(t, string(result), `{"x-test":"test value","scheme":"a","description":"b","type":"c","bearerFormat":"d","name":"e","in":"f"}`)
}
