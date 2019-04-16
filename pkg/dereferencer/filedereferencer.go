package dereferencer

type fileDereferencer struct {
	f dereferencer
	ref []byte
	yamlOrJSONDocument []byte
}

func (fdef *fileDereferencer) Dereference(ref string, document []byte) error {
	return nil
}