// build:
package timer

import (
	"time"
)

func Tick(duration time.Duration, callback func(ticker *time.Ticker)) *time.Ticker {
	ticker := time.NewTicker(duration)

	go func() {
		for {
			<-ticker.C

			callback(ticker)
		}
	}()

	return ticker
}

// After Nonblocking call to callback after duration.
func After(duration time.Duration, callback func()) *time.Timer {
	return time.AfterFunc(duration, callback)
}
