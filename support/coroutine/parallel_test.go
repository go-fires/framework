package coroutine

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParallel(t *testing.T) {
	t.Skip("skipping test in race mode.")

	p := NewParallel(10)

	start := time.Now()

	for i := 0; i < 20; i++ {
		p.Add(func() {
			fmt.Println(time.Now().String())
			time.Sleep(1 * time.Second)
		})
	}

	p.Wait()

	assert.True(t, time.Since(start) < 2*time.Second+300*time.Millisecond)
}
