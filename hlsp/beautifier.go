package hlsp

import (
	"encoding/json"
	"fmt"

	"github.com/asyncapi/parser/models"
)

// Beautify Create a list of messages on the root level of the document, using field name x-parser-messages
func Beautify(jsonDocument json.RawMessage) (json.RawMessage, error) {

	asyncAPIObj := models.AsyncapiDocument{}

	if err := json.Unmarshal(jsonDocument, &asyncAPIObj); err != nil {
		return nil, err
	}

	fmt.Println(asyncAPIObj.Info.Title)

	return nil, nil
}
