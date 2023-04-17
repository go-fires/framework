package container

import (
	"fmt"
	"reflect"
	"sync"
)

type Concrete func(container *Container) interface{}

type binding struct {
	name     string
	concrete Concrete
	shared   bool
}

type Container struct {
	bindings  map[string]binding
	instances map[string]interface{}

	mu sync.Mutex
}

func NewContainer() *Container {
	return &Container{
		bindings:  make(map[string]binding),
		instances: make(map[string]interface{}),
	}
}

func (c *Container) Bind(name string, concrete Concrete, shared bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.bindings[name] = binding{
		name:     name,
		concrete: concrete,
		shared:   shared,
	}
}

func (c *Container) Singleton(name string, concrete Concrete) {
	c.Bind(name, concrete, true)
}

func (c *Container) Make(name string, value interface{}) error {
	return c.resolve(name, value)
}

func (c *Container) Instance(name string, instance interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.instances[name] = instance
}
func (c *Container) resolve(name string, value interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// if an instance of the type is currently being managed as a shared
	if instance, ok := c.instances[name]; ok {
		return c.value(instance, value)
	}

	// if a binding exists for the name type
	binding, ok := c.bindings[name]
	if !ok {
		return fmt.Errorf("no binding found for %s", name)
	}

	// if the concrete type is a function
	concrete := binding.concrete(c)

	// if the concrete type is shared
	if binding.shared {
		c.instances[name] = concrete
		return c.value(concrete, value)
	}

	return c.value(concrete, value)
}

func (c *Container) value(src interface{}, dst interface{}) error {
	rv := reflect.ValueOf(dst)

	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("dst must be a pointer")
	}

	if rv.IsNil() {
		return fmt.Errorf("dst must not be nil")
	}

	rv.Elem().Set(reflect.ValueOf(src))

	return nil
}

func (c *Container) Has(id string) bool {
	_, ok := c.bindings[id]

	return ok
}

func (c *Container) Get(id string, value interface{}) error {
	return c.resolve(id, value)
}
