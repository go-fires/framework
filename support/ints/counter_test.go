package ints

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func BenchmarkCounter(b *testing.B) {
	c := &Counter{}
	wg := sync.WaitGroup{}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			wg.Add(2)

			go func() {
				defer wg.Done()

				c.Inc(1)
			}()

			go func() {
				defer wg.Done()

				c.Dec(1)
			}()
		}
	})

	wg.Wait()

	assert.Equal(b, 0, c.Val())
}
