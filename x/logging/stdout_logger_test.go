package logging

import (
	"testing"

	"github.com/go-fires/fires/contracts/logger"
)

func TestStdoutLogger(t *testing.T) {
	s := NewStdoutLogger("prod")

	s.Emergency("emergency")
	s.Alert("alert")
	s.Critical("critical")
	s.Error("error")
	s.Warning("warning")
	s.Notice("notice")
	s.Info("info")
	s.Debug("debug")

	s.Emergencyf("emergency %s", "test")
	s.Alertf("alert %s", "test")
	s.Criticalf("critical %s", "test")
	s.Errorf("error %s", "test")
	s.Warningf("warning %s", "test")
	s.Noticef("notice %s", "test")
	s.Infof("info %s", "test")
	s.Debugf("debug %s", "test")

	s.Log(logger.Emergency, "emergency")
}
