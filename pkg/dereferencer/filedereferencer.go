package dereferencer

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/pkg/errors"
	"github.com/stretchr/objx"
)

type fileDereferencer struct {
	f                  dereferencer
	ref                []byte
	yamlOrJSONDocument []byte
}

func (fdef *fileDereferencer) Dereference(ref string, document []byte) (value []byte, err error) {
	m, err := objx.FromJSON(string(document))
	if err != nil {
		return nil, err
	}
	path := strings.Split(strings.Trim(ref, "#"), "/")
	element := m.Get(trimFirstRune(strings.Join(path, ".")))
	if element == nil || (*element).Data() == nil {
		return nil, errors.Errorf("element %q is not found", ref)
	}

	return json.Marshal((*element).Data())
}

func checkFile(filename string) (fileData []byte, ref string, err error) {
	paths := strings.Split(filename, "#")
	fileData, err = ioutil.ReadFile(paths[0])
	if err != nil {
		return nil, "", err
	}
	return fileData, paths[1], nil
}
