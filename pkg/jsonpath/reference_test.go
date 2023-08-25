package jsonpath

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
			actual, err := DecodeEntryKey(test.name)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual)
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
			_, err := DecodeEntryKey(test.name)
			assert.Error(t, err)
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
			expected: Ref{
				pointer: "/test/me/plz",
				uri:     "/test/path",
				path:    []string{"test", "me", "plz"},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.strRef, func(t *testing.T) {
			ref, err := NewRef(test.strRef)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, ref)
		})
	}
}

func TestNewReference_ShouldReturnError(t *testing.T) {
	for _, test := range []string{
		"123",
		"",
	} {
		t.Run(test, func(t *testing.T) {
			_, err := NewRef(test)
			assert.Error(t, err)
		})
	}
}

func TestGetRefObject(t *testing.T) {
	v := map[string]interface{}{
		"test": map[string]interface{}{
			"me": map[string]interface{}{
				"plz": true,
			},
		},
	}
	actual, err := GetRefObject([]string{"test", "me"}, v)
	assert.NoError(t, err)

	expected := map[string]interface{}{
		"plz": true,
	}
	assert.Equal(t, expected, actual)
}
