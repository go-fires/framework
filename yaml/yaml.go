package yaml

import (
	"gopkg.in/yaml.v3"
)

func Parse(s string, v interface{}) error {
	return yaml.Unmarshal([]byte(s), v)
}

func Dump(v interface{}) (string, error) {
	bytes, err := yaml.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
