package v2

import (
	"github.com/asyncapi/parser-go/pkg/decode"
	"github.com/asyncapi/parser-go/pkg/jsonpath"
	"github.com/stretchr/testify/assert"

	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name             string
		blackListedPaths []string
		doc              map[string]interface{}
		expectedDoc      map[string]interface{}
		expectedErr      bool
	}{
		{
			name:             "anyOf",
			blackListedPaths: []string{"#/components/schemas"},
			doc:              testDataFromFile("input", "anyof.json"),
			expectedDoc:      testDataFromFile("output", "anyof.json"),
		},
		{
			name:             "loop1",
			blackListedPaths: []string{"#/components/schemas"},
			doc:              testDataFromFile("input", "loop1.json"),
			expectedErr:      true,
		},
		{
			name:             "loop2",
			blackListedPaths: []string{"#/components/schemas"},
			doc:              testDataFromFile("input", "loop2.json"),
			expectedErr:      true,
		},
		{
			name:             "invalid-ref",
			blackListedPaths: []string{"#/components/schemas"},
			doc:              testDataFromFile("input", "invalid-ref.json"),
			expectedErr:      true,
		},
		{
			name:             "recursive-references",
			blackListedPaths: []string{"#/components/schemas"},
			doc:              testDataFromFile("input", "recref0.json"),
			expectedDoc:      testDataFromFile("output", "recref.json"),
			expectedErr:      false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			refLoader := jsonpath.NewRefLoader(http.DefaultClient)
			p := NewParser(refLoader, test.blackListedPaths...)
			err := p.Parse(test.doc)
			if test.expectedErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.expectedDoc, test.doc)
		})
	}
}

func testDataFromFile(dirName, fileName string) map[string]interface{} {
	file, err := os.Open(fmt.Sprintf("./testdata/%s/%s", dirName, fileName))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileMap, err := decode.ToMap(file)
	if err != nil {
		panic(err)
	}
	return fileMap
}
