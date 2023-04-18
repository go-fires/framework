package serializer

import (
	"encoding/json"
	"github.com/go-fires/framework/contracts/support"
)

type JsonSerializer struct {
}

var _ support.Serializable = (*JsonSerializer)(nil)

func (serializer *JsonSerializer) Serialize(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func (serializer *JsonSerializer) Unserialize(src []byte, dest interface{}) error {
	return json.Unmarshal(src, dest)
}
