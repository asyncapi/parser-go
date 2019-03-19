# Technical proposal
The way I see us succeeding is by creating a pluggable architecture and splitting the work into the following components or plugins:

* High-level specification parser/validator
* Schema parser
* ProtocolInfo parser
* Specification extension parser

## Terminology

### High-level specification parser/validator
This component is in charge of parsing/validating the standard specification. It does not go into details on how to parse schemas. It's sole purpose is to guarantee the AsyncAPI document has the correct structure and provide a beautified version of it.

### Schema parser
This component must parse the schemas conforming to their type, e.g., JSON Schema Draft 04-07, Protobuf, Avro, OpenAPI schema, etc.

### ProtocolInfo parser
A protocolInfo object is a special type of specification extension. In version 2.0.0, protocol information will be contained in their own specification extension. For instance:

```yaml
...
publish:
  protocolInfo:
    mqtt:
      headers:
      qos:
        type: number
        enum: [1,2]
        description: Quality of Service. 0 is forbidden.
      anyProtocolSpecificConfiguration: its value
```

Following the example above, an MQTT SDK could use the `protocolInfo.mqtt` object to serialize/deserialize a message to/from an MQTT binary message format.

### Specification extension parser
An AsyncAPI document may contain custom information provided in the form of [specification extensions](https://github.com/asyncapi/asyncapi/blob/master/versions/1.2.0/asyncapi.md#specificationExtensions). Therefore, a specification extension parser is the piece of software that understands, interprets, and validates the extension.

###### Examples:

```yaml
info:
  x-twitter: '@awesome-user'
```

Or a more complex one:

```yaml
components:
  schemas:
    chatMessage:
      type: object
      properties:
        channel:
          type: string
          x-format: slack-channel # This adds an unsupported vendor-specific format. Notice this field may be parsed by other specification extensions. E.g., a Twitter specification extension may add support for x-format: twitter-handle.
```

## Requirements

### Parser/Validator

![](./assets/parser-diagram.png)

The flow is as follows:

1. The High-level specification parser (HLSP) receives either a YAML or JSON AsyncAPI document. It parses the document and checks if it's valid AsyncAPI. Skips specification extensions and schemas validation. If validation fails, the Parser/Validator should trigger an error. Produces a beautified version of the document in JSON Schema Draft 07.
2. The output of the HLSP serves as the input of the Schema parser, which will in turn identify the type of the schema and will pass it to the appropiate parser. It's the responsiblity of each parser to translate the schema format to valid JSON Schema Draft 07 and return it back to the Schema parser. E.g., Protobuf to JSON Schema Draft 07, Avro to JSON Schema Draft 07, JSON Schema Draft 04 to JSON Schema Draft 07, etc. The Schema parser will replace the original schema definition with the one returned from a specific format parser.
3. The output of the Schema parser serves as the input for the Specification extensions parser (SEP), which will seek for extensions and will pass them to the appropiate parser. It's the responsiblity of each parser to interpret, validate, and return the validation results to the SEP. It validation fails, the Parser/Validator should trigger an error.
4. To finish, we must check the whole document is still valid JSON Schema Draft 07. If it is, the output of the parser must be the JSON Schema document. Otherwise, it should trigger an error.

|Input|Output|Required by|
|-----|------|-----------|
|An AsyncAPI document in YAML or JSON format.| A beautified version of the document in JSON Schema format Draft 07.| Code generators, SDKs, Documentation generators, and potentially every future tool.|

###### Example code

> This code is just for explanation purposes. Please, **do not** take it as an example of how it should be implemented.

```go
import (
    "fmt"

    "github.com/asyncapi/hlsp"
    "github.com/asyncapi/schemaparser"
    "github.com/asyncapi/protojsonschema"
    "github.com/asyncapi/sep"
    "github.com/someone/asyncapiTwitterExtension"
    "github.com/xeipuuv/gojsonschema"
)

func main() {
    // Step 1
    document := `{"asyncapi": "2.0.0", ...}`
    result, err := hslp.Parse(document)
    if err != nil {
        panic(err.Error())
    }

    // Step 2
    result, err := schemaparser.Parse(result, [protojsonschema.Parser])
    if err != nil {
        panic(err.Error())
    }

    // Step 3
    result, err := sep.Parse(result, [asyncapiTwitterExtension])
    if err != nil {
        panic(err.Error())
    }

    // Step 4
    schemaLoader := gojsonschema.NewReferenceLoader("file:///path/to/asyncapi/2.0.0/schema.json")
    documentLoader := gojsonschema.NewStringLoader(document)
    result, err := gojsonschema.Validate(schemaLoader, documentLoader)
    if err != nil {
        panic(err.Error())
    }
}
```

### High-level specification parser (HLSP)

The role of HLSP is to:
1. Dereference all the `$ref`s, except the those in the schemas.
2. Apply and resolve all the `traits` in a document.
3. Validate the document, except messages and extensions.
4. Beautify the AsyncAPI document, by adding some handy information.

|Input|Output|Required by|
|-----|------|-----------|
|An AsyncAPI document in YAML or JSON format.| A beautified version of the document in JSON Schema Draft 07 format.| Parser/Validator

### Schema parser

This component is in charge of understanding all the messages in a document and passing them to the appropiate schema parser. As an example, if a message's `schemaFormat` is `protobuf`, this component should pass the message payload to the Protobuf Schema Parser, which in turn will return the schema converted to JSON Schema Draft 07.

Once the schema is parsed and converted, this component will update the messages in the document with the result.

Since we can't anticipate how many schema formats we'll have, this information must be provided by the user when calling the parser.

|Input|Output|Required by|
|-----|------|-----------|
|An AsyncAPI **schema** in YAML or JSON format.| A beautified version of the **schema** in JSON Schema Draft 07 format.| Parser/Validator

### Specification extension parser

This component is in charge of parsing the specification extensions and `protocolInfo` objects. The flow is pretty much the same as the one with schema parsers.

Since we can't anticipate the definition of the extensions, this information must be provided by the user when calling the parser.

|Input|Output|Required by|
|-----|------|-----------|
|An AsyncAPI specification extension in YAML or JSON format.| A resolved version of the specification extension in JSON Schema Draft 07 format.| Parser/Validator

## FAQ

### Why Go?

We chose Go because of three reasons:

1. It compiles to C shared objects, so it means we can reuse the work done here in another languages very easily, avoiding the cost of maintaining many versions in different programming languages.
2. Go performance is really good.
3. The Go community is already big and keeps growing, which makes it a safe bet. That was the key reason not to choose Rust, which is much better than Go regarding points 1 and 2.
