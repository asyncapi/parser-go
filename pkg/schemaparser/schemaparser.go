package schemaparser

import (
	"encoding/json"

	"github.com/asyncapi/parser/pkg/errs"
	"github.com/asyncapi/parser/pkg/models"
)

type SchemaParser interface {
	Parse(json.RawMessage) *errs.ParserError
}

// Parse parses all the message schemas in an AsyncAPI document.
// Note: It doesn't parse the messages under components that are not used.
func Parse(asyncapiJSON json.RawMessage) *errs.ParserError {
	var doc models.AsyncapiDocument
	if err := json.Unmarshal(asyncapiJSON, &doc); err != nil {
		return &errs.ParserError{
			ErrorMessage: err.Error(),
		}
	}

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

	j, err := message.Payload.MarshalJSON()
	if err != nil {
		return &errs.ParserError{
			ErrorMessage: err.Error(),
		}
	}

	var schemaParser SchemaParser

	switch message.SchemaFormat {
	case "":
	case "openapi":
	case "asyncapi":
	case "application/vnd.oai.openapi":
	case "application/vnd.asyncapi":
		schemaParser = OpenAPI{}
	}

	if schemaParser == nil {
		return nil
	}

	return schemaParser.Parse(j)
}
