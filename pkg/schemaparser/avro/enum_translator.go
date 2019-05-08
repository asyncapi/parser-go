package avro

import (
	"encoding/json"
	"fmt"

	"github.com/asyncapi/parser/pkg/errs"
)

// EnumAvro maps enum avro scheme
type EnumAvro struct {
	Enum []string `json:"enum"`
}

// Convert transforms avro formatted message to JSONSchema
func (ra *EnumAvro) Convert(message map[string]interface{}) (string, *errs.ParserError) {
	convertedMessage := string(`{
		"definitions" : {
		  "%s": %s
		},
		"$ref" : "#/definitions/%s"
	}`)
	var refKey, translated string
	refKey = fmt.Sprintf("%s:%s", message["type"].(string), message["name"].(string))
	var enumArray []string
	objectArray := message["symbols"].([]interface{})
	for _, o := range objectArray {
		enumArray = append(enumArray, o.(string))
	}
	eAvro := EnumAvro{Enum: enumArray}
	jEAvro, _ := json.Marshal(eAvro)
	translated = fmt.Sprintf("%s", string(jEAvro))
	return fmt.Sprintf(convertedMessage, refKey, translated, refKey), nil
}
