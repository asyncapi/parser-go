package jsonpath

import (
	"github.com/asyncapi/parser-go/pkg/decode"

	"net/http"
	"os"
	"strings"
)

type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type RefLoader func(string) (map[string]interface{}, error)

func buildHttpLoader(client HttpClient) func(string) (map[string]interface{}, error) {
	return func(url string) (map[string]interface{}, error) {
		resp, err := client.Get(url)
		if err != nil {
			return nil, err
		}
		return decode.ToMap(resp.Body)
	}
}

func (l RefLoader) fileLoader(path string) (map[string]interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return decode.ToMap(file)
}

func NewRefLoader(client HttpClient) RefLoader {
	return buildHttpLoader(client)
}

func (l RefLoader) Load(documentRef string) (map[string]interface{}, error) {
	switch {
	case strings.HasPrefix(documentRef, "http://") || strings.HasPrefix(documentRef, "https://"):
		return l(documentRef)
	default:
		return l.fileLoader(documentRef)
	}
}
