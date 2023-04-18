package support

// Tap calls the given callback with the given value then returns the value.
func Tap(value interface{}, callback ...func(interface{})) interface{} {
	for _, c := range callback {
		c(value)
	}

	return value
}

func With(value interface{}, callback ...func(interface{}) interface{}) interface{} {
	for _, c := range callback {
		value = c(value)
	}

	return value
}
