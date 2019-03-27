package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	parser "github.com/asyncapi/parser/pkg"
	"github.com/pkg/errors"
)

var (
	filename = flag.String("file", parser.DefaultDefinitionFile, fmt.Sprintf("-file %s", parser.DefaultDefinitionFile))
)

func init() {
	flag.Parse()
}

func main() {
	fBytes, err := ioutil.ReadFile(*filename)
	p, err := parser.Parse(fBytes)
	handleError(err, "error parsing definition file")
	fmt.Println("Definition %s", p)
}

func handleError(err error, msg string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, errors.Wrap(err, msg))
		os.Exit(2)
	}
}
