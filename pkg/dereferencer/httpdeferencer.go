package dereferencer

type httpDereferencer struct {
	f                  dereferencer
	ref                []byte
	yamlOrJSONDocument []byte
}

func (htpp *httpDereferencer) Dereference(ref string, document []byte) error {
	return nil
}
