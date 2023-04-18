package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSupport_Tap(t *testing.T) {
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

func TestSupport_With(t *testing.T) {
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

func TestSupport_ValueOf(t *testing.T) {
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
