package hlsp

import (
	"encoding/json"

	"github.com/asyncapi/parser/models"
)

// Beautify Create a list of messages on the root level of the document, using field name x-parser-messages
func Beautify(jsonDocument []byte) (json.RawMessage, error) {
	asyncAPIObj := models.AsyncapiDocument{}
	var messages models.ParserMessages

	if err := json.Unmarshal(jsonDocument, &asyncAPIObj); err != nil {
		return nil, err
	}

	for key, value := range asyncAPIObj.Channels {
		pm := getOperationData(value)
		pm.ChannelName = key
		messages = append(messages, pm)
	}

	messagesBytes, err := json.Marshal(messages)

	result, err := models.MergeExtensions(jsonDocument, map[string]json.RawMessage{
		"x-parser-messages": messagesBytes,
	})

	return result, err
}

func getOperationData(channelItem *models.ChannelItem) models.ParserMessage {
	if channelItem.Publish != nil {
		return models.ParserMessage{
			OperationName: "publish",
			OperationId:   channelItem.Publish.OperationId,
			Message:       channelItem.Publish.Message,
		}
	}
	return models.ParserMessage{
		OperationName: "subscribe",
		OperationId:   channelItem.Subscribe.OperationId,
		Message:       channelItem.Subscribe.Message,
	}
}
