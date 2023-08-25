package jsonpath

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func data() io.ReadCloser {
	data := map[string]interface{}{
		"test": "me",
	}
	result, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return ioutil.NopCloser(bytes.NewReader(result))
}

type ClientMock struct{}

func (c *ClientMock) Get(url string) (resp *http.Response, err error) {
	switch url {
	case "http://asyncapi.com":
		return &http.Response{
			Status: "200",
			Body:   data(),
		}, nil
	default:
		return nil, http.ErrServerClosed
	}
}

func TestLoader(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"http://asyncapi.com"},
		{"./testdata/sample.json"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			loader := NewRefLoader(&ClientMock{})
			actual, err := loader.Load(test.name)
			assert.NoError(t, err)
			expected := map[string]interface{}{
				"test": "me",
			}
			assert.Equal(t, expected, actual)
		})
	}
}

func TestLoaderErr(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"http://a"},
		{"./test"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			loader := NewRefLoader(&ClientMock{})
			_, err := loader.Load(test.name)
			assert.Error(t, err)
		})
	}
}
