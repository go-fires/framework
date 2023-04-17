package logger

import "fmt"

type Loggerable func(Level, string) error

func (l Loggerable) Emergencyf(format string, args ...interface{}) error {
	return l.Emergency(fmt.Sprintf(format, args...))
}

func (l Loggerable) Alertf(format string, args ...interface{}) error {
	return l.Alert(fmt.Sprintf(format, args...))
}

func (l Loggerable) Criticalf(format string, args ...interface{}) error {
	return l.Critical(fmt.Sprintf(format, args...))
}

func (l Loggerable) Errorf(format string, args ...interface{}) error {
	return l.Error(fmt.Sprintf(format, args...))
}

func (l Loggerable) Warningf(format string, args ...interface{}) error {
	return l.Warning(fmt.Sprintf(format, args...))
}

func (l Loggerable) Noticef(format string, args ...interface{}) error {
	return l.Notice(fmt.Sprintf(format, args...))
}

func (l Loggerable) Infof(format string, args ...interface{}) error {
	return l.Info(fmt.Sprintf(format, args...))
}

func (l Loggerable) Debugf(format string, args ...interface{}) error {
	return l.Debug(fmt.Sprintf(format, args...))
}

func (l Loggerable) Emergency(message string) error {
	return l.Log(Emergency, message)
}

func (l Loggerable) Alert(message string) error {
	return l.Log(Alert, message)
}

func (l Loggerable) Critical(message string) error {
	return l.Log(Critical, message)
}

func (l Loggerable) Error(message string) error {
	return l.Log(Error, message)
}

func (l Loggerable) Warning(message string) error {
	return l.Log(Warning, message)
}

func (l Loggerable) Notice(message string) error {
	return l.Log(Notice, message)
}

func (l Loggerable) Info(message string) error {
	return l.Log(Info, message)
}

func (l Loggerable) Debug(message string) error {
	return l.Log(Debug, message)
}

func (l Loggerable) Log(level Level, message string) error {
	return l(level, message)
}
