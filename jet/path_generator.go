package jet

import (
	"regexp"
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
	// $handledNamespace = Str::replaceLast('Service', '', end($handledNamespace));
	path := strings.ToLower(strings.Join(namespaces, "_"))

	if path[0] != '/' {
		path = "/" + path
	}

	return path + "/" + method
}

func (d *URLPathGenerator) snake(s string) string {
	sep := "_"

	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, "")
	s = regexp.MustCompile("[A-Z][a-z]").ReplaceAllStringFunc(s, func(s string) string {
		return sep + s
	})

	return strings.TrimLeft(strings.ToLower(s), sep)
}
