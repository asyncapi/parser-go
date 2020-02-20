package v2

import parseSchema "github.com/asyncapi/parser-go/pkg/schema"

var (
	parser = parseSchema.NewParser(schema)
	Labels = []string{"asyncapi"}
)

func Parse(v interface{}) error {
	return parser.Parse(v)
}
