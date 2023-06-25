package coroutine

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestConcurrent(t *testing.T) {
	t.Skip("skipping test in race mode.")

	c := NewConcurrent(10)
	defer c.Close()

	var (
		start = time.Now()
		wg    = sync.WaitGroup{}
	)

	for i := 0; i < 20; i++ {
		wg.Add(1)
		c.Run(func() {
			defer wg.Done()

			time.Sleep(1 * time.Second)
			fmt.Println(time.Now().String())
		})
	}

	wg.Wait()

	assert.True(t, time.Since(start) < 2*time.Second+300*time.Millisecond)
}
