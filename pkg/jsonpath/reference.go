package jsonpath

import (
	"github.com/pkg/errors"

	"fmt"
	"strings"
)

var (
	ErrInvalidPath      = errors.New("invalid path")
	ErrInvalidKey       = errors.New("key not found")
	ErrInvalidEncoding  = errors.New("invalid encoding")
	ErrInvalidReference = errors.New("invalid reference")
	decodeMap           = map[uint8]rune{
		'0': '~',
		'1': '/',
	}
	root = ""
)

type Ref struct {
	pointer string
	uri     string
	path    []string
}

func NewRef(ref interface{}) (Ref, error) {
	strRef := fmt.Sprintf("%v", ref)
	uri, pointer, err := ParseRefStr(root, strRef)
	if err != nil {
		return Ref{}, err
	}
	return Ref{
		pointer: pointer,
		uri:     uri,
		path:    strings.Split(pointer, "/")[1:],
	}, nil
}

func (r Ref) String() string {
	return fmt.Sprintf(`%s#%s`, r.uri, r.pointer)
}

func (r Ref) URI() string {
	return r.uri
}

func (r Ref) Path() []string {
	return r.path
}

func (r Ref) NewChild(name interface{}) (Ref, error) {
	childEncodedKey := EncodeEntryKey(fmt.Sprintf("%v", name))
	return NewRef(fmt.Sprintf("%s/%s", r.String(), childEncodedKey))
}

func DecodeEntryKey(name string) (string, error) {
	if name == "~" || name == "/" {
		return "", errors.Wrap(ErrInvalidEncoding, name)
	}
	nameLen := len(name)
	if nameLen < 2 {
		return name, nil
	}
	builder := strings.Builder{}
	for i := 0; i < nameLen; i++ {
		if name[i] != '~' {
			builder.WriteRune(rune(name[i]))
			continue
		}
		switch name[i+1] {
		case '0', '1':
			builder.WriteRune(decodeMap[name[i+1]])
			i++
			continue
		default:
			return "", errors.Wrap(ErrInvalidEncoding, name)
		}
	}
	return builder.String(), nil
}

func EncodeEntryKey(name string) string {
	builder := strings.Builder{}
	for _, r := range name {
		switch r {
		case '~':
			builder.WriteString("~0")
		case '/':
			builder.WriteString("~1")
		default:
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

func getValue(key string, v interface{}) (interface{}, error) {
	m, ok := v.(map[string]interface{})
	if !ok {
		return nil, ErrInvalidKey
	}
	result, exist := m[key]
	if !exist {
		return nil, ErrInvalidKey
	}
	return result, nil
}

func GetRefObject(path []string, v interface{}) (map[string]interface{}, error) {
	if len(path) < 1 {
		return nil, ErrInvalidPath
	}
	var (
		current = v
		err     error
	)
	for _, key := range path {
		current, err = getValue(key, current)
		if err != nil {
			return nil, err
		}
	}
	result, ok := current.(map[string]interface{})
	if !ok {
		return nil, ErrInvalidReference
	}
	return result, nil
}

func ParseRefStr(docName string, strRef string) (string, string, error) {
	index := strings.IndexByte(strRef, '#')
	if index < 0 {
		return "", "", errors.Wrapf(ErrInvalidReference, "%s#%s", docName, strRef)
	}
	document := strRef[:index]
	if document == "" {
		document = docName
	}
	pointer := strRef[index+1:]
	return document, pointer, nil
}
