package error

import (
	"testing"

	"github.com/pkg/errors"

	. "github.com/onsi/gomega"
)

var errTest = errors.New("test error")

func TestNewNil(t *testing.T) {
	g := NewWithT(t)
	err := New()
	g.Expect(err).To(BeNil())
}

func TestNew(t *testing.T) {
	g := NewWithT(t)
	err := New(errTest)
	g.Expect(err.Error()).To(Equal("test error"))
}

func TestJoin(t *testing.T) {
	testCases := []struct {
		desc            string
		errs            []error
		expectedMessage string
	}{
		{
			desc:            "empty",
			errs:            []error{},
			expectedMessage: "",
		},
		{
			desc:            "single error",
			errs:            []error{errTest},
			expectedMessage: "test error",
		},
		{
			desc:            "double error",
			errs:            []error{errTest, errTest},
			expectedMessage: "test error|test error",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			g := NewWithT(t)
			errMsg := Join(tC.errs, "|")
			g.Expect(errMsg).To(Equal(tC.expectedMessage))
		})
	}
}
