package dereferencer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/stretchr/objx"
)

type httpDereferencer struct {
	f                  dereferencer
	ref                []byte
	yamlOrJSONDocument []byte
}

func (htpp *httpDereferencer) Dereference(ref string, document []byte) (value []byte, err error) {
	m, err := objx.FromJSON(string(document))
	if err != nil {
		return nil, errors.Errorf("failed to create object from json %q: %v", string(document), err)
	}
	path := strings.Split(strings.Trim(ref, "#"), "/")

	key := strings.Join(path, ".")

	element := m.Get(trimFirstRune(key))
	if element == nil || (*element).Data() == nil {
		return nil, errors.Errorf("element %q is not found", ref)
	}

	return json.Marshal((*element).Data())
}

func resolveURL(url string) (URLData []byte, ref string, err error) {
	paths := strings.Split(url, "#")
	if len(paths) >= 2 {
		ref = paths[1]
	}

	resp, err := http.Get(paths[0])
	if err != nil {
		return nil, ref, err
	}
	defer resp.Body.Close()

	URLData, err = ioutil.ReadAll(resp.Body)

	return URLData, ref, err
}
