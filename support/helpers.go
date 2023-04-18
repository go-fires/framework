package support

// Tap calls the given callback with the given value then returns the value.
func Tap(value interface{}, callbacks ...func(interface{})) interface{} {
	for _, callback := range callbacks {
		callback(value)
	}

	return value
}

func With(value interface{}, callbacks ...func(interface{}) interface{}) interface{} {
	for _, callback := range callbacks {
		value = callback(value)
	}

	return value
}
