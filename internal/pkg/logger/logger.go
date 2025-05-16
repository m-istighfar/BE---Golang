package logger

import "io"

var Log Logger

type Logger interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Warn(args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
	WithFields(args map[string]interface{}) Logger
	GetWriter() io.Writer
	Printf(format string, args ...interface{})
}

func SetLogger(logger Logger) {
	Log = logger
}
