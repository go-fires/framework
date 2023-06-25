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

var testmessage string

func NewTestLogger() Logger {
	return &testLogger{
		Loggerable: func(level Level, s string) {
			testmessage = fmt.Sprintf("%s: %s", level, s)
		},
	}
}

func TestLoggerable(t *testing.T) {
	l := NewTestLogger()

	tests := []struct {
		name string
		fn   func()
		want string
	}{
		{"emergency", func() { l.Emergency("test") }, "emergency: test"},
		{"alert", func() { l.Alert("test") }, "alert: test"},
		{"critical", func() { l.Critical("test") }, "critical: test"},
		{"error", func() { l.Error("test") }, "error: test"},
		{"warning", func() { l.Warning("test") }, "warning: test"},
		{"notice", func() { l.Notice("test") }, "notice: test"},
		{"info", func() { l.Info("test") }, "info: test"},
		{"debug", func() { l.Debug("test") }, "debug: test"},

		{"emergencyf", func() { l.Emergencyf("test %s", "test") }, "emergency: test test"},
		{"alertf", func() { l.Alertf("test %s", "test") }, "alert: test test"},
		{"criticalf", func() { l.Criticalf("test %s", "test") }, "critical: test test"},
		{"errorf", func() { l.Errorf("test %s", "test") }, "error: test test"},
		{"warningf", func() { l.Warningf("test %s", "test") }, "warning: test test"},
		{"noticef", func() { l.Noticef("test %s", "test") }, "notice: test test"},
		{"infof", func() { l.Infof("test %s", "test") }, "info: test test"},
		{"debugf", func() { l.Debugf("test %s", "test") }, "debug: test test"},

		{"log", func() { l.Log(Emergency, "test") }, "emergency: test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			assert.Equal(t, tt.want, testmessage)
		})
	}
}
