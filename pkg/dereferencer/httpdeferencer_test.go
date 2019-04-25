package dereferencer

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"github.com/ghodss/yaml"
	"testing"
	"github.com/stretchr/testify/assert"
)

const (
	refHTTPFile = `channels:
      event/{streetlightId}/lighting/measured:
        parameters:
          - $ref: '%s/references.json#/components/parameters/streetlightId'
        subscribe:
          summary: Receive information about environmental lighting conditions of a particular streetlight.
          operationId: receiveLightMeasurement`
)




func TestDereferenceHttp(t *testing.T){
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/references.json" {
				w.Header().Add("Content-Type", "application/json")
				w.Write([]byte(externalContent))
			}
		}),
	  )
	defer ts.Close()
	jsonDocument, err := yaml.YAMLToJSON([]byte(fmt.Sprintf(refHTTPFile, ts.URL)))
	assert.NoError(t, err, "error converting yaml to json")
	resolvedDoc, err := Dereference(jsonDocument)
	assert.NoError(t, err, "error Dereferencing")
	assert.Contains(t, string(resolvedDoc), expectedResolved, "does not contain resolved $ref")
}