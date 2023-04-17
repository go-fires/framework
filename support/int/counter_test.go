package int

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkCounter(b *testing.B) {
	c := &Counter{}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Inc(1)
			c.Dec(1)
		}
	})

	assert.Equal(b, 0, c.Val())
}
