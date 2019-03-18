package main

// typedef struct {
//  char *result;
//  char *err;
// } _ParseResult_;
import "C"
import (
	"fmt"

	"github.com/asyncapi/parser/hlsp"
)

//export Parse
func Parse(yamlOrJSONDocument string) C._ParseResult_ {
	fmt.Println(string(yamlOrJSONDocument))
	jsonDoc, err := hlsp.Parse([]byte(yamlOrJSONDocument))

	if err != nil {
		return C._ParseResult_{
			result: nil,
			err:    C.CString(err.Error()),
		}
	}

	// var parsedAsyncAPI models.ParsedAsyncAPI
	// err = json.Unmarshal(jsonDoc, &parsedAsyncAPI)
	// if err != nil {
	//  return nil, errors.New(err.Error())
	// }

	return C._ParseResult_{
		result: C.CString(string(jsonDoc)),
		err:    nil,
	}
}

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library
}
