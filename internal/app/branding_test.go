package app

import (
	"bytes"
	"strings"
	"testing"
)

// TestVersionOutputUsesCiberbalBranding verifies that the `version` command
// prints "ciberbal-ai" as the binary name, not "ciberbal-ai".
func TestVersionOutputUsesCiberbalBranding(t *testing.T) {
	var buf bytes.Buffer
	err := RunArgs([]string{"version"}, &buf)
	if err != nil {
		t.Fatalf("version should not fail: %v", err)
	}
	if !strings.Contains(buf.String(), "ciberbal-ai") {
		t.Errorf("version output should contain 'ciberbal-ai'; got: %q", buf.String())
	}
}

func TestHelpOutputUsesCiberbalBranding(t *testing.T) {
	var buf bytes.Buffer
	err := RunArgs([]string{"help"}, &buf)
	if err != nil {
		t.Fatalf("help should not fail: %v", err)
	}
	if !strings.Contains(buf.String(), "ciberbal-ai") {
		t.Errorf("help output should contain 'ciberbal-ai'; got: %q", buf.String())
	}
}
