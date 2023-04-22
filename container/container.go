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

func (c *Container) Has(id string) bool {
	if _, ok := c.instances[id]; ok {
		return true
	}
	if _, ok := c.bindings[id]; ok {
		return true
	}

	return false
}

func (c *Container) Get(id string) (interface{}, error) {
	return c.resolve(id)
}

func (c *Container) MustGet(id string) interface{} {
	instance, err := c.Get(id)
	if err != nil {
		panic(err)
	}

	return instance
}

func (c *Container) Make(name string, value interface{}) error {
	concrete, err := c.resolve(name)
	if err != nil {
		return err
	}

	return helper.ValueOf(concrete, value)
}

func (c *Container) Instance(name string, instance interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.instances[name] = instance
}
func (c *Container) resolve(name string) (interface{}, error) {
	// if an instance of the type is currently being managed as a shared
	if instance, ok := c.instances[name]; ok {
		return instance, nil
	}

	// if a binding exists for the name type
	binding, ok := c.bindings[name]
	if !ok {
		return nil, fmt.Errorf("no binding found for %s", name)
	}

	// if the concrete type is a function
	concrete := binding.concrete(c)

	// if the concrete type is shared
	if binding.shared {
		c.mu.Lock()
		c.instances[name] = concrete
		c.mu.Unlock()
	}

	return concrete, nil
}
