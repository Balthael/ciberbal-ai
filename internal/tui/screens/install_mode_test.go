package screens

import (
	"strings"
	"testing"
)

// TestInstallModeOptionsReturnsExactlyTwo verifies that InstallModeOptions()
// returns exactly 2 options — Quick install and Advanced install.
func TestInstallModeOptionsReturnsExactlyTwo(t *testing.T) {
	opts := InstallModeOptions()
	if len(opts) != 2 {
		t.Fatalf("InstallModeOptions() returned %d options, want 2; got: %v", len(opts), opts)
	}
}

// TestInstallModeOptionNamesCorrect verifies the exact option labels.
func TestInstallModeOptionNamesCorrect(t *testing.T) {
	opts := InstallModeOptions()
	if opts[0] != "Quick install" {
		t.Errorf("opts[0] = %q, want %q", opts[0], "Quick install")
	}
	if opts[1] != "Advanced install" {
		t.Errorf("opts[1] = %q, want %q", opts[1], "Advanced install")
	}
}

// TestRenderInstallModeIsNonEmpty verifies that RenderInstallMode returns
// a non-empty string (basic smoke check for the render function).
func TestRenderInstallModeIsNonEmpty(t *testing.T) {
	out := RenderInstallMode(0)
	if out == "" {
		t.Fatalf("RenderInstallMode(0) returned empty string")
	}
}

// TestRenderInstallModeContainsBothOptions verifies that both menu options
// are present in the rendered output regardless of which cursor is active.
func TestRenderInstallModeContainsBothOptions(t *testing.T) {
	out := RenderInstallMode(0)
	if !strings.Contains(out, "Quick install") {
		t.Errorf("RenderInstallMode(0): missing 'Quick install' in output: %q", out)
	}
	if !strings.Contains(out, "Advanced install") {
		t.Errorf("RenderInstallMode(0): missing 'Advanced install' in output: %q", out)
	}
}

// TestRenderInstallModeCursorOne verifies cursor=1 (Advanced) also renders both options.
func TestRenderInstallModeCursorOne(t *testing.T) {
	out := RenderInstallMode(1)
	if !strings.Contains(out, "Quick install") {
		t.Errorf("RenderInstallMode(1): missing 'Quick install' in output: %q", out)
	}
	if !strings.Contains(out, "Advanced install") {
		t.Errorf("RenderInstallMode(1): missing 'Advanced install' in output: %q", out)
	}
}

// TestRenderInstallModeContainsHelpHint verifies the screen contains navigation hints.
func TestRenderInstallModeContainsHelpHint(t *testing.T) {
	out := RenderInstallMode(0)
	if !strings.Contains(out, "enter") {
		t.Errorf("RenderInstallMode: missing 'enter' in help hint: %q", out)
	}
	if !strings.Contains(out, "esc") {
		t.Errorf("RenderInstallMode: missing 'esc' in help hint: %q", out)
	}
}
