package filesystem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var f = &Filesystem{}

func TestFilesystem_Exists(t *testing.T) {
	assert.False(t, f.Exists("_testdata/non-existent-file"))

	f.Put("_testdata/existent-file", "")

	assert.True(t, f.Exists("_testdata/existent-file"))
}
