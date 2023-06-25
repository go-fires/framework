package deps

import (
	"github.com/golang-module/carbon/v2"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCarbon(t *testing.T) {
	assert.Equal(t, time.Now().Add(-24*time.Hour).Format("2006-01-02"), carbon.Yesterday().Format("Y-m-d"))
}

func TestDotenv(t *testing.T) {
	assert.Error(t, godotenv.Load())
}

func TestViper(t *testing.T) {
	viper.Set("foo", "bar")
	assert.Equal(t, "bar", viper.GetString("foo"))
}
