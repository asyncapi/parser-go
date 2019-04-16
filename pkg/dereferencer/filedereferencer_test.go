package dereferencer

import (
	"testing"
)


const (
	refInFile = `
	event.{streetlightId}.lighting.measured:
    parameters:
	  - $ref: '#/components/parameters/streetlightId'
	components:
	  parameters:
		streetlightId:
		  name: streetlightId
		  description: The ID of the streetlight.
		  schema:
			type: string
	`
	expectedResolved = `
	event.{streetlightId}.lighting.measured:
    parameters:
	  - name: streetlightId
		description: The ID of the streetlight.
		schema:
			type: string
	`
)

func TestDereferenceInFile(t *testing.T){

}