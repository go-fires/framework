package coroutine

import "sync"

type Parallel struct {
	callbacks  []func()
	concurrent *Concurrent
	wg         sync.WaitGroup
}

func NewParallel(max int) *Parallel {
	return &Parallel{
		callbacks:  []func(){},
		concurrent: NewConcurrent(max),
	}
}

func (p *Parallel) Add(fn ...func()) {
	p.callbacks = append(p.callbacks, fn...)
}

func (p *Parallel) Wait() {
	p.wg.Add(len(p.callbacks))
	defer p.wg.Wait()

	for _, fn := range p.callbacks {
		var f = fn

		p.concurrent.Run(func() {
			defer p.wg.Done()

			f()
		})
	}
}
