package klog

import (
	"github.com/op/go-logging"
	"os"
	"io"
)

type KLogger struct {
	opLogger   *logging.Logger
	configured bool
}

type writer struct {
	logF func(template string, args ...interface{})
}

func (w *writer) Write(p []byte) (n int, err error) {
	w.logF(string(p))
	return len(p), nil
}

func (l *KLogger) Configure(name string, prefix string, timeFormat string, output io.Writer) {
	if l.configured {
		panic("KLogger already configured")
	}
	logger := logging.MustGetLogger(name)
	backend := logging.NewLogBackend(output, prefix, 0)
	format := "%{color}%{time:" + timeFormat + "} ▶ %{level:.4s} %{color:reset} %{message}"
	fmtBackend := logging.NewBackendFormatter(
		backend,
		logging.MustStringFormatter(format),
	)
	logging.AddModuleLevel(backend)
	logging.SetBackend(fmtBackend)
	logging.SetLevel(logging.DEBUG, "")

	l.opLogger = logger
	l.configured = true
}

func (l *KLogger) AsWriter(logF func(template string, args ...interface{})) io.Writer {
	return &writer{
		logF,
	}
}

func (l *KLogger) Printf(template string, args ...interface{}) {
	l.maybeConfigure()
	l.opLogger.Infof(template, args...)
}

func (l *KLogger) Debugf(template string, args ...interface{}) {
	l.maybeConfigure()
	l.opLogger.Debugf(template, args...)
}

func (l *KLogger) Errorf(template string, args ...interface{}) {
	l.maybeConfigure()
	l.opLogger.Errorf(template, args...)
}

func (l *KLogger) Fatalf(template string, args ...interface{}) {
	l.maybeConfigure()
	l.opLogger.Fatalf(template, args...)
}

/*
Configuring with default setting when the defaultLogger wasn't configured
 */
func (l *KLogger) maybeConfigure() {
	if !l.configured {
		l.Configure("", "", "2006-01-02T15:04:05", os.Stdout)
	}
}
