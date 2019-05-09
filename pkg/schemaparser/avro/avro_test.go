package avro

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/asyncapi/parser/pkg/dereferencer"
	"github.com/xeipuuv/gojsonschema"

	"github.com/linkedin/goavro/v2"
	"github.com/stretchr/testify/assert"
)

// TODO: improve all tests

func getAvroSchemaForTest() []byte {
	return []byte(`
	{
		"type" : "record",
		"name" : "twitter_schema",
		"namespace" : "com.miguno.avro",
		"fields" : [ {
		  "name" : "username",
		  "type" : "string",
		  "doc" : "Name of the user account on Twitter.com"
		}, {
		  "name" : "tweet",
		  "type" : "string",
		  "doc" : "The content of the user's Twitter message"
		}, {
		  "name" : "timestamp",
		  "type" : "long",
		  "doc" : "Unix epoch time in seconds"
		} ],
		"doc:" : "A basic schema for storing Twitter messages"
	  }
`)
}

func getUnionAvroSchemaForTest() []byte {
	return []byte(`
	[
	{ "type": "string" },
	"null",
	{
			"type": "map",
			"values": {
					"type": "enum",
					"name": "gender",
					"symbols": [ "MALE", "FEMALE", "WHOCARES" ]
			}
	}
]
`)
}

func getSimpleArrayAvroSchemaForTest() []byte {
	return []byte(`
	{
		"type": "array",
		"items": "boolean"
}
`)
}

func getMapArrayAvroSchemaForTest() []byte {
	return []byte(`
	{
		"type": "array",
		"items": {
				"type": "map",
				"values": "bytes"
		}
}
`)
}

func getEnumAvroSchemaForTest() []byte {
	return []byte(`
	{
		"name": "gender",
		"type": "enum",
		"symbols": [ "MALE", "FEMALE", "WHOCARES" ]
	}
`)
}

func getFixedAvroSchemaForTest() []byte {
	return []byte(`
	{
		"type": "fixed",
		"name": "md5",
		"size": 16
	}
`)
}

func getFixedMapAvroSchemaForTest() []byte {
	return []byte(`
	{
		"type": "map",
		"values": {
				"type": "fixed",
				"name": "md5",
				"size": 16
		}
	}
`)
}

func getSimpleAvroSchemaForTest() [][]byte {
	input := []string{`{ "type": "boolean" }`, `{ "type": "null" }`, `{ "type": "int" }`, `{ "type": "float" }`, `{ "type": "long" }`}
	output := make([][]byte, len(input))
	for i, v := range input {
		output[i] = []byte(v)
	}
	return output
}

func getSimpleMapAvroSchemaForTest() []byte {
	return []byte(`
	{
		"type": "map",
		"values": "boolean"
	}
`)
}

func getComposedMapAvroSchemaForTest() []byte {
	return []byte(`
	{
		"type": "map",
		"values": {
				"type": "map",
				"values": "bytes"
		}
}
`)
}

func getJSONSchemaForTest() []byte {
	return []byte(`
	{
			"definitions" : {
			  "record:com.miguno.avro.twitter_schema" : {
				"type" : "object",
				"required" : [ "username", "tweet", "timestamp" ],
				"additionalProperties" : false,
				"properties" : {
				  "username" : {
					"type" : "string"
				  },
				  "tweet" : {
					"type" : "string"
				  },
				  "timestamp" : {
					"type" : "integer",
					"minimum" : -9223372036854775808,
					"maximum" : 9223372036854775807
				  }
				}
			  }
			},
			"$ref" : "#/definitions/record:com.miguno.avro.twitter_schema"
		  }
`)
}

func TestAvro2Json(t *testing.T) {
	avroSchema := getAvroSchemaForTest()
	codec, err := goavro.NewCodec(string(avroSchema))
	assert.NoError(t, err)
	bschema := json.RawMessage(codec.Schema())
	err = Parse(&bschema)
	log.Printf("Avro schema: %s", bschema)
	assert.Contains(t, string(bschema), `"record:com.miguno.avro:twitter_schema": {"type":"object","additionalProperties":false,"required":["username","tweet","timestamp"],"properties":{"timestamp":{"type":"integer"},"tweet":{"type":"string"},"username":{"type":"string"}}}`)
	resolvedDoc, _ := dereferencer.Dereference(bschema, true)
	assert.NoError(t, checkJSONSchema(string(resolvedDoc)))
}

// func TestUnionAvro2Json(t *testing.T) {
// 	avroSchema := getUnionAvroSchemaForTest()
// 	codec, err := goavro.NewCodec(string(avroSchema))
// 	assert.NoError(t, err)
// 	// log.Printf("Avro schema: %s", codec.Schema())
// 	bschema := json.RawMessage(codec.Schema())
// 	err = Parse(&bschema)
// 	log.Printf("Avro schema: %s", bschema)
// 	assert.Contains(t, string(bschema), getJSONSchemaForTest())
//	assert.NoError(t, checkJSONSchema(string(bschema)))
// }

