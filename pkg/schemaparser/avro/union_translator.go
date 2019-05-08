package avro

import "encoding/json"

// UnionAvro maps union avro scheme
type UnionAvro struct {
	Content string `json:"-"`
}

// UnmarshalJSON unmarshals JSON
func (value *UnionAvro) UnmarshalJSON(data []byte) error {
	type UnionAvroAlias UnionAvro
	jsonMap := UnionAvroAlias{}
	var err error
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		return err
	}

	value.Content = jsonMap.Content

	return nil
}

// MarshalJSON marshals JSON
func (value UnionAvro) MarshalJSON() ([]byte, error) {
	type UnionAvroAlias UnionAvro
	jsonIntArray, err := json.Marshal(&UnionAvroAlias{
		Content: value.Content,
	})
	if err != nil {
		return nil, err
	}
	return jsonIntArray, nil
}
