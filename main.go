package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	parser "github.com/asyncapi/parser/pkg"
	"github.com/pkg/errors"
)

var (
	defaultDefinitionFile = "asyncapi/2.0.0/example.yaml"
	filename = flag.String("file", defaultDefinitionFile, fmt.Sprintf("-file %s", defaultDefinitionFile))
)

func init() {
	flag.Parse()
}

func main() {
	fBytes, err := ioutil.ReadFile(*filename)
	handleError(err, "Reading file")
	p, perr := parser.Parse(fBytes)
	if perr != nil {
		fmt.Fprintln(os.Stderr, errors.Wrap(perr, perr.Error()))
		os.Exit(2)
	}
	jOut, err:= json.MarshalIndent(p, "","  ")
	fmt.Println("Definition %s", string(jOut))
}

func handleError(err error, msg string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, errors.Wrap(err, msg))
		os.Exit(2)
	}
}
