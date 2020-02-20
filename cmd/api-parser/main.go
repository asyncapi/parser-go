package main

import (
	"fmt"
	"log"
	"os"

	"github.com/asyncapi/parser-go/internal/cli"

	"github.com/docopt/docopt-go"
)

const version = "asyncapi-parser 0.2"

func main() {
	usage := fmt.Sprintf(`
	Parse AsyncAPI %s version documents.
	
  Usage:
  asyncapi-parser <PATH>
  asyncapi-parser -h | --help | --version
		
  Arguments:
	PATH  a path to asyncapi document (either url or local file, supports json and yaml format)`,
		"2.0.0")

	opts, err := docopt.ParseArgs(usage, nil, version)
	if err != nil {
		log.Fatal(err)
	}
	asyncapiCli := cli.New(opts)
	parser, reader, err := asyncapiCli.NewParserAndReader()
	if err != nil {
		log.Fatal(err)
	}
	err = parser(reader, os.Stdout)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
