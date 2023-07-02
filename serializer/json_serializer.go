package serializer

import (
	"encoding/json"
)

var Json Serializer = &jsonSerializer{}

type jsonSerializer struct{}

func (s *jsonSerializer) Serialize(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func (s *jsonSerializer) Unserialize(src []byte, dest interface{}) error {
	return json.Unmarshal(src, dest)
}
