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

	for _, fn := range p.callbacks {
		p.concurrent.Run(func() {
			fn()
			p.wg.Done()
		})
	}

	p.wg.Wait()
}
