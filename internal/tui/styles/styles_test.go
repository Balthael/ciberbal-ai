package styles

import (
	"strings"
	"testing"
)

// TestTaglineContainsCiberbalBranding verifies that Tagline() returns
// ciberbal-ai branding, not the old gentle-ai branding.
func TestTaglineContainsCiberbalBranding(t *testing.T) {
	tag := Tagline("1.0.0")
	if !strings.Contains(tag, "ciberbal-ai") {
		t.Fatalf("Tagline() = %q, want string containing 'ciberbal-ai'", tag)
	}
}

// TestTaglineContainsVersion verifies that Tagline() embeds the version string.
func TestTaglineContainsVersion(t *testing.T) {
	tag := Tagline("2.3.4")
	if !strings.Contains(tag, "2.3.4") {
		t.Fatalf("Tagline(%q) = %q, does not contain version", "2.3.4", tag)
	}
}
