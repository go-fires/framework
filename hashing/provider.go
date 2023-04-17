package hashing

import (
	"github.com/go-fires/framework/contracts/container"
	"github.com/go-fires/framework/contracts/foundation"
)

const Hash = "hash"

type Provider struct {
	container.Container
	*foundation.UnimplementedProvider
}

var _ foundation.Provider = (*Provider)(nil)

func NewProvider(c container.Container) *Provider {
	return &Provider{
		Container: c,
	}
}

func (h *Provider) Register() {
	h.Singleton(Hash, func(c container.Container) interface{} {
		return NewManager(&Config{
			Driver: "bcrypt",
		})
	})
}
