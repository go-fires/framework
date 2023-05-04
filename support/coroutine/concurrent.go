package coroutine

type Concurrent struct {
	channel chan struct{}
}

func NewConcurrent(max int) *Concurrent {
	return &Concurrent{
		channel: make(chan struct{}, max),
	}
}

func (c *Concurrent) Run(fn func()) {
	c.channel <- struct{}{}

	go func() {
		defer func() {
			<-c.channel
		}()

		fn()
	}()
}
