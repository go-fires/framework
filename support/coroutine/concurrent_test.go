package coroutine

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestConcurrent(t *testing.T) {
	c := NewConcurrent(10)

	var (
		start = time.Now()
		wg    = sync.WaitGroup{}
	)

	for i := 0; i < 20; i++ {
		c.Run(func() {
			wg.Add(1)
			defer wg.Done()

			time.Sleep(1 * time.Second)
			fmt.Println(time.Now().String())
		})
	}

	wg.Wait()

	assert.True(t, time.Since(start) < 2*time.Second+300*time.Millisecond)
}
