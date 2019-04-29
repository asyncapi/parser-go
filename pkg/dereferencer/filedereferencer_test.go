package dereferencer

import (
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	refInFile = `channels:
event/{streetlightId}/lighting/measured:
  parameters:
    - $ref: '#/components/parameters/streetlightId'
  subscribe:
    summary: Receive information about environmental lighting conditions of a particular streetlight.
    operationId: receiveLightMeasurement
components:
  parameters:
    streetlightId:
      name: streetlightId
      description: The ID of the streetlight.
      schema:
      type: string`
	externalFileName = "references.json"
	refExternalFile  = `channels:
      event/{streetlightId}/lighting/measured:
        parameters:
          - $ref: '%s/references.json#/components/parameters/streetlightId'
        subscribe:
          summary: Receive information about environmental lighting conditions of a particular streetlight.
          operationId: receiveLightMeasurement`
	externalContent = `{"components":{"parameters":{"streetlightId":{"description":"The ID of the streetlight.","name":"streetlightId","schema":null,"type":"string"}}}}`

	expectedResolved = `[{"description":"The ID of the streetlight.","name":"streetlightId","schema":null,"type":"string"}]`
)

func TestDereferenceInFile(t *testing.T) {
	jsonDocument, err := yaml.YAMLToJSON([]byte(refInFile))
	assert.NoError(t, err, "error converting yaml to json")
	resolvedDoc, err := Dereference(jsonDocument)
	assert.NoError(t, err, "error Dereferencing")
	assert.Contains(t, string(resolvedDoc), expectedResolved, "does not contain resolved $ref")
}

func TestDereferenceExternalFile(t *testing.T) {
	// if runtime.GOOS == "windows" {
	// 	t.Skip("TODO: Fix it for windows.")
	// }
	externalFilePath := fmt.Sprintf("%s/%s", ".", externalFileName)
	f, err := os.OpenFile(externalFilePath, os.O_RDONLY|os.O_CREATE, 0666)
	err = ioutil.WriteFile(externalFilePath, []byte(externalContent), 0666)
	assert.NoError(t, err, "error writing to file")
	defer func() {
		f.Close()
		err := os.Remove(externalFilePath)
		assert.NoError(t, err, "error removing file")
	}()
	jsonDocument, err := yaml.YAMLToJSON([]byte(fmt.Sprintf(refExternalFile, ".")))
	assert.NoError(t, err, "error converting yaml to json")
	resolvedDoc, err := Dereference(jsonDocument)
	assert.NoError(t, err, "error Dereferencing")
	assert.Contains(t, string(resolvedDoc), expectedResolved, "does not contain resolved $ref")
}
