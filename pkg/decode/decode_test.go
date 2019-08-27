package decode

import (
	. "github.com/onsi/gomega"

	"strings"
	"testing"
)

func TestToMap1(t *testing.T) {
	g := NewWithT(t)
	reader := strings.NewReader("123")
	_, err := ToMap(reader)
	g.Expect(err).Should(HaveOccurred())
}
