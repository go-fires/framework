# Logger

Logger is a simple logger interface for Go. It is reference on the [PSR-3](https://www.php-fig.org/psr/psr-3/) standard.

## Features

- Simple interface
- 8 log levels
- [`stdout_logger.go`](stdout_logger.go) for `StdoutLogger

## Example

```go
package main

import (
	"fmt"
	"github.com/go-packagist/contracts/logger"
	"time"
)

type CustomLogger struct {
	logger.Loggerable
}

func NewCustomLogger(name string) *CustomLogger {
	return &CustomLogger{
		Loggerable: func(level logger.Level, message string) error {
			fmt.Println(
				fmt.Sprintf("[%s] %s.%s: %s",
					time.Now().Format("2006-01-02 15:04:05"),
					name, level.UpperString(), message))

			return nil
		},
	}
}

func main() {
	c := NewCustomLogger("test")

	c.Emergency("test")
	c.Alert("test")
	c.Critical("test")
	c.Error("test")
	c.Warning("test")
	c.Notice("test")
	c.Info("test")
	c.Debug("test")
	c.Log(logger.Info, "test")

	c.Emergencyf("test %s", "test")
	c.Alertf("test %s", "test")
	c.Criticalf("test %s", "test")
	c.Errorf("test %s", "test")
	c.Warningf("test %s", "test")
	c.Noticef("test %s", "test")
	c.Infof("test %s", "test")
	c.Debugf("test %s", "test")

	c.Log(logger.Info, "test")

	// Output:
	//
	//  [2023-04-09 19:26:45] test.EMERGENCY: test
	// 	[2023-04-09 19:26:45] test.ALERT: test
	// 	[2023-04-09 19:26:45] test.CRITICAL: test
	// 	[2023-04-09 19:26:45] test.ERROR: test
	// 	[2023-04-09 19:26:45] test.WARNING: test
	// 	[2023-04-09 19:26:45] test.NOTICE: test
	// 	[2023-04-09 19:26:45] test.INFO: test
	// 	[2023-04-09 19:26:45] test.DEBUG: test
	// 	[2023-04-09 19:26:45] test.INFO: test
	// 	[2023-04-09 19:26:45] test.EMERGENCY: test test
	// 	[2023-04-09 19:26:45] test.ALERT: test test
	// 	[2023-04-09 19:26:45] test.CRITICAL: test test
	// 	[2023-04-09 19:26:45] test.ERROR: test test
	// 	[2023-04-09 19:26:45] test.WARNING: test test
	// 	[2023-04-09 19:26:45] test.NOTICE: test test
	// 	[2023-04-09 19:26:45] test.INFO: test test
	// 	[2023-04-09 19:26:45] test.DEBUG: test test
	// 	[2023-04-09 19:26:45] test.INFO: test
}
```