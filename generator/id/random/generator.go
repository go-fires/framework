package random

import (
	"github.com/go-fires/fires/generator"
	"github.com/go-fires/fires/support/strs"
)

type Generator struct{}

var _ generator.IDGenerator = (*Generator)(nil)

func (g *Generator) Generate() string {
	return strs.RandomString(32)
}

var global = &Generator{}

func Generate() string {
	return global.Generate()
}
