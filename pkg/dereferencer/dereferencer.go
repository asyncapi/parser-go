package dereferencer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"unicode/utf8"

	"github.com/pkg/errors"
)

type dereferencer interface {
	Dereference(ref string, document []byte) error
}

const (
	inFileRef = "#"
	httpRef   = "http://"
)

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func eachJSONValue(obj *interface{}, handler func(*string, *int, *interface{})) {
	if obj == nil {
		return
	}
	// Yield all key/value pairs for objects.
	o, isObject := (*obj).(map[string]interface{})
	if isObject {
		for k, v := range o {
			handler(&k, nil, &v)
			eachJSONValue(&v, handler)
		}
	}
	// Yield each index/value for arrays.
	a, isArray := (*obj).([]interface{})
	if isArray {
		for i, x := range a {
			handler(nil, &i, &x)
			eachJSONValue(&x, handler)
		}
	}
	// Do nothing for primitives since the handler got them.
}

// Dereference resolves all references in the document
func Dereference(document []byte, circularReferenceOption bool) (resolvedDoc []byte, err error) {
	var objmap map[string]interface{}
	err = json.Unmarshal(document, &objmap)
	if err != nil {
		return nil, err
	}

	var replacements = make(map[string]interface{})
	replacements, err = resolve(objmap, document)
	if err != nil {
		return nil, err
	}
	i := 1
	for len(replacements) > 0 {
		// Replace strings for its references
		for k, v := range replacements {
			find := []byte(fmt.Sprintf("{\"$ref\":\"%s\"}", k))
			document = bytes.Replace(document, find, v.([]byte), -1)
		}
		var objmap map[string]interface{}
		err = json.Unmarshal(document, &objmap)
		if err != nil {
			return nil, err
		}
		var oldReplacements = make(map[string]interface{})
		oldReplacements = copyMap(replacements)
		replacements = make(map[string]interface{})
		replacements, err = resolve(objmap, document)
		// After 9 loops or multiples of 10 resolve circulars
		if i%10 == 0 {
			replacements, err = resolveCircular(oldReplacements, replacements, circularReferenceOption)
			if err != nil {
				return nil, errors.WithMessage(err, "failed to resolve circular references")
			}
		}
		if err != nil {
			return nil, err
		}
		i++

		if i >= 1000 {
			return document, fmt.Errorf("error finding references, check the format of your document please")
		}
	}
	resolvedDoc = document
	return resolvedDoc, nil
}

func copyMap(originalMap map[string]interface{}) map[string]interface{} {
	mapCopy := make(map[string]interface{})
	// Copy from the original map to the target map
	for key, value := range originalMap {
		mapCopy[key] = value
	}
	return mapCopy
}

func resolveCircular(oldReplacements, newReplacements map[string]interface{}, circularReferenceOption bool) (map[string]interface{}, error) {
	for k := range oldReplacements {
		if newReplacements[k] != nil {
			if !circularReferenceOption {
				return newReplacements, fmt.Errorf("you a circular reference at %s please review it", k)
			}
			newReplacements[k] = []byte("{\"circular\": \"circular\"}")
		}
	}
	return newReplacements, nil
}

func resolve(objmap map[string]interface{}, document []byte) (replacements map[string]interface{}, err error) {
	replacements = make(map[string]interface{})
	fDef := fileDereferencer{}
	httpDef := httpDereferencer{}
	var dv []byte
	for _, v := range objmap {
		eachJSONValue(&v, func(key *string, index *int, value *interface{}) {
			if key != nil { // It's an object key/value pair...
				if *key == "$ref" {
					ref := (*value).(string)

					if strings.HasPrefix(ref, inFileRef) {
						dv, err = fDef.Dereference(ref, document)
					} else if strings.HasPrefix(ref, httpRef) {
						dv, ref, err = resolveURL(ref)
						if err != nil {
							log.Fatalf("failed to resolve URL: %v\n", err)
						}
						if ref != "" {
							dv, err = httpDef.Dereference(ref, dv)
						}
					} else {
						dv, ref, err = checkFile(ref)
						if err != nil {
							log.Fatalf("failed to check file: %v\n", err)
						}
						if ref != "" {
							dv, err = fDef.Dereference(ref, dv)
						}
					}

					if err != nil {
						log.Fatalf("failed to process reference %q: %v\n", (*value).(string), err)
					}

					//we have to marshal document back in order to be able to substitute reference "{\"$ref\":\"%s\"}"
					//with the dereferenced document
					replacements[(*value).(string)], err = json.Marshal(json.RawMessage(dv))
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		})
	}
	return
}
