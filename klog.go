package klog

import (
	"io"
	opLogging "github.com/op/go-logging"
)

var defaultLogger = NewLogger()

func NewLogger() *KLogger {
	return &KLogger{
		&opLogging.Logger{},
		false,
	}
}

func Printf(template string, args ...interface{}) {
	defaultLogger.Printf(template, args...)
}

func Debugf(template string, args ...interface{}) {
	defaultLogger.Debugf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	defaultLogger.Errorf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	defaultLogger.Fatalf(template, args...)
}

func Configure(name string, prefix string, timeFormat string, output io.Writer) {
	defaultLogger.Configure(name, prefix, timeFormat, output)
}
