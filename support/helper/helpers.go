package helper

import (
	"fmt"
	"io"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-fires/framework/support/try"
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
//
// Example:
//
//	var foo string
//	ValueOf("bar", &foo)
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

// Call calls the given function with the given params.
//
// Example:
//
//	Call(func(name string) string {
//		return "Hello " + name
//	}, "world")
func Call(fn interface{}, params ...interface{}) interface{} {
	return CallWithCtx(nil, fn, params...)
}

// CallWithCtx calls the given function with the given context and params.
//
// Example:
//
//	type Foo struct {
//		Name string
//	}
//
//	result := CallWithCtx(&Foo{Name: "Hello"}, func(ts *Foo, name string) string {
//		return ts.Name + name
//	}, "world")
func CallWithCtx(ctx interface{}, fn interface{}, params ...interface{}) interface{} {
	fv := reflect.ValueOf(fn)

	if fv.Kind() != reflect.Func {
		panic("fn must be a function")
	}

	var args []reflect.Value

	if ctx != nil {
		args = append(args, reflect.ValueOf(ctx))
	}

	for _, param := range params {
		args = append(args, reflect.ValueOf(param))
	}

	return fv.Call(args)[0].Interface()
}

// Dump dumps the given values.
// It's a wrapper for spew.Dump(https://github.com/davecgh/go-spew)
// This is useful for debugging.
func Dump(v ...interface{}) {
	spew.Dump(v...)
}

// Sdump dumps the given values and returns the result.
// It's a wrapper for spew.Sdump(https://github.com/davecgh/go-spew)
// This is useful for debugging.
func Sdump(v ...interface{}) string {
	return spew.Sdump(v...)
}

// Fdump dumps the given values to the given writer.
// It's a wrapper for spew.Fdump(https://github.com/davecgh/go-spew)
// This is useful for debugging.
func Fdump(w io.Writer, v ...interface{}) {
	spew.Fdump(w, v...)
}

// Try calls the given function and returns a Try instance.
// It's a wrapper for try.NewTry().
func Try(fn func()) *try.Try {
	return try.NewTry(fn)
}
