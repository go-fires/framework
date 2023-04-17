package container

import (
	"github.com/go-fires/framework/contracts/container"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainer_Instance(t *testing.T) {
	c := NewContainer()

	// string
	c.Instance("aa", "aaa")
	c.Instance("bb", "bbb")

	var aa, bb, cc string
	assert.Nil(t, c.Make("aa", &aa))
	assert.Nil(t, c.Make("bb", &bb))
	assert.Error(t, c.Make("cc", &cc))
	assert.Equal(t, "aaa", aa)
	assert.Equal(t, "bbb", bb)
	assert.Equal(t, "", cc)

	// callback
	c.Instance("dd", func() string {
		return "ddd"
	})

	var dd, ee func() string
	assert.Nil(t, c.Make("dd", &dd))
	assert.Nil(t, c.Make("dd", &ee))
	assert.Equal(t, "ddd", dd())
	assert.NotSame(t, dd, ee)

	// type struct
	type A struct {
		Name string
	}

	c.Instance("ff", A{Name: "fff"})
	var ff, gg A
	assert.Nil(t, c.Make("ff", &ff))
	assert.Nil(t, c.Make("ff", &gg))
	assert.Equal(t, "fff", ff.Name)
	assert.NotSame(t, ff, gg)

	c.Instance("gg", &A{Name: "fff"})
	var hh, ii *A
	assert.Nil(t, c.Make("gg", &hh))
	assert.Nil(t, c.Make("gg", &ii))
	assert.Equal(t, "fff", hh.Name)
	assert.Same(t, hh, ii)
}

type A struct {
	name  string
	value string
}

func NewA(name string) *A {
	return &A{
		name:  name,
		value: "value: " + name,
	}
}

func (a *A) GetName() string {
	return a.name
}

func (a *A) GetValue() string {
	return a.value
}

func TestContainer_Bind(t *testing.T) {
	c := NewContainer()

	// shared
	c.Bind("shared", func(c container.Container) interface{} {
		return NewA("shared")
	}, true)

	var shared *A
	assert.Nil(t, c.Make("shared", &shared))
	assert.Equal(t, "shared", shared.GetName())
	assert.Equal(t, "value: shared", shared.GetValue())

	var shared2 *A
	assert.Nil(t, c.Make("shared", &shared2))
	assert.Same(t, shared, shared2)

	// not shared
	c.Bind("not_shared", func(c container.Container) interface{} {
		return NewA("not_shared")
	}, false)

	var notShared *A
	assert.Nil(t, c.Make("not_shared", &notShared))
	assert.Equal(t, "not_shared", notShared.GetName())
	assert.Equal(t, "value: not_shared", notShared.GetValue())
	var notShared2 *A
	assert.Nil(t, c.Make("not_shared", &notShared2))
	assert.NotSame(t, notShared, notShared2)
}

func TestContainer_Singleton(t *testing.T) {
	c := NewContainer()

	c.Singleton("singleton", func(c container.Container) interface{} {
		return NewA("singleton")
	})

	var singleton *A
	assert.Nil(t, c.Make("singleton", &singleton))
	assert.Equal(t, "singleton", singleton.GetName())
	assert.Equal(t, "value: singleton", singleton.GetValue())

	var singleton2 *A
	assert.Nil(t, c.Make("singleton", &singleton2))
	assert.Same(t, singleton, singleton2)
}

func TestContainer_Make(t *testing.T) {
	c := NewContainer()

	c.Bind("aa", func(c container.Container) interface{} {
		return NewA("aa")
	}, true)

	c.Bind("bb", func(c container.Container) interface{} {
		return NewA("bb")
	}, false)

	var aa, bb, cc *A
	assert.Nil(t, c.Make("aa", &aa))
	assert.Nil(t, c.Make("bb", &bb))
	assert.Error(t, c.Make("cc", &cc))
	assert.Equal(t, "aa", aa.GetName())
	assert.Equal(t, "bb", bb.GetName())
}

func BenchmarkContainer(b *testing.B) {
	c := NewContainer()

	c.Bind("aa", func(c container.Container) interface{} {
		return NewA("aa")
	}, true)

	c.Bind("bb", func(c container.Container) interface{} {
		return NewA("bb")
	}, false)

	var aa, bb *A
	for i := 0; i < b.N; i++ {
		c.Make("aa", &aa)
		c.Make("bb", &bb)
	}
}
