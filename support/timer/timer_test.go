package timer

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestTick(t *testing.T) {
	var (
		r  int
		mu sync.Mutex
	)

	Tick(1*time.Second, func(ticker *time.Ticker) {
		mu.Lock()
		r++
		mu.Unlock()

		if r == 3 {
			ticker.Stop()
		}
	})

	assert.Equal(t, 0, r)

	// wait for 4 seconds
	time.Sleep(4 * time.Second)
	assert.Equal(t, 3, r)
}

func TestTimer_After(t *testing.T) {
	var (
		flag = true
		mu   sync.Mutex
	)

	After(1*time.Second, func() {
		mu.Lock()
		defer mu.Unlock()

		flag = false
	})

	assert.True(t, flag)

	time.Sleep(2 * time.Second)
	assert.False(t, flag)
}
