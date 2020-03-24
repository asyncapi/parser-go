# AsyncAPI Parser

[![Build Status](https://godoc.org/github.com/asyncapi/parser?status.svg)](https://godoc.org/github.com/asyncapi/parser) [![Build Status](https://travis-ci.org/asyncapi/parser.svg?branch=master)](https://travis-ci.org/asyncapi/parser) [![Go Report Card](https://goreportcard.com/badge/github.com/asyncapi/parser)](https://goreportcard.com/report/github.com/asyncapi/parser)

## Overview

The AsyncAPI Parser validates AsyncAPI documents according to dedicated schemas. The supported schemas are:

- AsyncAPI schema
- OpenAPI schema
- JSON schema
  If you use the parser as a package, you can also register external schemas. For example, you can write your own schema.

Documents provided for the AsyncAPI Parser can be in the `.yaml` or `.json` formats. If a document is invalid, the parser provides a message listing all errors. If a document is valid, the parser provides dereferenced output. During the dereference process the AsyncAPI parser substitutes a reference with a full definition. The dereferenced output is always in the `.json` format.

## Prerequisites

- [Golang](https://golang.org/dl/) version 1.12+

## Installation

To install the AsyncAPI Parser package, run:

```bash
go get github.com/asyncapi/parser-go/...
```

## Usage

You can use the AsyncAPI Parser in two ways:

- Before you use the AsyncAPI Parser in the terminal, build the application. Run:

  ```bash
  git clone https://github.com/asyncapi/parser.git
  cd ./parser
  go build -o=asyncapi-parser ./cmd/api-parser/main.go
  ```

  To use the AsyncAPI Parser run the following command:

  ```text
  asyncapi-parser <document_path>
  ```

  where `document_path` is a mandatory argument that is either a URL or a file path to an AsyncAPI document.

- You can also use the AsyncAPI Parser without building the application, using Golang. Run:
  `bash go run ./cmd/api-parser/main.go <document_path>`
  where `document_path` is a mandatory argument that is either a URL or a file path to an AsyncAPI document.
  **Examples**
  See the following examples of the AsyncAPI Parser usage in the terminal:
- Validation of the `gitter-streaming.yaml` valid file:

  ```text
  asyncapi-parser https://raw.githubusercontent.com/asyncapi/asyncapi/master/examples/2.0.0/gitter-streaming.yml
  ```

- Validation of the `oneof.yml` invalid file:

  ```text
  go run ./cmd/api-parser/main.go https://raw.githubusercontent.com/asyncapi/asyncapi/master/examples/1.1.0/oneof.yml
  ```

  Output:

  ```text
  (root): id is required
  (root): channels is required
  (root): Additional property topics is not allowed
  asyncapi: asyncapi must be one of the following: "2.0.0"
  ```

## Contribution

If you have a feature request, add it as an issue or propose changes in a pull request (PR).
If you create a feature request, use the dedicated **Feature request** issue template. When you create a PR, follow the contributing rules described in the [`CONTRIBUTING.md`](CONTRIBUTING.md) document.

## Roadmap

- `avro` schema support
- extensions support
- json-schema `$id` property support

## Credits

- Fran Mendez
- Raisel Melian
- Ruben Hervas
- Marcin Witalis from [Kyma](https://kyma-project.io/)
