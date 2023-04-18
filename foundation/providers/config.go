package providers

import (
	"github.com/go-fires/framework/config"
	"github.com/go-fires/framework/contracts/foundation"
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
