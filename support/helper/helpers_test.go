package helper

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTap(t *testing.T) {
	Tap("foo", func(value interface{}) {
		assert.Equal(t, "foo", value)
	})

	type Foo struct {
		Name string
	}

	f := Tap(&Foo{Name: "foo"}).(*Foo)
	assert.Equal(t, "foo", f.Name)

	f = Tap(&Foo{Name: "foo"}, func(foo interface{}) {
		foo.(*Foo).Name = "bar"
	}).(*Foo)
	assert.Equal(t, "bar", f.Name)

	f = Tap(&Foo{Name: "foo"}, func(foo interface{}) {
		foo.(*Foo).Name = "bar"
	}, func(foo interface{}) {
		foo.(*Foo).Name = "baz"
	}).(*Foo)
	assert.Equal(t, "baz", f.Name)
}

func TestWith(t *testing.T) {
	type Foo struct {
		Name string
	}

	f := With(&Foo{Name: "foo"}).(*Foo)
	assert.Equal(t, "foo", f.Name)

	f = With(&Foo{Name: "foo"}, func(foo interface{}) interface{} {
		foo.(*Foo).Name = "bar"
		return foo
	}).(*Foo)
	assert.Equal(t, "bar", f.Name)

	f = (With(&Foo{Name: "foo"}, func(foo interface{}) interface{} {
		foo.(*Foo).Name = "bar"
		return foo
	}, func(foo interface{}) interface{} {
		foo.(*Foo).Name = "baz"
		return foo
	})).(*Foo)
	assert.Equal(t, "baz", f.Name)
}

func TestValueOf(t *testing.T) {
	var foo string
	err := ValueOf("foo", &foo)
	assert.Nil(t, err)
	assert.Equal(t, "foo", foo)

	var bar int
	err = ValueOf(1, &bar)
	assert.Nil(t, err)
	assert.Equal(t, 1, bar)

	type Baz struct {
		Name string
	}
	var baz Baz
	err = ValueOf(Baz{
		Name: "baz",
	}, &baz)
	assert.Nil(t, err)
	assert.Equal(t, "baz", baz.Name)
}

func TestCall(t *testing.T) {
	result := Call(func() string {
		return "foo"
	})
	assert.Equal(t, "foo", result)

	result = Call(func(name string) string {
		return name
	}, "foo")
	assert.Equal(t, "foo", result)

	result = Call(func(name string, age int) string {
		return name + strconv.Itoa(age)
	}, "foo", 1)
	assert.Equal(t, "foo1", result)
}

func TestCallWithCtx(t *testing.T) {
	type Foo struct {
		Name string
	}

	result := CallWithCtx(&Foo{Name: "Hello"}, func(ts *Foo, name string) string {
		return ts.Name + name
	}, "world")
	assert.Equal(t, "Helloworld", result)

	assert.Panics(t, func() {
		CallWithCtx(&Foo{Name: "Hello"}, func(ts *Foo, name string) string {
			return ts.Name + name
		})
	})
}
