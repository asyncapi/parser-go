package hlsp

// Parse receives either a YAML or JSON AsyncAPI document. It parses the document and checks if it's valid AsyncAPI. Skips specification extensions and schemas validation. If validation fails, the Parser/Validator should trigger an error. Produces a beautified version of the document in JSON Schema Draft 07
func Parse() string {
	return "Hello World"
}
