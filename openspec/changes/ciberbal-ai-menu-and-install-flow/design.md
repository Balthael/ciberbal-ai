# Design: ciberbal-ai Menu and Install Flow

## Technical Approach

Insert a single new screen (`ScreenInstallMode`) between `ScreenDetection` and `ScreenAgents` in the install flow. Quick install pre-populates `Selection` with full-cyber defaults and jumps directly to `ScreenDependencyTree → Review → Installing`, skipping Agents/Persona/Preset screens. Advanced install proceeds exactly as today. Branding changes are isolated to a handful of presentation files with zero logic changes.

## Architecture Decisions

| Decision | Choice | Alternatives | Rationale |
|----------|--------|-------------|-----------|
| Install-mode entry point | New `ScreenInstallMode` sub-screen after Detection | Split Welcome menu into two items | Avoids cursor index shifts and model_test regression; preserves optionCount() integrity |
| Quick install defaults | Pre-fill `Selection` + `QuickInstall bool` in Model; skip Agents/Persona/Preset | Separate pipeline path | Reuses existing `buildDependencyPlan` + `startInstalling` unchanged; minimal new code |
| Preset constant for Quick | Reuse `PresetFullGentleman` (display label renamed) | New `PresetFullCyber` constant | Avoids Tier 3 breakage until grep of test hardcodes is done; constant rename is a separate sub-task |
| goBack() from ScreenInstallMode | Return to `ScreenDetection` | Return to `ScreenWelcome` | Consistent with existing linear back-navigation in `goBack()` / `linearRoutes` |
| go.mod rename | Deferred (post-MVP) | Rename now | High risk: cascades to all imports, CI, goreleaser, 40+ files |
| Binary rename | `cmd/ciberbal-ai/` new dir with `main.go` delegating to `app.Run()` | Rename `cmd/gentle-ai/` | Additive; old binary still buildable; goreleaser targets added separately |

## Screen Graph (install flow only)

```
ScreenWelcome
    │ cursor=0 "Start installation"
    ▼
ScreenDetection
    │ cursor=0 "Continue"
    ▼
ScreenInstallMode        ← NEW
    ├─ cursor=0 "Quick install"
    │       │  populate Selection with full-cyber defaults
    │       │  set m.QuickInstall = true
    │       ▼
    │   ScreenDependencyTree ──► ScreenReview ──► ScreenInstalling ──► ScreenComplete
    │
    └─ cursor=1 "Advanced install"
            │  set m.QuickInstall = false
            ▼
        ScreenAgents ──► ScreenPersona ──► ScreenPreset ──► ... (unchanged)
```

## State Changes in Model

```go
// New field in tui.Model (additive)
QuickInstall bool  // true when Quick mode chosen; gates Agents/Persona/Preset skip

// New screen constant (additive, inserted between ScreenDetection and ScreenAgents)
ScreenInstallMode Screen
```

`NewModel()` initializes `QuickInstall = false` (no change to existing default path).

## Data Flow — Quick Install

```
ScreenInstallMode (Quick selected)
    │
    ├── m.Selection.Agents     = preselectedAgents(detection)   // all detected
    ├── m.Selection.Persona    = model.PersonaGentleman          // unchanged default
    ├── m.Selection.Preset     = model.PresetFullGentleman       // full preset
    ├── m.Selection.Components = componentsForPreset(PresetFullGentleman)
    ├── m.QuickInstall         = true
    └── m.buildDependencyPlan()
        └──► setScreen(ScreenDependencyTree)
```

No new planner/pipeline logic. The plan is built identically to Advanced; the only difference is skipping interactive screens.

## File Changes

