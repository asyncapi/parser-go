package parser

import (
	. "github.com/onsi/gomega"

	"bytes"
	"os"
	"testing"
)

var (
	noopMessageProcessor MessageProcessor = func(_ *map[string]interface{}) error {
		return nil
	}
)

func TestMessageProcessor_BuildParse(t *testing.T) {
	g := NewWithT(t)
	parse := noopMessageProcessor.BuildParse()
	writer := bytes.NewBufferString("")
	reader, err := os.Open("./testdata/oneof.json")
	g.Expect(err).ShouldNot(HaveOccurred())
	err = parse(reader, writer)
	g.Expect(err).ShouldNot(HaveOccurred())
}
