package dereferencer

import (
  "github.com/ghodss/yaml"
  "testing"

  "github.com/stretchr/testify/assert"
)

const (
  refInFile = `
channels:
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
      type: string
`
  expectedResolved = `
channels:
event/{streetlightId}/lighting/measured:
  parameters:
    - name: streetlightId
      description: The ID of the streetlight.
      schema:
      type: string
  subscribe:
    summary: Receive information about environmental lighting conditions of a particular streetlight.
    operationId: receiveLightMeasurement
`
)

func TestDereferenceInFile(t *testing.T) {
  jsonDocument, err := yaml.YAMLToJSON([]byte(refInFile))
  assert.NoError(t, err, "error converting yaml to json")
  resolvedDoc, err := Dereference(jsonDocument)
  assert.NoError(t, err, "error Dereferencing")
  resolvedYaml, err := yaml.JSONToYAML(resolvedDoc)
  assert.Equal(t, expectedResolved, string(resolvedYaml), "No equal")
}
