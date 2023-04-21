package container

import (
	"fmt"
	"sync"

	"github.com/go-fires/framework/contracts/container"
	"github.com/go-fires/framework/support/helper"
)

type binding struct {
	name     string
	concrete container.Concrete
	shared   bool
}

type Container struct {
	bindings  map[string]binding
	instances map[string]interface{}

	mu sync.Mutex
}

var _ container.Container = (*Container)(nil)

func NewContainer() *Container {
	return &Container{
		bindings:  make(map[string]binding),
		instances: make(map[string]interface{}),
	}
}

func (c *Container) Bind(name string, concrete container.Concrete, shared bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.bindings[name] = binding{
		name:     name,
		concrete: concrete,
		shared:   shared,
	}
}

func (c *Container) Singleton(name string, concrete container.Concrete) {
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
	// if an instance of the type is currently being managed as a shared
	if instance, ok := c.instances[name]; ok {
		return c.valueOf(instance, value)
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
		c.mu.Lock()
		defer c.mu.Unlock()

		c.instances[name] = concrete

		return c.valueOf(concrete, value)
	}

	return c.valueOf(concrete, value)
}

func (c *Container) valueOf(src interface{}, dst interface{}) error {
	return helper.ValueOf(src, dst)
}

func (c *Container) Has(id string) bool {
	_, ok := c.bindings[id]

	return ok
}

func (c *Container) Get(id string, value interface{}) error {
	return c.resolve(id, value)
}
