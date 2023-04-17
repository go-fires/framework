package logger

type Logger interface {
	Emergencyf(format string, args ...interface{}) error
	Alertf(format string, args ...interface{}) error
	Criticalf(format string, args ...interface{}) error
	Errorf(format string, args ...interface{}) error
	Warningf(format string, args ...interface{}) error
	Noticef(format string, args ...interface{}) error
	Infof(format string, args ...interface{}) error
	Debugf(format string, args ...interface{}) error

	Emergency(message string) error
	Alert(message string) error
	Critical(message string) error
	Error(message string) error
	Warning(message string) error
	Notice(message string) error
	Info(message string) error
	Debug(message string) error

	Log(level Level, message string) error
}
