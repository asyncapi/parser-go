package v2

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse_VersionNotFound(t *testing.T) {
	b := []byte(`
{
  "asyncapi": "9.9.9-fake",
  "info": {
    "title": "Test",
    "version": "1.0.0"
  },
  "channels": {
    "test": {
      "publish": {
        "message": {
          "name": "foo",
          "payload": {
            "properties": {
              "thing": {
                "type": "integer"
              }
            },
            "type": "object"
          }
        }
      }
    }
  }
}
`)

	data := make(map[string]interface{})
	require.NoError(t, json.Unmarshal(b, &data))
	assert.EqualError(t, Parse(&data), `version "9.9.9-fake" is not supported`)
}

func TestParse_AsyncAPIFieldMissing(t *testing.T) {
	b := []byte(`
{
  "info": {
    "title": "Test",
    "version": "1.0.0"
  },
  "channels": {
    "test": {
      "publish": {
        "message": {
          "name": "foo",
          "payload": {
            "properties": {
              "thing": {
                "type": "integer"
              }
            },
            "type": "object"
          }
        }
      }
    }
  }
}
`)

	data := make(map[string]interface{})
	require.NoError(t, json.Unmarshal(b, &data))
	assert.EqualError(t, Parse(&data), "error extracting AsyncAPI Spec version from provided document: the `asyncapi` field is missing")
}

func TestSpec_2_1_0(t *testing.T) {
	// 1. defaultContentType
	// 2. name and summary fields on examples
	// 3. ibmmq new protocol
	b := []byte(`
{
  "asyncapi": "2.2.0",
  "info": {
    "title": "Test",
    "version": "1.0.0"
  },
  "defaultContentType": "application/json",
  "servers": {
    "production1": {
      "url": "ibmmq://qmgr1host:1414/qm1/DEV.APP.SVRCONN",
      "protocol": "ibmmq",
      "description": "Production Instance 1"
    }
  },
  "channels": {
    "test": {
      "publish": {
        "message": {
          "examples": [
            {
              "name": "Example for Foo Message",
              "payload": {
                "thing": 98221
              },
              "summary": "Example of an Foo message that contains a thing."
            }
          ],
          "name": "foo",
          "payload": {
            "properties": {
              "thing": {
                "type": "integer"
              }
            },
            "type": "object"
          }
        }
      }
    }
  }
}
`)

	data := make(map[string]interface{})
	require.NoError(t, json.Unmarshal(b, &data))
	assert.NoError(t, Parse(&data))
}

func TestSpec_2_2_0(t *testing.T) {
	// 1. assign channels to servers
	// 2. anypointmq new protocol
	b := []byte(`
{
  "asyncapi": "2.2.0",
  "info": {
    "title": "Test",
    "version": "1.0.0"
  },
  "servers": {
    "production1": {
      "url": "https://mq-eu-central-1.eu1.anypoint.my-asyncapi.com/api",
      "protocol": "anypointmq",
      "description": "Production Instance 1"
    }
  },
  "channels": {
    "test": {
      "servers": [
        "production1"
      ],
      "publish": {
        "message": {
          "name": "foo",
          "payload": {
            "properties": {
              "thing": {
                "type": "integer"
              }
            },
            "type": "object"
          }
        }
      }
    }
  }
}
`)

	data := make(map[string]interface{})
	require.NoError(t, json.Unmarshal(b, &data))
	assert.NoError(t, Parse(&data))
}
