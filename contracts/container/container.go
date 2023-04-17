package container

type Container interface {
	Has(id string) bool
	Get(id string, value interface{}) error
}
