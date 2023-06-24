package facade

import "github.com/go-fires/fires/foundation"

func App() *foundation.Application {
	return foundation.GetInstance()
}
