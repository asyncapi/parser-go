package v2

import (
	. "github.com/onsi/gomega"

	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		msg      string
		expected string
	}{
		{
			name:     "parse",
			msg:      "./testdata/given/msg1.json",
			expected: "./testdata/expected/msg1.json",
		},
		{
			name:     "with nullable and without type",
			msg:      "./testdata/given/with_nullable_and_without_type.json",
			expected: "./testdata/expected/with_nullable_and_without_type.json",
		},
		{
			name:     "without nullable and without type",
			msg:      "./testdata/given/without_nullable_and_without_type.json",
			expected: "./testdata/expected/without_nullable_and_without_type.json",
		},
		{
			name:     "parse without examples",
			msg:      "./testdata/given/parse_without_examples.json",
			expected: "./testdata/expected/parse_without_examples.json",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := NewWithT(t)
			msg := loadMsg(test.msg, t)
			err := Parse(&msg)
			g.Expect(err).ShouldNot(HaveOccurred())

			payload, err := json.Marshal(msg)
			if err != nil {
				t.Fatal("data marshalling error", err)
			}
			expected := open(test.expected, t)
			g.Expect(payload).Should(MatchJSON(expected))
		})
	}
}

func open(path string, t *testing.T) []byte {
	file, err := os.Open(path)
	if err != nil {
		t.Fatal("loading data", err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal("reading data", err)
	}
	return bytes
}

func loadMsg(path string, t *testing.T) map[string]interface{} {
	var v map[string]interface{}
	file, err := os.Open(path)
	if err != nil {
		t.Fatal("loading data", err)
	}
	err = json.NewDecoder(file).Decode(&v)
	if err != nil {
		t.Fatal("unmarshalling data", err)
	}
	return v
}
