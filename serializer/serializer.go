package serializer

type Serializer interface {
	Serialize(data interface{}) ([]byte, error)
	Unserialize(src []byte, dest interface{}) error
}
