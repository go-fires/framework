package generator

import "github.com/google/uuid"

var UUID Generator = &uuidGenerator{}

type uuidGenerator struct{}

var _ Generator = (*uuidGenerator)(nil)

func (g *uuidGenerator) Generate() string {
	return uuid.New().String()
}
