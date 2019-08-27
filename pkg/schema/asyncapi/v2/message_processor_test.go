package v2

import (
	. "github.com/onsi/gomega"

	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"
)

func Test_schemaFormat(t *testing.T) {
	type args struct {
		m *map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "schemaFormat present",
			args: args{
				m: &map[string]interface{}{
					"schemaFormat": "test",
				},
			},
			want: "test",
		},
		{
			name: "schemaFormat missing",
			args: args{
				m: &map[string]interface{}{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			format := schemaFormat(tt.args.m)
			g.Expect(format).To(Equal(tt.want))
		})
	}
}

func TestDispatcher_Add(t *testing.T) {
	type args struct {
		pm     func(interface{}) error
		labels []string
	}
	tests := []struct {
		name    string
		d       Dispatcher
		args    args
		wantErr bool
	}{
		{
			name: "schema parser with multiple labels",
			d:    map[string]func(interface{}) error{},
			args: args{
				pm: func(i interface{}) error {
					return nil
				},
				labels: []string{"test1", "test2"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			err := tt.d.Add(tt.args.pm, tt.args.labels...)
			g.Expect(err).ShouldNot(HaveOccurred())
			for _, label := range tt.args.labels {
				g.Expect(tt.d[label]).NotTo(BeNil())
			}
		})
	}
}

func Test_extractMessages(t *testing.T) {
	tests := []struct {
		name        string
		channelFile string
		expectedMsg string
		wantErr     bool
	}{
		{
			name:        "just message",
			channelFile: "./testdata/given/anyOf_channel.json",
			expectedMsg: "./testdata/expected/expected_anyOf_messages.json",
			wantErr:     false,
		},
		{
			name:        "just message",
			channelFile: "./testdata/given/oneOf_channel.json",
			expectedMsg: "./testdata/expected/expected_oneOf_messages.json",
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		var channel interface{}
		err := load(tt.channelFile, &channel, t)
		if err != nil {
			panic(fmt.Sprintf("invalid test data in: '%s'", tt.channelFile))
		}
		var expectedMsgs []*map[string]interface{}
		err = load(tt.expectedMsg, &expectedMsgs, t)
		if err != nil {
			panic(fmt.Sprintf("invalid test data in: '%s'", tt.channelFile))
		}
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			got, err := extractMessages(channel)
			if tt.wantErr {
				g.Expect(err).Should(HaveOccurred())
				return
			}
			g.Expect(got).To(Equal(expectedMsgs))
		})
	}
}

func TestBuildMessageProcessor(t *testing.T) {
	g := NewWithT(t)
	testErr := errors.New("test error")
	d := Dispatcher{
		"test1": func(_ interface{}) error {
			return nil
		},
		"test2": func(_ interface{}) error {
			return testErr
		},
	}
	var document *map[string]interface{}
	docPath := "./testdata/given/anyofdoc.json"
	err := load(docPath, &document, t)
	if err != nil {
		panic(fmt.Sprintf("invalid test data in: '%s'", docPath))
	}
	processMessages := BuildMessageProcessor(d)
	err = processMessages(document)
	g.Expect(err).Should(HaveOccurred())
}

func load(path string, v interface{}, t *testing.T) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	err = json.NewDecoder(file).Decode(v)
	if err != nil {
		return err
	}
	return nil
}
