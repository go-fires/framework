package yaml

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type foo struct {
	Bar string `yaml:"bar"`
}

func TestParse(t *testing.T) {
	var f foo
	err := Parse("bar: baz", &f)
	assert.Nil(t, err)
	assert.Equal(t, "baz", f.Bar)
}

func TestDump(t *testing.T) {
	f := foo{Bar: "baz"}

	s, err := Dump(f)
	assert.Nil(t, err)
	assert.Equal(t, "bar: baz\n", s)
}
