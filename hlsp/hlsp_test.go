package hlsp

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/asyncapi/parser/models"
)

func TestA(t *testing.T) {
	yamlFile, err := os.Open("../asyncapi/2.0.0/example.yaml")
	if err != nil {
		t.Log(err)
	}
	defer yamlFile.Close()

	fileBytes, err := ioutil.ReadAll(yamlFile)
	if err != nil {
		t.Log(err)
	}

	jsonFile, e := Parse(fileBytes)
	if e != nil {
		t.Log(e.Error())
		t.Log(e.ParsingErrors())
	}

	t.Log(string(jsonFile))

	var AsyncAPI models.AsyncapiDocument
	err = json.Unmarshal(jsonFile, &AsyncAPI)
	if err != nil {
		t.Log(err)
	}
	j, _ := AsyncAPI.Channels["event/{streetlightId}/lighting/measured"].Subscribe.MarshalJSON()
	t.Log(string(j))

	// j, err := json.Marshal(AsyncAPI.Channels)
	// t.Log(string(j))
}
