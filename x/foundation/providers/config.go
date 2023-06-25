package providers

import (
	"github.com/go-fires/fires/config"
	foundation2 "github.com/go-fires/fires/x/contracts/foundation"
)

type ConfigProvider struct {
	app foundation2.Application

	*foundation2.UnimplementedProvider
}

var _ foundation2.Provider = (*ConfigProvider)(nil)

func NewConfigProvider(app foundation2.Application) *ConfigProvider {
	return &ConfigProvider{
		app: app,
	}
}

func (p *ConfigProvider) Register() {
	p.app.Instance("config", config.NewConfig())
}
