package error

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pkg/errors"
)

var errTest = errors.New("test error")

func TestNewNil(t *testing.T) {
	err := New()
	assert.NoError(t, err)
}

func TestNew(t *testing.T) {
	err := New(errTest)
	assert.Error(t, err, "test error")
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
			errMsg := Join(tC.errs, "|")
			assert.Equal(t, tC.expectedMessage, errMsg)
		})
	}
}
