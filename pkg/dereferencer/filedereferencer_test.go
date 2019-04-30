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

	circularRef = `definitions:
  child:
    title: child
    type: object
    properties:
      name:
        type: string
      pet:
        $ref: '#/definitions/pet'

  thing:
    $ref: "#/definitions/thing"         # <--- circular reference to self

  pet:
    title: pet
    type: object
    properties:
      name:
        type: string
      age:
        type: number
      species:
        type: string
        enum:
          - cat
          - dog
          - bird
          - fish`

	indirectCircularRef = `definitions:
    parent:
      title: parent
      properties:
        name:
          type: string
        children:
          type: array
          items:
            $ref: "#/definitions/child"       # <--- indirect circular reference
  
    child:
      title: child
      properties:
        name:
          type: string
        pet:
          $ref: '#/definitions/pet'
        parents:
          type: array
          items:
            $ref: "#/definitions/parent"      # <--- indirect circular reference
  
    pet:
      title: pet
      type: object
      properties:
        name:
          type: string
        age:
          type: number
        species:
          type: string
          enum:
            - cat
            - dog
            - bird
            - fish`

	ancestorCircularRef = `definitions:
  person:
    title: person
    properties:
      name:
        type: string
      pet:
        $ref: '#/definitions/pet'
      spouse:
        $ref: "#/definitions/person"       # <--- circular reference to ancestor
      age:
        type: number

  pet:
    title: pet
    type: object
    properties:
      name:
        type: string
      age:
        type: number
      species:
        type: string
        enum:
          - cat
          - dog
          - bird
          - fish`
	indirectAncestorCircularRef = `definitions:
  parent:
    title: parent
    properties:
      name:
        type: string
      child:
        $ref: '#/definitions/child'           # <--- indirect circular reference

  child:
    title: child
    properties:
      name:
        type: string
      pet:
        $ref: '#/definitions/pet'
      children:
        description: children
        type: array
        items:
          $ref: '#/definitions/child'         # <--- circular reference to ancestor

  pet:
    title: pet
    type: object
    properties:
      name:
        type: string
      age:
        type: number
      species:
        type: string
        enum:
          - cat
          - dog
          - bird
          - fish`
	circularString = `{"circular": "circular"}`
)

func TestDereferenceInFile(t *testing.T) {
	jsonDocument, err := yaml.YAMLToJSON([]byte(refInFile))
	assert.NoError(t, err, "error converting yaml to json")
	resolvedDoc, err := Dereference(jsonDocument, true)
	assert.NoError(t, err, "error Dereferencing")
	assert.Contains(t, string(resolvedDoc), expectedResolved, "does not contain resolved $ref")
}

func TestDereferenceExternalFile(t *testing.T) {
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
	resolvedDoc, err := Dereference(jsonDocument, true)
	assert.NoError(t, err, "error Dereferencing")
	assert.Contains(t, string(resolvedDoc), expectedResolved, "does not contain resolved $ref")
}

func TestCircularDereferencesInFile(t *testing.T) {
	circularReferences := []string{circularRef, indirectCircularRef, ancestorCircularRef, indirectAncestorCircularRef}
	for i := 0; i < len(circularReferences); i++ {
		jsonDocument, err := yaml.YAMLToJSON([]byte(circularReferences[i]))
		assert.NoError(t, err, "error converting yaml to json")
		resolvedDoc, err := Dereference(jsonDocument, true)
		assert.NoError(t, err, "error Dereferencing")
		assert.Contains(t, string(resolvedDoc), circularString, "does not contain circular String in circular $ref")
	}
}

func TestCircularDereferencesInFileFails(t *testing.T) {
	circularReferences := []string{circularRef, indirectCircularRef, ancestorCircularRef, indirectAncestorCircularRef}
	for i := 0; i < len(circularReferences); i++ {
		jsonDocument, err := yaml.YAMLToJSON([]byte(circularReferences[i]))
		assert.NoError(t, err, "error converting yaml to json")
		_, err = Dereference(jsonDocument, false)
		assert.Error(t, err, "Circulare reference not detected")
	}
}
