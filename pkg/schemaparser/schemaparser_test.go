package schemaparser

import (
	"testing"

	"github.com/asyncapi/parser/pkg/hlsp"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestParse(t *testing.T) {
	asyncapi := []byte(`{
		"asyncapi": "2.0.0-rc1",
		"id": "urn:myapi",
		"info": {
			"title": "My API",
			"version": "1.0.0"
		},
		"channels": {
			"event/lighting/measured": {
				"subscribe": {
					"operationId": "receiveLightMeasurement",
					"message": {
						"schemaFormat": "protobuf",
						"name": "lightMeasured",
						"title": "Light measured",
						"contentType": "application/json",
						"payload": {
							"type": "object",
							"properties": {
								"lumens": {
									"type": "integer",
									"minimum": 0,
									"description": "Light intensity measured in lumens."
								}
							}
						}
					}
				}
			},
			"test": {
				"publish": {
					"operationId": "receiveLightMeasurement",
					"message": {
						"oneOf": [
							{
								"schemaFormat": "application/vnd.asyncapi",
								"name": "lightMeasured",
								"title": "Light measured",
								"contentType": "application/json",
								"payload": {
									"type": "object",
									"properties": {
										"lumens": {
											"type": "integer",
											"minimum": 0,
											"description": "Light intensity measured in lumens."
										}
									}
								}
							},
							{
								"name": "lightMeasured",
								"title": "Light measured",
								"contentType": "application/json",
								"payload": {
									"type": "object",
									"properties": {
										"lumens": {
											"type": "integer",
											"minimum": 0,
											"description": "Light intensity measured in lumens."
										}
									}
								}
							}
						]
					}
				}
			}
		}
	}`)

	doc, err := hlsp.Parse(asyncapi, true)
	assert.Assert(t, is.Nil(err))

	err = Parse(doc)
	assert.Assert(t, is.Nil(err))
}
