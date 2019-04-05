package schemaparser

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestParse(t *testing.T) {
	asyncapi := []byte(`{
		"asyncapi": "2.0.0",
		"id": "myapi",
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

	err := Parse(asyncapi)
	assert.Assert(t, is.Nil(err))
}
