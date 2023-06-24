package facade

import (
	foundation2 "github.com/go-fires/fires/x/foundation"
)

func App() *foundation2.Application {
	return foundation2.GetInstance()
}
