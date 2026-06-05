# Tasks: ciberbal-ai Menu and Install Flow

## Phase 1: Batch 1 — MVP-safe branding first

- [x] 1.1 Update MVP-facing branding in `internal/tui/styles/logo.go`, `internal/tui/styles/styles.go`, `internal/tui/screens/preset.go`, `internal/tui/screens/complete.go`, and `internal/app/app.go`; keep `PresetFullGentleman` and `go.mod` unchanged.
- [x] 1.2 Add or update focused assertions in existing TUI/app tests (for example `internal/tui/screens/complete_test.go` and related view tests) to lock the new `ciberbal-ai` copy without touching non-MVP screens.

## Phase 2: Batch 2 — Install-mode RED tests

- [x] 2.1 Extend `internal/tui/model_test.go` with failing navigation tests for `ScreenWelcome -> ScreenDetection -> ScreenInstallMode` and back-navigation from `ScreenInstallMode` to `ScreenDetection`.
- [x] 2.2 Add failing tests in `internal/tui/model_test.go` for Quick install: `QuickInstall=true`, default selection prefill, and direct jump to `ScreenDependencyTree`.
- [x] 2.3 Add failing tests in `internal/tui/model_test.go` for Advanced install: `QuickInstall=false` and forward navigation to `ScreenAgents`.
- [x] 2.4 Create `internal/tui/screens/install_mode_test.go` with a render test asserting both options (`Quick install`, `Advanced install`) and non-empty output.

## Phase 3: Batch 3 — Install-mode GREEN implementation

- [x] 3.1 Create `internal/tui/screens/install_mode.go` with `InstallModeOptions()` and `RenderInstallMode(cursor int)` following existing Bubbletea screen patterns.
- [x] 3.2 Update `internal/tui/model.go` to add `QuickInstall bool`, `ScreenInstallMode`, and the new `View`, `optionCount`, `confirmSelection`, and `goBack` branches.
- [x] 3.3 Update `internal/tui/router.go` so Detection continues into `ScreenInstallMode`, Advanced keeps the legacy path, and Quick routes to `ScreenDependencyTree` without Agents/Persona/Preset.
- [x] 3.4 Reuse existing selection/planning helpers in `internal/tui/model.go` to prefill Quick install defaults and call `buildDependencyPlan()`/`startInstalling()` unchanged.

## Phase 4: Batch 4 — Verification and optional binary entrypoint

- [x] 4.1 Run targeted verification for the touched area with `go test ./internal/tui/... ./internal/app/...`; fix regressions in welcome/install navigation before widening scope.
- [ ] 4.2 If Phases 1-4.1 stay MVP-safe, add `cmd/ciberbal-ai/main.go` delegating to `app.Run()` and keep `cmd/ciberbal-ai/main.go` intact; defer goreleaser/go.mod/deep rename work.
- [x] 4.3 Re-run `go test ./...` and `go vet ./...`, then update this checklist to mark completed items for the first `sdd-apply` batch.
