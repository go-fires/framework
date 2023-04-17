package container

type Concrete func(container Container) interface{}

type Container interface {
	Has(id string) bool
	Get(id string, value interface{}) error

	Bind(name string, concrete Concrete, shared bool)
	Singleton(name string, concrete Concrete)
	Instance(name string, instance interface{})
	Make(name string, value interface{}) error
}
