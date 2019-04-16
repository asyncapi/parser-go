package dereferencer

import (
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	"log"
	"strings"
	"fmt"
	"encoding/json"
)

type dereferencer interface {
    Dereference(ref string, document []byte) error
}

const (
    inFileRef = "#"
    httpRef = "http://"
)

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
func Dereference(document []byte) (resolvedDoc []byte, err error){
    var objmap map[string]interface{}
    err = json.Unmarshal(document, &objmap)
    if err != nil {
        return nil, err
    }
    fDef := fileDereferencer{}
    httpDef := httpDereferencer{}
    for k, v := range objmap { 
        fmt.Printf("key[%s]\n", k)
        eachJSONValue(&v, func(key *string, index *int, value *interface{}) {
            if key != nil { // It's an object key/value pair...
                //fmt.Printf("OBJ: key=%q, value=%#v\n", *key, *value)
                if *key == "$ref" {
                    if  strings.HasPrefix((*value).(string), inFileRef){
                        fmt.Printf("inFileRef %s", (*value).(string))
                        err = fDef.Dereference((*value).(string), document)
                    } else if strings.HasPrefix((*value).(string), httpRef){
                        fmt.Printf("httpRef %s", (*value).(string))
                        err = httpDef.Dereference((*value).(string), document)
                    } else {
                        fileData, err := checkFile((*value).(string))
                        if err != nil {
                            fmt.Printf("can't detect which reference are you using for %s", (*value).(string))
                            log.Fatal(err)
                        }
                        err = fDef.Dereference((*value).(string), fileData)
                    }
                }
            }
        })
    }
    return resolvedDoc, nil
}

func checkFile(filename string) (fileData []byte, err error) {
    fileData, err = ioutil.ReadFile(filename)
    fmt.Printf("externalFileRef %s", filename)
    schemaLoader := gojsonschema.NewBytesLoader(fileData)
    documentLoader := gojsonschema.NewBytesLoader(fileData)
    result, err := gojsonschema.Validate(schemaLoader, documentLoader)
    if result.Valid() {
		return fileData, nil
    }
    return nil, err
}