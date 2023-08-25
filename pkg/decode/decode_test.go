package decode

import (
	"github.com/stretchr/testify/assert"

	"strings"
	"testing"
)

func TestToMap1(t *testing.T) {
	reader := strings.NewReader("123")
	_, err := ToMap(reader)
	assert.Error(t, err)
}
