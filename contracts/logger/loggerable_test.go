package logger

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testLogger struct {
	Loggerable
}

var _ Logger = (*testLogger)(nil)

func NewTestLogger() Logger {
	return &testLogger{
		Loggerable: func(level Level, s string) error {
			return fmt.Errorf("%s: %s", level, s)
		},
	}
}

func TestLoggerable(t *testing.T) {
	l := NewTestLogger()

	assert.Errorf(t, l.Emergency("test"), "emergency: test")
	assert.Errorf(t, l.Alert("test"), "alert: test")
	assert.Errorf(t, l.Critical("test"), "critical: test")
	assert.Errorf(t, l.Error("test"), "error: test")
	assert.Errorf(t, l.Warning("test"), "warning: test")
	assert.Errorf(t, l.Notice("test"), "notice: test")
	assert.Errorf(t, l.Info("test"), "info: test")
	assert.Errorf(t, l.Debug("test"), "debug: test")
	assert.Errorf(t, l.Log(Info, "test"), "info: test")

	assert.Errorf(t, l.Emergencyf("test %s", "test"), "emergency: test test")
	assert.Errorf(t, l.Alertf("test %s", "test"), "alert: test test")
	assert.Errorf(t, l.Criticalf("test %s", "test"), "critical: test test")
	assert.Errorf(t, l.Errorf("test %s", "test"), "error: test test")
	assert.Errorf(t, l.Warningf("test %s", "test"), "warning: test test")
	assert.Errorf(t, l.Noticef("test %s", "test"), "notice: test test")
	assert.Errorf(t, l.Infof("test %s", "test"), "info: test test")
	assert.Errorf(t, l.Debugf("test %s", "test"), "debug: test test")
}
