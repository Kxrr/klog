package klog

import (
	"github.com/op/go-logging"
	"os"
	"io"
)

type Logger struct {
	opLogger   *logging.Logger
	configured bool
}

func (l *Logger) Configure(name string, prefix string, output io.Writer) {
	if l.configured {
		panic("Logger already configured")
	}
	logger := logging.MustGetLogger(name)
	backend := logging.NewLogBackend(output, prefix, 0)
	fmtBackend := logging.NewBackendFormatter(
		backend,
		logging.MustStringFormatter(
			`%{color}%{time:15:04:05 MST} ▶ %{level:.4s} %{color:reset} %{message}`),
	)
	logging.AddModuleLevel(backend)
	logging.SetBackend(fmtBackend)
	logging.SetLevel(logging.DEBUG, "")

	l.opLogger = logger
	l.configured = true
}

func (l *Logger) Printf(template string, args ...interface{}) {
	l.maybeConfigure()
	l.opLogger.Infof(template, args...)

}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.maybeConfigure()
	l.opLogger.Debugf(template, args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.maybeConfigure()
	l.opLogger.Errorf(template, args...)
}


/*
Configuring with default setting when the defaultLogger wasn't configured
 */
func (l *Logger) maybeConfigure() {
	if !l.configured {
		l.Configure("", "", os.Stdout)
	}
}
