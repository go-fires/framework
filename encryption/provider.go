package encryption

import (
	"encoding/base64"
	"github.com/go-fires/framework/contracts/container"
	"github.com/go-fires/framework/contracts/foundation"
	"strings"
)

const EncrypterName = "encrypter"

type Provider struct {
	app foundation.Application

	*foundation.UnimplementedProvider
}

var _ foundation.Provider = (*Provider)(nil)

func NewProvider(app foundation.Application) *Provider {
	return &Provider{
		app: app,
	}
}

func (e *Provider) Register() {
	e.app.Singleton(EncrypterName, func(c container.Container) interface{} {
		return NewEncrypter(e.parseKey("base64:4IgTbtwvozDVo4DTaTEZME8n64yLSjAvdB+VkIj/OsA=")) // TODO: get key from config
	})
}

// parseKey parses the key and returns the result.
// If the key starts with "base64:", it will be decoded
// with base64.StdEncoding.
// Otherwise, it will be returned as-is.
func (e *Provider) parseKey(key string) string {
	if strings.Contains(key, "base64:") == false {
		return key
	}

	decoded, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(key, "base64:"))

	if err != nil {
		return ""
	}

	return string(decoded)
}
