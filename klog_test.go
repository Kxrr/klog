package klog

import (
	"testing"
	"os"
)

func TestConfiguredWhenCallDefaultLoggerMethod(t *testing.T) {
	if defaultLogger.configured != false {
		t.Fatalf("Expect default defaultLogger not be configured")
	}
	Printf("hello, world from test")
	if defaultLogger.configured != true {
		t.Fatalf("Expect default defaultLogger be configured after call Printf")
	}
}

func TestConfigure(t *testing.T) {
	Configure("Test", "Test", os.Stdout)
}
