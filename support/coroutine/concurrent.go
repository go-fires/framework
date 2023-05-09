package coroutine

type Concurrent struct {
	channel chan struct{}
}

func NewConcurrent(max int) *Concurrent {
	return &Concurrent{
		channel: make(chan struct{}, max),
	}
}

// Run runs a function in a goroutine, but limits the number of goroutines running at the same time.
// Warning: Please pay attention to the scope of variables in the function.
func (c *Concurrent) Run(fn func()) {
	c.channel <- struct{}{}

	go func() {
		defer func() {
			<-c.channel
		}()

		fn()
	}()
}
