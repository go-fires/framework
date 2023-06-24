package providers

import (
	"github.com/go-fires/fires/config"
	"github.com/go-fires/fires/contracts/foundation"
)

type ConfigProvider struct {
	app foundation.Application

	*foundation.UnimplementedProvider
}

var _ foundation.Provider = (*ConfigProvider)(nil)

func NewConfigProvider(app foundation.Application) *ConfigProvider {
	return &ConfigProvider{
		app: app,
	}
}

func (p *ConfigProvider) Register() {
	p.app.Instance("config", config.NewConfig())
}
