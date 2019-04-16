package dereferencer

import (
	// "log"
	"encoding/json"
	"unicode/utf8"
	"strings"
	"github.com/stretchr/objx"
)

type fileDereferencer struct {
	f dereferencer
	ref []byte
	yamlOrJSONDocument []byte
}

func trimFirstRune(s string) string {
    _, i := utf8.DecodeRuneInString(s)
    return s[i:]
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