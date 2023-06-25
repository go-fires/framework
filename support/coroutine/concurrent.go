package coroutine

type Concurrent struct {
	ch chan struct{}
}

func NewConcurrent(max int) *Concurrent {
	return &Concurrent{
		ch: make(chan struct{}, max),
	}
}

// Run runs a function in a goroutine, but limits the amount goroutines running at the same time.
// Warning: Please pay attention to the scope of variables in the function.
func (c *Concurrent) Run(fn func()) {
	c.ch <- struct{}{}

	go func() {
		defer func() {
			<-c.ch
		}()

		fn()
	}()
}

// Close the concurrent ch
func (c *Concurrent) Close() {
	close(c.ch)
}
