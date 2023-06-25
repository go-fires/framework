package serializer

import (
	"encoding/json"
)

type JsonSerializer struct{}

func (s *JsonSerializer) Serialize(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func (s *JsonSerializer) Unserialize(src []byte, dest interface{}) error {
	return json.Unmarshal(src, dest)
}
