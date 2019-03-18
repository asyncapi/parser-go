package models

import (
	"encoding/json"
	"errors"
)

// ParsedAsyncAPI represents an AsyncAPI document after being parsed.
type ParsedAsyncAPI struct {
	AsyncapiDocument
	parserMessages   ParserMessages
	producedMessages ParserMessages
	consumedMessages ParserMessages
}

// ListMessages retrieves a conveniently formatted list of messages and their channels.
func (doc *ParsedAsyncAPI) ListMessages() (ParserMessages, error) {
	// If the list of messages has been unmarshaled already, don't do it again.
	if doc.parserMessages != nil {
		return doc.parserMessages, nil
	}

	err := json.Unmarshal(doc.Extensions["x-parser-messages"], &doc.parserMessages)
	if err != nil {
		return nil, err
	}
	return doc.parserMessages, nil
}

// ListProducedMessages retrieves a convenient list of messages this AsyncAPI produces/publishes.
func (doc *ParsedAsyncAPI) ListProducedMessages() (ParserMessages, error) {
	// If the list of messages has been unmarshaled already, don't do it again.
	if doc.producedMessages != nil {
		return doc.producedMessages, nil
	}

	msgs, err := doc.ListMessages()
	if err != nil {
		return nil, err
	}

	doc.producedMessages = filter(msgs, func(msg ParserMessage) bool {
		if msg.OperationName == "publish" {
			return true
		}
		return false
	})

	return doc.producedMessages, nil
}

// ListConsumedMessages retrieves a convenient list of messages this AsyncAPI consumes/subscribes to.
func (doc *ParsedAsyncAPI) ListConsumedMessages() (ParserMessages, error) {
	// If the list of messages has been unmarshaled already, don't do it again.
	if doc.consumedMessages != nil {
		return doc.consumedMessages, nil
	}

	msgs, err := doc.ListMessages()
	if err != nil {
		return nil, err
	}

	doc.consumedMessages = filter(msgs, func(msg ParserMessage) bool {
		if msg.OperationName == "subscribe" {
			return true
		}
		return false
	})

	return doc.consumedMessages, nil
}

// ApplyOverlay applies an AsyncAPI overlay and returns the resulting document.
func (doc *ParsedAsyncAPI) ApplyOverlay(yamlOrJSONDocument []byte) (*ParsedAsyncAPI, error) {
	return nil, errors.New("not yet implemented")
}

func filter(vs []ParserMessage, f func(ParserMessage) bool) ParserMessages {
	vsf := make([]ParserMessage, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
