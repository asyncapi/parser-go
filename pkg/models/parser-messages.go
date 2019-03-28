package models

//ParserMessage maps AsyncAPI "x-parser-messages" array item
type ParserMessage struct {
	Message       *OperationMessage `json:"message,omitempty"`
	ChannelName   string            `json:"channelName,omitempty"`
	OperationName string            `json:"operationName,omitempty"`
	OperationId   string            `json:"operationId,omitempty"`
}

//ParserMessages maps AsyncAPI "x-parser-messages" array
type ParserMessages []ParserMessage
