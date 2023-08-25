package parser

import (
	"bytes"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const oneOfJSONFile = "./testdata/oneof.json"

var (
	noopMessageProcessor MessageProcessor = func(_ map[string]interface{}) error {
		return nil
	}
)

func TestMessageProcessor_BuildParser(t *testing.T) {
	parse := noopMessageProcessor.BuildParser()
	writer := bytes.NewBufferString("")
	reader, err := os.Open(oneOfJSONFile)
	assert.NoError(t, err)
	err = parse(reader, writer)
	assert.NoError(t, err)
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
			r, err := NewReader(test.doc)
			assert.NoError(t, err)
			assert.NotNil(t, r)
		})
	}
}
