package generator

import uuid "github.com/satori/go.uuid"

var UUIDGenerator Generator = &uuidGenerator{}

type uuidGenerator struct{}

var _ Generator = (*uuidGenerator)(nil)

func (g *uuidGenerator) Generate() string {
	return uuid.NewV4().String()
}
