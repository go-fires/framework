package macroable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var macros = NewMacroable()

type test struct {
	Name string

	*Macroable
}

func newTest(name string) *test {
	t := &test{
		Name: name,
	}

	t.Macroable = macros.WithCtx(t)

	return t
}

func TestMacroable(t *testing.T) {
	macros.Macro("foo", func(ts *test) string {
		return "bar"
	})

	macros.Macro("bar", func(ts *test, bar string) string {
		return ts.Name + bar
	})

	macros.Macro("modify", func(ts *test, bar string) string {
		ts.Name = bar

		return ts.Name
	})

	ts := newTest("foo")
	assert.Equal(t, "bar", ts.Call("foo"))
	assert.Equal(t, "foobar", ts.Call("bar", "bar"))
	assert.Equal(t, "bar-modify", ts.Call("modify", "bar-modify"))

	ts2 := newTest("bar")
	assert.Equal(t, "bar", ts2.Call("foo"))
	assert.Equal(t, "barbar", ts2.Call("bar", "bar"))
	assert.Equal(t, "bar2-modify", ts2.Call("modify", "bar2-modify"))

	assert.Equal(t, "bar-modify", ts.Name)
	assert.Equal(t, "bar2-modify", ts2.Name)
}
