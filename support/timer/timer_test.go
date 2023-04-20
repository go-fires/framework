package timer

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTick(t *testing.T) {
	var r int

	Tick(1*time.Second, func(ticker *time.Ticker) {
		r++

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
	var flag = true

	After(1*time.Second, func() {
		flag = false
	})

	assert.True(t, flag)

	time.Sleep(2 * time.Second)
	assert.False(t, flag)
}
