package v2

import (
	parserErrors "github.com/asyncapi/parser-go/pkg/error"

	"github.com/pkg/errors"

	"fmt"
)

var ErrInvalidValue = errors.New("invalid value")

func schemaFormat(m map[string]interface{}) string {
	schemaFormat, found := m["schemaFormat"]
	if !found {
		return ""
	}
	return fmt.Sprintf("%v", schemaFormat)
}

type Dispatcher map[string]func(interface{}) error

func (d Dispatcher) do(messages []map[string]interface{}) error {
	var errs []error
	for _, msg := range messages {
		schemaFormat := schemaFormat(msg)
		pm, found := d[schemaFormat]
		if !found {
			continue
		}
		err := pm(msg)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return parserErrors.New(errs...)
}

func (d *Dispatcher) Add(pm func(interface{}) error, labels ...string) error {
	for _, key := range labels {
		(*d)[key] = pm
	}
	return nil
}

func BuildMessageProcessor(dispatcher Dispatcher) func(map[string]interface{}) error {
	return func(doc map[string]interface{}) error {
		var errs []error
		channels, found := (doc)["channels"].(map[string]interface{})
		if !found {
			return errors.Wrap(ErrInvalidValue, "channels")
		}
		for _, channel := range channels {
			chanMessages, err := extractMessages(channel)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			if err := dispatcher.do(chanMessages); err != nil {
				errs = append(errs, err)
			}
		}
		return parserErrors.New(errs...)
	}
}

func extractMessages(channel interface{}) ([]map[string]interface{}, error) {
	var (
		messages []map[string]interface{}
		errs     []error
	)
	channelMap, ok := channel.(map[string]interface{})
	if !ok {
		return nil, errors.Wrap(ErrInvalidValue, "channel")
	}
	for _, key := range []string{"publish", "subscribe"} {
		channel, ok := channelMap[key].(map[string]interface{})
		if ok {
			pubMsg, err := extractMessage(channel["message"])
			switch err != nil {
			case true:
				errs = append(errs, err)
			default:
				messages = append(messages, pubMsg...)
			}
		}
	}
	return messages, parserErrors.New(errs...)
}

func extractMessage(message interface{}) ([]map[string]interface{}, error) {
	msg, ok := message.(map[string]interface{})
	if !ok {
		return nil, errors.Wrap(ErrInvalidValue, "message")
	}
	oneOf, ok := msg["oneOf"]
	if !ok {
		return []map[string]interface{}{
			msg,
		}, nil
	}
	oneOfList, ok := oneOf.([]interface{})
	if !ok {
		return nil, errors.Wrap(ErrInvalidValue, "oneOf")
	}
	var result []map[string]interface{}
	var errs []error
	for _, msg := range oneOfList {
		msgMap, ok := msg.(map[string]interface{})
		if !ok {
			errs = append(errs, errors.Wrap(ErrInvalidValue, "message"))
			continue
		}
		result = append(result, msgMap)
	}
	return result, parserErrors.New(errs...)
}
