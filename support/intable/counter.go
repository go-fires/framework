package intable

import "sync"

type Counter struct {
	mu  sync.Mutex
	val int
}

func (c *Counter) Inc(n int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.val += n
}

func (c *Counter) Dec(n int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.val -= n
}

func (c *Counter) Val() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.val
}
