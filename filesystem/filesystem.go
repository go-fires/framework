package filesystem

import (
	"errors"
	"io/fs"
	"os"
)

var (
	ErrNotExist = fs.ErrNotExist
)

type Filesystem struct {
}

func (f *Filesystem) Exists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

func (f *Filesystem) Missing(path string) bool {
	return !f.Exists(path)
}

func (f *Filesystem) Get(path string) (string, error) {
	if f.Missing(path) {
		return "", ErrNotExist
	}

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer file.Close()

	var content string
	_, err = file.Read([]byte(content))
	if err != nil {
		return "", err
	}

	return content, nil
}

func (f *Filesystem) Put(path, content string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			file, err = os.Create(path)
			if err != nil {
				return 0, err
			}
		} else {
			return 0, err
		}
	}

	defer file.Close()

	return file.Write([]byte(content))
}
