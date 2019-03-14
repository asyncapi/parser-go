package hlsp

import (
	"encoding/json"

	"github.com/asyncapi/parser/models"
)

//ParserMessage maps AsyncAPI "x-parser-message" object
type ParserMessage struct {
	Message       *models.OperationMessage `json:"message,omitempty"`
	ChannelName   string                   `json:"channelName,omitempty"`
	OperationName string                   `json:"operationName,omitempty"`
	OperationId   string                   `json:"operationId,omitempty"`
}

type ParserMessages []ParserMessage

// Beautify Create a list of messages on the root level of the document, using field name x-parser-messages
func Beautify(jsonDocument []byte) (json.RawMessage, error) {
	asyncAPIObj := models.AsyncapiDocument{}
	var messages ParserMessages

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

func getOperationData(channelItem *models.ChannelItem) ParserMessage {
	if channelItem.Publish != nil {
		return ParserMessage{
			OperationName: "publish",
			OperationId:   channelItem.Publish.OperationId,
			Message:       channelItem.Publish.Message,
		}
	}
	return ParserMessage{
		OperationName: "subscribe",
		OperationId:   channelItem.Subscribe.OperationId,
		Message:       channelItem.Subscribe.Message,
	}
}
