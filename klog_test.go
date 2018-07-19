package klog

import (
	"testing"
	"bytes"
	"regexp"
)

func TestAutoConfigured(t *testing.T) {
	if defaultLogger.configured != false {
		t.Fatalf("Expect default defaultLogger not be configured")
	}
	Printf("foo")
	if defaultLogger.configured != true {
		t.Fatalf("Expect default defaultLogger be configured after call Printf")
	}
	Printf("bar")

}

func TestConfigure(t *testing.T) {
	buf := bytes.Buffer{}
	Configure("Test", "Test - ", &buf)
	Printf("baz")
	r := regexp.MustCompile("Test - .*baz")
	if !r.Match(buf.Bytes()) {
		t.Fatalf("Incorrect output: %#v", buf.String())
	}
}
