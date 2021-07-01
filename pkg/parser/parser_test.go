package parser

import (
	"bytes"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"
)

const oneOfJSONFile = "./testdata/oneof.json"

var (
	noopMessageProcessor MessageProcessor = func(_ *map[string]interface{}) error {
		return nil
	}
)

func TestMessageProcessor_BuildParser(t *testing.T) {
	g := NewWithT(t)
	parse := noopMessageProcessor.BuildParser()
	writer := bytes.NewBufferString("")
	reader, err := os.Open(oneOfJSONFile)
	g.Expect(err).ShouldNot(HaveOccurred())
	err = parse(reader, writer)
	g.Expect(err).ShouldNot(HaveOccurred())
}

func TestNewReader(t *testing.T) {
	dir, filename := filepath.Split(oneOfJSONFile)
	http.Handle("/", http.FileServer(http.Dir(dir)))
	go func() {
		_ = http.ListenAndServe(":3001", nil)
	}()

	tests := []struct {
		name string
		doc  string
	}{
		{
			name: "Local file",
			doc:  oneOfJSONFile,
		},
		{
			name: "File served via HTTP",
			doc:  "localhost:3001/" + filename,
		},
		{
			name: "Plain text",
			doc:  "whatever",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := NewWithT(t)
			r, err := NewReader(test.doc)
			g.Expect(err).ShouldNot(HaveOccurred())
			g.Expect(r).ShouldNot(BeNil())
		})
	}
}

func TestNew(t *testing.T) {
	g := NewWithT(t)
	p, err := New()
	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(p).ShouldNot(BeNil())
}
