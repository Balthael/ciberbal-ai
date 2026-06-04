package app

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// writeMinimalSkill creates a minimal SKILL.md with valid frontmatter at path.
func writeMinimalSkill(t *testing.T, path, name, description string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("MkdirAll(%s): %v", filepath.Dir(path), err)
	}
	content := "---\nname: " + name + "\ndescription: " + description + "\n---\n\n## Hard Rules\n\n- Demo rule.\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("WriteFile(%s): %v", path, err)
	}
}

// TestSkillRegistryRefreshGeneratesRegistry verifies that
// `skill-registry refresh --cwd <tmp> --no-gitignore` writes
// .atl/skill-registry.md without requiring system detection.
func TestSkillRegistryRefreshGeneratesRegistry(t *testing.T) {
	tmp := t.TempDir()
	home := t.TempDir()

	// Set HOME so os.UserHomeDir() in runSkillRegistry resolves cleanly.
	origHome := os.Getenv("HOME")
	t.Cleanup(func() { os.Setenv("HOME", origHome) })
	os.Setenv("HOME", home)

	// Create a minimal skill so the registry is non-empty.
	writeMinimalSkill(t, filepath.Join(tmp, "skills", "demo", "SKILL.md"), "demo", "Demo skill for testing")

	var buf bytes.Buffer
	err := RunArgs([]string{"skill-registry", "refresh", "--cwd", tmp, "--no-gitignore"}, &buf)
	if err != nil {
		t.Fatalf("RunArgs(skill-registry refresh) error = %v", err)
	}

	// Registry file must exist.
	registryPath := filepath.Join(tmp, ".atl", "skill-registry.md")
	data, readErr := os.ReadFile(registryPath)
	if readErr != nil {
		t.Fatalf("registry file not written at %s: %v", registryPath, readErr)
	}

	// Content must reference the demo skill.
	if !strings.Contains(string(data), "demo") {
		t.Errorf("registry should contain 'demo'; got:\n%s", data)
	}

	// Output must mention refreshed.
	out := buf.String()
	if !strings.Contains(out, "Skill registry refreshed") {
		t.Errorf("stdout should confirm refresh; got: %q", out)
	}
}

// TestSkillRegistryRefreshQuietProducesNoOutput verifies that --quiet
// suppresses stdout but still writes the registry file.
func TestSkillRegistryRefreshQuietProducesNoOutput(t *testing.T) {
	tmp := t.TempDir()
	home := t.TempDir()

	origHome := os.Getenv("HOME")
	t.Cleanup(func() { os.Setenv("HOME", origHome) })
	os.Setenv("HOME", home)

	writeMinimalSkill(t, filepath.Join(tmp, "skills", "demo", "SKILL.md"), "demo", "Demo skill for testing")

	var buf bytes.Buffer
	err := RunArgs([]string{"skill-registry", "refresh", "--cwd", tmp, "--no-gitignore", "--quiet"}, &buf)
	if err != nil {
		t.Fatalf("RunArgs(skill-registry refresh --quiet) error = %v", err)
	}

	// --quiet must produce no output.
	if buf.Len() != 0 {
		t.Errorf("--quiet should produce no output; got: %q", buf.String())
	}

	// But the registry file must still exist.
	registryPath := filepath.Join(tmp, ".atl", "skill-registry.md")
	if _, statErr := os.Stat(registryPath); statErr != nil {
		t.Fatalf("registry file must exist even with --quiet: %v", statErr)
	}
}

// TestSkillRegistryRefreshShortQuietFlag verifies -q alias works identically to --quiet.
func TestSkillRegistryRefreshShortQuietFlag(t *testing.T) {
	tmp := t.TempDir()
	home := t.TempDir()

	origHome := os.Getenv("HOME")
	t.Cleanup(func() { os.Setenv("HOME", origHome) })
	os.Setenv("HOME", home)

	writeMinimalSkill(t, filepath.Join(tmp, "skills", "demo", "SKILL.md"), "demo", "Demo skill")

	var buf bytes.Buffer
	err := RunArgs([]string{"skill-registry", "refresh", "--cwd", tmp, "--no-gitignore", "-q"}, &buf)
	if err != nil {
		t.Fatalf("RunArgs(skill-registry refresh -q) error = %v", err)
	}
	if buf.Len() != 0 {
		t.Errorf("-q should produce no output; got: %q", buf.String())
	}
}

// TestSkillRegistryRefreshNoArgsReturnsUsageError verifies that
// `skill-registry` without subcommand returns a usage error.
func TestSkillRegistryRefreshNoArgsReturnsUsageError(t *testing.T) {
	var buf bytes.Buffer
	err := RunArgs([]string{"skill-registry"}, &buf)
	if err == nil {
		t.Fatal("skill-registry with no subcommand should return error")
	}
	if !strings.Contains(err.Error(), "ciberbal-ai skill-registry refresh") {
		t.Errorf("error should contain usage hint; got: %v", err)
	}
}

// TestSkillRegistryRefreshUnknownFlagReturnsError verifies that an unknown
// flag returns a proper error.
func TestSkillRegistryRefreshUnknownFlagReturnsError(t *testing.T) {
	tmp := t.TempDir()
	home := t.TempDir()

	origHome := os.Getenv("HOME")
	t.Cleanup(func() { os.Setenv("HOME", origHome) })
	os.Setenv("HOME", home)

	var buf bytes.Buffer
	err := RunArgs([]string{"skill-registry", "refresh", "--cwd", tmp, "--unknown-flag"}, &buf)
	if err == nil {
		t.Fatal("unknown flag should return error")
	}
}

// TestHelpContainsSkillRegistry verifies that help output lists skill-registry.
func TestHelpContainsSkillRegistry(t *testing.T) {
	var buf bytes.Buffer
	err := RunArgs([]string{"help"}, &buf)
	if err != nil {
		t.Fatalf("help should not fail: %v", err)
	}
	if !strings.Contains(buf.String(), "skill-registry") {
		t.Errorf("help output should contain 'skill-registry'; got:\n%s", buf.String())
	}
}

// TestSkillRegistryDispatchedBeforeSystemDetection verifies that skill-registry
// is dispatched without triggering system detection (no EnsureCurrentOSSupported call).
// This is a structural guard — if skill-registry were placed after system detection,
// it would fail on unsupported platforms.
func TestSkillRegistryDispatchedBeforeSystemDetection(t *testing.T) {
	tmp := t.TempDir()
	home := t.TempDir()

	origHome := os.Getenv("HOME")
	t.Cleanup(func() { os.Setenv("HOME", origHome) })
	os.Setenv("HOME", home)

	// No skill files — the registry will be written but empty.
	var buf bytes.Buffer
	// The command must not error due to platform detection; it runs cleanly.
	err := RunArgs([]string{"skill-registry", "refresh", "--cwd", tmp, "--no-gitignore"}, &buf)
	if err != nil {
		t.Fatalf("skill-registry must not require system detection: %v", err)
	}
}
