package system

import (
	"github.com/stretchr/testify/assert"
	"os"
	"runtime"
	"testing"
)

func TestNumCPU(t *testing.T) {
	assert.True(t, NumCPU() > 0)
	assert.Equal(t, runtime.NumCPU(), NumCPU())
}

func TestHostname(t *testing.T) {
	hostname, _ := os.Hostname()

	assert.Equal(t, hostname, Hostname())
	assert.True(t, len(Hostname()) > 0)
}

func TestIsWindows(t *testing.T) {
	assert.Equal(t, runtime.GOOS == "windows", IsWindows())
}

func TestLocalIP(t *testing.T) {
	assert.True(t, len(LocalIP()) > 0)
	assert.Equal(t, LocalIP(), LocalIPs()[0])
}
