package size

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSize(t *testing.T) {
	assert.Equal(t, uint64(1536), Size(1536).Bytes())

	tests := []struct {
		size Size
		want string
	}{
		{Size(1), "1.00 B"},
		{Size(1024), "1.00 KB"},
		{Size(1024 * 1024), "1.00 MB"},
		{Size(1024 * 1024 * 1024), "1.00 GB"},
		{Size(1024 * 1024 * 1024 * 1024), "1.00 TB"},
		{Size(1024 * 1024 * 1024 * 1024 * 1024), "1.00 PB"},
		{Size(1024 * 1024 * 1024 * 1024 * 1024 * 1024), "1.00 EB"},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, test.size.String())
	}
}