| File | Action | Description |
|------|--------|-------------|
| `internal/tui/model.go` | Modify | Add `QuickInstall bool` to `Model`; add `ScreenInstallMode` constant; handle new screen in `confirmSelection`, `goBack`, `View`, `optionCount` |
| `internal/tui/router.go` | Modify | Add `ScreenInstallMode` route: `{Forward: ScreenAgents, Backward: ScreenDetection}`; also add transition to `ScreenDependencyTree` for quick path |
| `internal/tui/screens/install_mode.go` | Create | `RenderInstallMode(cursor int) string` — 2-option screen: Quick / Advanced |
| `internal/tui/styles/logo.go` | Modify | Replace braille ASCII art with ciberbal-ai logo |
| `internal/tui/styles/styles.go` | Modify | Update `Tagline()` → `"ciberbal-ai Stack " + version + " — One command. Any agent. Any OS."` |
| `internal/tui/screens/preset.go` | Modify | Update `presetDescriptions[PresetFullGentleman]` → cybersecurity-framed description |
| `internal/tui/screens/complete.go` | Modify | Replace `"Run gentle-ai again to retry"` → `"Run ciberbal-ai again to retry"` |
| `internal/app/app.go` | Modify | Update version print `"gentle-ai %s"` → `"ciberbal-ai %s"`; `backupRoot` path `".gentle-ai"` → `".ciberbal-ai"` |
| `cmd/ciberbal-ai/main.go` | Create | New entry point delegating to `app.Run()` |

**Unchanged**: all backup/restore/profile/agent-builder/sync/upgrade/model-config screens, packages, and tests.

## Interfaces / Contracts

```go
// install_mode.go
package screens

func RenderInstallMode(cursor int) string
func InstallModeOptions() []string  // ["Quick install", "Advanced install"]
```

`confirmSelection` in `model.go` for `ScreenInstallMode`:

```go
case ScreenInstallMode:
    switch m.Cursor {
    case 0: // Quick
        m.QuickInstall = true
        m.Selection.Agents = preselectedAgents(m.Detection)
        m.Selection.Persona = model.PersonaGentleman
        m.Selection.Preset = model.PresetFullGentleman
        m.Selection.Components = componentsForPreset(model.PresetFullGentleman)
        m.buildDependencyPlan()
        m.setScreen(ScreenDependencyTree)
    case 1: // Advanced
        m.QuickInstall = false
        m.setScreen(ScreenAgents)
    }
```

## Testing Strategy

| Layer | What to Test | Approach |
|-------|-------------|----------|
| Unit | `ScreenWelcome → ScreenDetection → ScreenInstallMode` transition | Extend `TestNavigationWelcomeToDetection`; add `TestNavigationDetectionToInstallMode` |
| Unit | Quick install → `ScreenDependencyTree` + `QuickInstall=true` + `Selection` pre-filled | `TestQuickInstallPopulatesSelectionAndJumps` |
| Unit | Advanced install → `ScreenAgents` + `QuickInstall=false` | `TestAdvancedInstallGoesToAgents` |
| Unit | Esc from `ScreenInstallMode` → `ScreenDetection` | `TestInstallModeBackNavigation` |
| Unit | `RenderInstallMode` returns non-empty string with both options | `TestRenderInstallMode` in `screens` package |
| Regression | All existing Welcome cursor navigation tests pass unchanged | Re-run `TestNavigationWelcomeToDetection` and sibling tests |
| Regression | `optionCount()` for `ScreenInstallMode` returns 2 | Add to `optionCount` coverage |

Tests follow existing pattern in `model_test.go`: construct `Model`, send `tea.KeyMsg`, assert `state.Screen` and fields.

## Migration / Rollout

No data migration required. `QuickInstall` is an in-session boolean, not persisted. The backup path change (`~/.gentle-ai` → `~/.ciberbal-ai`) affects new backups only; existing backups remain readable via their absolute paths in manifests.

## Open Questions

- [ ] Should `preselectedAgents(detection)` in Quick mode include ALL detected agents or only a curated cybersec subset? (Current design: all detected — safe default, matches existing `NewModel` behavior)
- [ ] Is the backup dir rename (`~/.gentle-ai` → `~/.ciberbal-ai`) in scope for MVP or deferred with go.mod rename? (Recommendation: defer both together to avoid partial identity split)
