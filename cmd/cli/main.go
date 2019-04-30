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
	filename              = flag.String("file", defaultDefinitionFile, fmt.Sprintf("-file %s", defaultDefinitionFile))
	circularReferences    = flag.Bool("c", true, "-c=[true,false]")
)

func init() {
	flag.Parse()
}

func main() {
	fBytes, err := ioutil.ReadFile(*filename)
	handleError(err, "Reading file")
	p, perr := parser.Parse(fBytes, *circularReferences)
	if perr != nil {
		fmt.Fprintln(os.Stderr, errors.Wrap(perr, perr.Error()))
		fmt.Fprintln(os.Stderr, perr.ParsingErrors)
		os.Exit(2)
	}
	jOut, err := json.MarshalIndent(p, "", "  ")
	fmt.Printf("Definition %s", string(jOut))
}

func handleError(err error, msg string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, errors.Wrap(err, msg))
		os.Exit(2)
	}
}
