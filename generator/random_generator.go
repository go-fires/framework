package generator

import "github.com/go-fires/fires/support/strs"

var Random Generator = &randomGenerator{}

type randomGenerator struct{}

var _ Generator = (*randomGenerator)(nil)

func (g *randomGenerator) Generate() string {
	return strs.RandomString(32)
}
