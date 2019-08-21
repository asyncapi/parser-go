package error

import "strings"

type Error struct {
	errors []error
}

var sep = "\n"

func New(errs ...error) error {
	if len(errs) < 1 {
		return nil
	}
	return &Error{
		errors: errs,
	}
}

func Join(a []error, sep string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return a[0].Error()
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i].Error())
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(a[0].Error())
	for _, e := range a[1:] {
		b.WriteString(sep)
		b.WriteString(e.Error())
	}
	return b.String()
}

func (e Error) Error() string {
	return Join(e.errors, sep)
}
