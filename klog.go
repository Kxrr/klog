package klog

import "io"

var defaultLogger = Logger{}

func Printf(template string, args ...interface{}) {
	defaultLogger.Printf(template, args...)
}

func Debugf(template string, args ...interface{}) {
	defaultLogger.Debugf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	defaultLogger.Errorf(template, args...)
}

func Configure(name string, prefix string, output io.Writer) {
	defaultLogger.Configure(name, prefix, output)
}
