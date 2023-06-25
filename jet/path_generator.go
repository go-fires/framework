package jet

import (
	"strings"
)

type PathGenerator interface {
	Generate(service string, method string) string
}

var DefaultPathGenerator PathGenerator = &URLPathGenerator{}

type URLPathGenerator struct{}

func (d *URLPathGenerator) Generate(service string, method string) string {
	namespaces := strings.Split(service, "\\")
	// namespaces[len(namespaces)-1] = strings.TrimSuffix(namespaces[len(namespaces)-1], "Service")
	path := strings.ToLower(strings.Join(namespaces, "_"))

	if path[0] != '/' {
		path = "/" + path
	}

	return path + "/" + method
}
