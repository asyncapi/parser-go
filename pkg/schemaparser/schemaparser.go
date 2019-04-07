package schemaparser

import (
	"encoding/json"

	"github.com/asyncapi/parser/pkg/errs"
	"github.com/asyncapi/parser/pkg/models"
	"github.com/asyncapi/parser/pkg/schemaparser/openapi"
)

type SchemaParser interface {
	Parse(json.RawMessage) *errs.ParserError
}

// Parse parses all the message schemas in an AsyncAPI document.
// Note: It doesn't parse the messages under components that are not used.
func Parse(doc *models.AsyncapiDocument) *errs.ParserError {
	for _, channel := range doc.Channels {
		if channel.Publish != nil {
			if err := parseMessages(channel.Publish.Message); err != nil {
				return err
			}
		}
		if channel.Subscribe != nil {
			if err := parseMessages(channel.Subscribe.Message); err != nil {
				return err
			}
		}
	}

	return nil
}

func parseMessages(message *models.Message) *errs.ParserError {
	if message.OneOf != nil {
		for _, m := range message.OneOf {
			if err := parseMessage(m); err != nil {
				return err
			}
		}
	} else {
		if err := parseMessage(message); err != nil {
			return err
		}
	}

	return nil
}

func parseMessage(message *models.Message) *errs.ParserError {
	if message.Payload == nil {
		return nil
	}

	switch message.SchemaFormat {
	case "", "openapi", "asyncapi", "application/vnd.oai.openapi", "application/vnd.asyncapi":
		return openapi.OpenAPI{}.Parse(message)
	}

	return nil
}
