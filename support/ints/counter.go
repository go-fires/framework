package ints

import "sync"

type Counter struct {
	rw  sync.RWMutex
	val int
}

func (c *Counter) Inc(n int) {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.val += n
}

func (c *Counter) Dec(n int) {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.val -= n
}

func (c *Counter) Val() int {
	c.rw.RLock()
	defer c.rw.RUnlock()

	return c.val
}