func TestSimpleArrayAvro2Json(t *testing.T) {
	avroSchema := getSimpleArrayAvroSchemaForTest()
	codec, err := goavro.NewCodec(string(avroSchema))
	assert.NoError(t, err)
	// log.Printf("Avro schema: %s", codec.Schema())
	bschema := json.RawMessage(codec.Schema())
	err = Parse(&bschema)
	log.Printf("Avro schema: %s", bschema)
	assert.Contains(t, string(bschema), `{"type":"array","items":{"type":"boolean"}}`)
	resolvedDoc, _ := dereferencer.Dereference(bschema, true)
	assert.NoError(t, checkJSONSchema(string(resolvedDoc)))
}

func TestMapArrayAvro2Json(t *testing.T) {
	avroSchema := getMapArrayAvroSchemaForTest()
	codec, err := goavro.NewCodec(string(avroSchema))
	assert.NoError(t, err)
	// log.Printf("Avro schema: %s", codec.Schema())
	bschema := json.RawMessage(codec.Schema())
	err = Parse(&bschema)
	log.Printf("Avro schema: %s", bschema)
	assert.Contains(t, string(bschema), `{"type":"array","items":{"type":"object","additionalProperties":{"type":"string","pattern":"^[\u0000-ÿ]*$"}}}`)
	assert.NoError(t, checkJSONSchema(string(bschema)))
}

func TestEnumAvro2Json(t *testing.T) {
	avroSchema := getEnumAvroSchemaForTest()
	codec, err := goavro.NewCodec(string(avroSchema))
	assert.NoError(t, err)
	bschema := json.RawMessage(codec.Schema())
	err = Parse(&bschema)
	log.Printf("Avro schema: %s", bschema)
	assert.Contains(t, string(bschema), "enum:gender")
	assert.Contains(t, string(bschema), `{"enum":["MALE","FEMALE","WHOCARES"]}`)
	assert.Contains(t, string(bschema), `"$ref" : "#/definitions/enum:gender"`)
	assert.NoError(t, checkJSONSchema(string(bschema)))
}

func TestFixedAvro2Json(t *testing.T) {
	avroSchema := getFixedAvroSchemaForTest()
	codec, err := goavro.NewCodec(string(avroSchema))
	assert.NoError(t, err)
	bschema := json.RawMessage(codec.Schema())
	err = Parse(&bschema)
	log.Printf("Avro schema: %s", bschema)
	assert.Contains(t, string(bschema), `fixed:md5`)
	assert.NoError(t, checkJSONSchema(string(bschema)))
}

func TestFixedMapAvro2Json(t *testing.T) {
	avroSchema := getFixedMapAvroSchemaForTest()
	codec, err := goavro.NewCodec(string(avroSchema))
	assert.NoError(t, err)
	bschema := json.RawMessage(codec.Schema())
	err = Parse(&bschema)
	log.Printf("Avro schema: %s", bschema)
	assert.Contains(t, string(bschema), `fixed:md5`)
}

// TODO: improve tests
func TestSimpleAvro2Json(t *testing.T) {
	avroSchema := getSimpleAvroSchemaForTest()
	for _, v := range avroSchema {
		codec, err := goavro.NewCodec(string(v))
		assert.NoError(t, err)
		bschema := json.RawMessage(codec.Schema())
		err = Parse(&bschema)
		log.Printf("Avro schema: %s", bschema)
		assert.Contains(t, string(bschema), `type`)
		assert.NoError(t, checkJSONSchema(string(bschema)))
	}
}

func TestMapAvro2Json(t *testing.T) {
	avroSchema := getSimpleMapAvroSchemaForTest()
	codec, err := goavro.NewCodec(string(avroSchema))
	assert.NoError(t, err)
	bschema := json.RawMessage(codec.Schema())
	err = Parse(&bschema)
	log.Printf("Avro schema: %s", bschema)
	assert.Contains(t, string(bschema), `type`)
	assert.NoError(t, checkJSONSchema(string(bschema)))
}

func TestComposeMapAvro2Json(t *testing.T) {
	avroSchema := getComposedMapAvroSchemaForTest()
	codec, err := goavro.NewCodec(string(avroSchema))
	assert.NoError(t, err)
	bschema := json.RawMessage(codec.Schema())
	err = Parse(&bschema)
	log.Printf("Avro schema: %s", bschema)
	assert.Contains(t, string(bschema), `{"type":"object","additionalProperties":{"type":"object","additionalProperties":{"type":"string","pattern":"^[\u0000-ÿ]*$"}}}`)
	assert.NoError(t, checkJSONSchema(string(bschema)))
}

func checkJSONSchema(schema string) error {
	sl := gojsonschema.NewSchemaLoader()
	loader := gojsonschema.NewStringLoader(string(schema))
	_, err := sl.Compile(loader)
	return err
}
