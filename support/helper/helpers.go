package helper

import (
	"fmt"
	"reflect"
)

// Tap calls the given callback with the given value then returns the value.
func Tap(value interface{}, callbacks ...func(interface{})) interface{} {
	for _, callback := range callbacks {
		callback(value)
	}

	return value
}

// With calls the given callbacks with the given value then returns the value.
func With(value interface{}, callbacks ...func(interface{}) interface{}) interface{} {
	for _, callback := range callbacks {
		value = callback(value)
	}

	return value
}

// ValueOf sets the value of dest to the value of src.
func ValueOf(src interface{}, dest interface{}) error {
	rv := reflect.ValueOf(dest)

	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("dest must be a pointer")
	}

	if rv.IsNil() {
		return fmt.Errorf("dest must not be nil")
	}

	rv.Elem().Set(reflect.ValueOf(src))

	return nil
}
