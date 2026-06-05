package tui

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gentleman-programming/ciberbal-ai/internal/model"
	"github.com/gentleman-programming/ciberbal-ai/internal/system"
)

// ─── Task 2.1: Navigation tests for ScreenInstallMode ─────────────────────────

// TestNavigationDetectionToInstallMode verifies that pressing Enter on
// ScreenDetection (cursor=0 "Continue") transitions to ScreenInstallMode,
// NOT directly to ScreenAgents as before.
func TestNavigationDetectionToInstallMode(t *testing.T) {
	m := NewModel(system.DetectionResult{}, "dev")
	m.Screen = ScreenDetection
	m.Cursor = 0

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	state := updated.(Model)

	if state.Screen != ScreenInstallMode {
		t.Fatalf("screen = %v, want ScreenInstallMode (%v)", state.Screen, ScreenInstallMode)
	}
}

// TestInstallModeBackNavigationToDetection verifies that pressing Esc from
// ScreenInstallMode returns to ScreenDetection (consistent linear back-nav).
func TestInstallModeBackNavigationToDetection(t *testing.T) {
	m := NewModel(system.DetectionResult{}, "dev")
	m.Screen = ScreenInstallMode
	m.Cursor = 0

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEsc})
	state := updated.(Model)

	if state.Screen != ScreenDetection {
		t.Fatalf("Esc from ScreenInstallMode: screen = %v, want ScreenDetection (%v)", state.Screen, ScreenDetection)
	}
}

// ─── Task 2.2: Quick install tests ───────────────────────────────────────────

// TestQuickInstallSetsQuickInstallFlag verifies that selecting Quick install
// (cursor=0) from ScreenInstallMode sets QuickInstall=true on the model.
func TestQuickInstallSetsQuickInstallFlag(t *testing.T) {
	m := NewModel(system.DetectionResult{}, "dev")
	m.Screen = ScreenInstallMode
	m.Cursor = 0 // Quick install

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	state := updated.(Model)

	if !state.QuickInstall {
		t.Fatalf("QuickInstall = false after selecting Quick install; want true")
	}
}

// TestQuickInstallJumpsToDependencyTree verifies that Quick install skips
// Agents/Persona/Preset screens and goes directly to ScreenDependencyTree.
func TestQuickInstallJumpsToDependencyTree(t *testing.T) {
	m := NewModel(system.DetectionResult{}, "dev")
	m.Screen = ScreenInstallMode
	m.Cursor = 0 // Quick install

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	state := updated.(Model)

	if state.Screen != ScreenDependencyTree {
		t.Fatalf("screen after Quick install = %v, want ScreenDependencyTree (%v)", state.Screen, ScreenDependencyTree)
	}
}

// TestQuickInstallPreSelectsAgents verifies that Quick install pre-populates
// Selection.Agents with detected (or default) agents.
func TestQuickInstallPreSelectsAgents(t *testing.T) {
	m := NewModel(system.DetectionResult{}, "dev")
	m.Screen = ScreenInstallMode
	m.Cursor = 0 // Quick install

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	state := updated.(Model)

	if len(state.Selection.Agents) == 0 {
		t.Fatalf("Quick install: Selection.Agents is empty; want at least one pre-selected agent")
	}
}

// ─── Task 2.3: Advanced install tests ────────────────────────────────────────

// TestAdvancedInstallKeepsQuickInstallFalse verifies that selecting Advanced
// install (cursor=1) does NOT set QuickInstall=true.
func TestAdvancedInstallKeepsQuickInstallFalse(t *testing.T) {
	m := NewModel(system.DetectionResult{}, "dev")
	m.Screen = ScreenInstallMode
	m.Cursor = 1 // Advanced install

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	state := updated.(Model)

	if state.QuickInstall {
		t.Fatalf("QuickInstall = true after selecting Advanced install; want false")
	}
}

// TestAdvancedInstallGoesToAgents verifies that selecting Advanced install
// (cursor=1) proceeds to ScreenAgents (the granular flow).
func TestAdvancedInstallGoesToAgents(t *testing.T) {
	m := NewModel(system.DetectionResult{}, "dev")
	m.Screen = ScreenInstallMode
	m.Cursor = 1 // Advanced install

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	state := updated.(Model)

	if state.Screen != ScreenAgents {
		t.Fatalf("screen after Advanced install = %v, want ScreenAgents (%v)", state.Screen, ScreenAgents)
	}
}

// TestQuickInstallSetsFullPreset verifies that Quick install sets the preset to
// PresetFullPentest (the full cybersecurity defaults).
func TestQuickInstallSetsFullPreset(t *testing.T) {
	m := NewModel(system.DetectionResult{}, "dev")
	m.Screen = ScreenInstallMode
	m.Cursor = 0 // Quick install

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	state := updated.(Model)

	// The preset must be set to the full preset (full-pentest = full cyber defaults).
	if state.Selection.Preset != model.PresetFullPentest {
		t.Fatalf("Selection.Preset = %q after Quick install, want %q", state.Selection.Preset, model.PresetFullPentest)
	}
}

// ─── optionCount regression ───────────────────────────────────────────────────

// TestInstallModeOptionCount verifies that ScreenInstallMode has exactly 2
// options (Quick install + Advanced install). This guards against cursor
// wrap-around bugs (the critical optionCount guard from the design).
func TestInstallModeOptionCount(t *testing.T) {
	m := NewModel(system.DetectionResult{}, "dev")
	m.Screen = ScreenInstallMode

	count := m.optionCount()
	if count != 2 {
		t.Fatalf("optionCount() for ScreenInstallMode = %d, want 2", count)
	}
}
