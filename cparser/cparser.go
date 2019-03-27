package main

// typedef struct {
//  char *result;
//  char **err;
//  int  errCount;
//  _Bool hasErrors;
// } _ParseResult_;
import "C"
import (
	"unsafe"
	"github.com/asyncapi/parser/pkg/hlsp"
)

//Parse is the C-friendly version of the parser.Parse method.
//export Parse
func Parse(yamlOrJSONDocument string) C._ParseResult_ {
	jsonDoc, err := hlsp.Parse([]byte(yamlOrJSONDocument))

	if err != nil {
		ea, count := makeCErrorArray(err)
		return C._ParseResult_{
			result:    nil,
			err:       ea,
			errCount:  count,
			hasErrors: count > 0,
		}
	}

	return C._ParseResult_{
		result:    C.CString(string(jsonDoc)),
		err:       nil,
		errCount:  0,
		hasErrors: false,
	}
}

func makeCErrorArray(err *hlsp.ParserError) (**C.char, C.int) {
	var arr []string
	arr = append(arr, err.Error())
	if err.ParsingErrors != nil {
		for _, msg := range err.ParsingErrors {
			arr = append(arr, string(msg.String()))
		}
	}
	cArray := C.malloc(C.size_t(len(arr)) * C.size_t(unsafe.Sizeof(uintptr(0))))

	// convert the C array to a Go Array so we can index it
	a := (*[1<<30 - 1]*C.char)(cArray)

	for idx, str := range arr {
		a[idx] = C.CString(str)
	}

	return (**C.char)(cArray), C.int(len(arr))
}

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library
}
