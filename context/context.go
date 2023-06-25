package context

import (
	"errors"
	"sync"
)

var ctx = New()

func Set(key string, value interface{}) {
	ctx.Set(key, value)
}

func Get(key string) (interface{}, error) {
	return ctx.Get(key)
}

func Has(key string) bool {
	return ctx.Has(key)
}

func Delete(key string) {
	ctx.Delete(key)
}

func Clear() {
	ctx.Clear()
}

func All() map[string]interface{} {
	return ctx.All()
}

var (
	ErrNotFound = errors.New("context: key not found")
)

type Context struct {
	data map[string]interface{}
	rw   sync.RWMutex
}

func New() *Context {
	return &Context{
		data: make(map[string]interface{}),
	}
}

func (c *Context) Set(key string, value interface{}) {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.data[key] = value
}

func (c *Context) Get(key string) (interface{}, error) {
	c.rw.RLock()
	defer c.rw.RUnlock()

	if value, ok := c.data[key]; ok {
		return value, nil
	} else {
		return nil, ErrNotFound
	}
}

func (c *Context) Has(key string) bool {
	c.rw.RLock()
	defer c.rw.RUnlock()

	_, ok := c.data[key]

	return ok
}

func (c *Context) Delete(key string) {
	c.rw.Lock()
	defer c.rw.Unlock()

	delete(c.data, key)
}

func (c *Context) Clear() {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.data = make(map[string]interface{})
}

func (c *Context) All() map[string]interface{} {
	c.rw.RLock()
	defer c.rw.RUnlock()

	return c.data
}
