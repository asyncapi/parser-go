package jsonpath

import (
	"fmt"
	. "github.com/onsi/gomega"
	"testing"
)

func TestDecodeItemName(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "",
			expected: "",
		},
		{
			name:     "a",
			expected: "a",
		},
		{
			name:     "~1plz~1test~0me",
			expected: "/plz/test~me",
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf(`Expect:"%s"`, test.expected), func(t *testing.T) {
			g := NewWithT(t)
			actual, err := DecodeEntryKey(test.name)
			g.Expect(err).ShouldNot(HaveOccurred())
			g.Expect(actual).To(Equal(test.expected))
		})
	}
}

func TestDecodeItemNameErr(t *testing.T) {
	tests := []struct {
		name     string
		expected error
	}{
		{
			name: "~",
		},
		{
			name: "/",
		},
		{
			name: "~wrong~",
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf(`Expect err:"%s"`, test.name), func(t *testing.T) {
			g := NewWithT(t)
			_, err := DecodeEntryKey(test.name)
			g.Expect(err).Should(HaveOccurred())
		})
	}
}

func TestNewReference(t *testing.T) {
	tests := []struct {
		strRef   string
		expected Ref
	}{
		{
			strRef: "/test/path#/test/me/plz",
			expected:
			Ref{
				pointer: "/test/me/plz",
				uri:     "/test/path",
				path:    []string{"test", "me", "plz"},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.strRef, func(t *testing.T) {
			g := NewWithT(t)
			ref, err := NewRef(test.strRef)
			g.Expect(err).ShouldNot(HaveOccurred())
			g.Expect(ref).To(Equal(test.expected))
		})
	}
}

func TestNewReference_ShouldReturnError(t *testing.T) {
	for _, test := range []string{
		"123",
		"",
	} {
		t.Run(test, func(t *testing.T) {
			g := NewWithT(t)
			_, err := NewRef(test)
			g.Expect(err).Should(HaveOccurred())
		})
	}
}

func TestGetRefObject(t *testing.T) {
	g := NewWithT(t)
	v := map[string]interface{}{
		"test": map[string]interface{}{
			"me": map[string]interface{}{
				"plz": true,
			},
		},
	}
	actual, err := GetRefObject([]string{"test", "me"}, v)
	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(actual).To(Equal(map[string]interface{}{
		"plz": true,
	}))
}
