package avro

import (
	"encoding/json"
	"fmt"

	"github.com/asyncapi/parser/pkg/errs"
)

// MapAvro maps a avro map object
type MapAvro struct {
	Content string `json:"-"`
}

// Convert transforms avro formatted message to JSONSchema
func (ra *MapAvro) Convert(message map[string]interface{}) (string, *errs.ParserError) {
	convertedMessage := string(`{
		"definitions" : {
		  "%s": %s
		},
		"$ref" : "#/definitions/%s"
	}`)

	refKey := fmt.Sprintf("%s:%s", message["type"].(string), message["name"].(string))
	fAvro := NewFixedAvro(message["type"].(string), refKey, message["size"].(float64))

	jFAvro, _ := json.Marshal(fAvro)
	translated := fmt.Sprintf("%s", string(jFAvro))
	return fmt.Sprintf(convertedMessage, refKey, translated, refKey), nil
}
