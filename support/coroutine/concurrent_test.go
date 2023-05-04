package coroutine

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestConcurrent(t *testing.T) {
	c := NewConcurrent(10)

	start := time.Now()

	for i := 0; i <= 20; i++ {
		c.Run(func() {
			time.Sleep(1 * time.Second)
			fmt.Println(time.Now().String())
		})
	}

	assert.True(t, time.Since(start) < 2*time.Second+300*time.Millisecond)
}
