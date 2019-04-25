package dereferencer

import (
	"encoding/json"
	"github.com/stretchr/objx"
	"io/ioutil"
	"strings"
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
	// log.Printf("Obj %v \n", (*element).Data())
	belement, err := json.Marshal((*element).Data())
	if element != nil {
		value = belement
	}
	return value, nil
}

func checkFile(filename string) (fileData []byte, ref string, err error) {
	paths := strings.Split(filename, "#")
	fileData, err = ioutil.ReadFile(paths[0])
	if err != nil {
		return nil, "", err
	}
	return fileData, paths[1], nil
}
