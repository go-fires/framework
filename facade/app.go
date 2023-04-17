package facade

import "github.com/go-fires/framework/foundation"

func App() *foundation.Application {
	return foundation.GetInstance()
}
