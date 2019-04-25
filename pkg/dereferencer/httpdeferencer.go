package dereferencer

import (
	"encoding/json"
	"github.com/stretchr/objx"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	"net/http"
	"strings"
)

type httpDereferencer struct {
	f                  dereferencer
	ref                []byte
	yamlOrJSONDocument []byte
}

func (htpp *httpDereferencer) Dereference(ref string, document []byte) (value []byte, err error) {
	m, err := objx.FromJSON(string(document))
	if err != nil {
		return nil, err
	}
	path := strings.Split(strings.Trim(ref, "#"), "/")
	element := m.Get(trimFirstRune(strings.Join(path, ".")))
	// log.Printf("Obj %v \n", (*element).Data())
	belement, err := json.Marshal((*element).Data())
	if element != nil {
		value = belement
	}
	return value, nil
}

func resolveURL(url string) (URLData []byte, ref string, err error) {
	paths := strings.Split(url, "#")
	resp, err := http.Get(paths[0])
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	URLData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}
	schemaLoader := gojsonschema.NewBytesLoader(URLData)
	documentLoader := gojsonschema.NewBytesLoader(URLData)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)

	if result.Valid() {
		return URLData, paths[1], nil
	}

	return nil, paths[1], err
}
