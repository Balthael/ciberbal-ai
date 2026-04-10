# Exploration: ciberbal-ai Menu and Install Flow

## Current State

gentle-ai is a Go/Bubbletea CLI/TUI that configures AI agent ecosystems.
Its welcome menu and install wizard are spread across these key files:

| File | Role |
|------|------|
| `internal/tui/screens/welcome.go` | Menu labels + RenderWelcome |
| `internal/tui/styles/logo.go` | Braille ASCII art logo (17 lines) |
| `internal/tui/styles/styles.go` | Tagline string, color palette (Rose Pine) |
| `internal/tui/screens/preset.go` | Preset options + descriptions |
| `internal/tui/screens/persona.go` | Persona options |
| `internal/tui/screens/agents.go` | Agent checkboxes |
| `internal/tui/screens/skill_picker.go` | Skill groups + labels |
| `internal/tui/screens/sdd_mode.go` | SDD mode descriptions |
| `internal/tui/screens/complete.go` | Post-install "next steps" text |
| `internal/tui/model.go` | Screen constants, Model struct, NewModel defaults, componentsForPreset, preselectedAgents |
| `internal/tui/router.go` | Linear route map between all screens |
| `internal/catalog/agents.go` | AllAgents() list |
| `internal/catalog/components.go` | MVPComponents() list |
| `internal/catalog/skills.go` | MVPSkills() list |
| `internal/model/types.go` | Preset, Persona, Component, Skill, Agent ID constants |
| `internal/app/app.go` | app name strings (version printout, error msg, backup path) |
| `cmd/gentle-ai/main.go` | Binary entrypoint (no product name hardcoded except import path) |

### Full Menu Tree (gentle-ai today)

**Welcome screen options:**
```
0  Start installation
1  Upgrade tools [★ | (up to date) badge]
2  Sync configs
3  Upgrade + Sync
4  Configure models
5  Create your own Agent [(no agents) if no engines]
6  OpenCode SDD Profiles (N)  [conditional: only if OpenCode detected]
7/6  Manage backups
8/7  Quit
```

**Install wizard screens (in order):**
```
Welcome → Detection → Agents (checkboxes) → Persona (radio) → Preset (radio)
  → [Claude+SDD] ClaudeModelPicker
  → [OpenCode+SDD] SDDMode → [multi+cache] ModelPicker → StrictTDD → DependencyTree
  → [Custom preset] DependencyTree(picker) → [...conditionals...] → SkillPicker
  → Review → Installing → Complete
```

**Preset options (preset.go):**
- `full-gentleman` — Everything: memory, SDD, skills, docs, persona & security
- `ecosystem-only` — Core tools only: memory, SDD, skills & docs
- `minimal` — Just Engram persistent memory
- `custom` — Pick individual components yourself

**Persona options:**
- `gentleman`, `neutral`, `custom`

**Components (catalog/components.go):**
- engram, sdd, skills, context7, persona, permissions, gga, theme

**Skills (catalog/skills.go + skill_picker.go):**
- SDD group: sdd-init, sdd-explore, sdd-propose, sdd-spec, sdd-design, sdd-tasks, sdd-apply, sdd-verify, sdd-archive, sdd-onboard, judgment-day
- Foundation group: go-testing, skill-creator, branch-pr, issue-creation, skill-registry

**SDD mode options:**
- `single` — one orchestrator handles all SDD phases
- `multi` — dedicated sub-agent per SDD phase

---

## Affected Areas

### Branding / Identity (MUST change)
- `internal/tui/styles/logo.go` — replace braille ASCII art with ciberbal-ai logo
- `internal/tui/styles/styles.go` — update `Tagline()` string; optionally change color palette to cybersecurity-themed
- `internal/app/app.go` — product name in version print, error messages, backup path (`~/.gentle-ai/backups` → `~/.ciberbal-ai/backups`)
- `cmd/gentle-ai/main.go` → rename to `cmd/ciberbal-ai/main.go`
- `go.mod` — module path: `github.com/gentleman-programming/gentle-ai` → `github.com/gentleman-programming/ciberbal-ai` (or keep as-is for MVP)
- `complete.go` — post-install next steps copy text references "gentle-ai" in the retry instruction

### Install Wizard Screens (MUST change for ciberbal-ai semantics)
- `internal/tui/screens/preset.go` — rename presets + descriptions for cybersecurity context
- `internal/tui/screens/persona.go` — rename personas for cybersecurity context
- `internal/tui/screens/complete.go` — update next steps copy text

### Install Wizard Screens (NO CHANGE needed in content — preserved 1:1)
- `internal/tui/screens/agents.go` — unchanged (agent list stays the same)
- `internal/tui/screens/detection.go` — unchanged
- `internal/tui/screens/sdd_mode.go` — unchanged (SDD mode options are domain-agnostic)
- `internal/tui/screens/skill_picker.go` — unchanged (SDD skills are product-agnostic)
- `internal/tui/screens/dependency_tree.go` — unchanged
- `internal/tui/screens/review.go` — minimal if any changes
- `internal/tui/screens/installing.go` — unchanged
- All backup/restore screens — unchanged
- All upgrade/sync screens — unchanged
- All agent builder screens — unchanged
- All profile screens — unchanged
- `internal/tui/router.go` — unchanged
- `internal/tui/model.go` — logic preserved; only NewModel defaults and componentsForPreset need adjusting

### Core Infrastructure (PRESERVE 100%)
- `internal/catalog/agents.go` — all 8 agents kept as-is
- `internal/catalog/components.go` — all 8 components kept as-is
- `internal/catalog/skills.go` — all skills kept as-is
- `internal/model/` — all types/IDs preserved
- `internal/pipeline/`, `internal/planner/`, `internal/installcmd/` — zero changes
- `internal/backup/`, `internal/state/`, `internal/verify/` — zero changes
- `internal/agents/`, `internal/opencode/`, `internal/agentbuilder/` — zero changes
- `internal/components/`, `internal/assets/` — zero changes
- `internal/system/`, `internal/update/` — zero changes
- `internal/cli/`, `internal/app/app.go` (logic) — zero changes to logic; only product-name strings change

---

## Quick Install vs Advanced Install Design

### The Core User Intent
> Quick install = install EVERYTHING for cybersecurity/pentesting, no specialty selection needed.
> Advanced install = preserve granular configuration (current install wizard as-is).

This maps cleanly to the existing preset architecture:

| Intent | gentle-ai today | ciberbal-ai mapping |
|--------|----------------|---------------------|
| Quick install | `full-gentleman` preset auto-selected | `full-cyber` preset auto-selected, skip Preset screen |
| Advanced install | Full wizard (agents → persona → preset → SDD → review) | Full wizard preserved 1:1 |

### Quick Install Flow (proposed)
```
Welcome → Detection → [auto: select all detected/available agents]
        → [auto: preset = full-cyber] → [auto: SDDMode = single] → [auto: StrictTDD = false]
        → Review → Installing → Complete
```

The **only screen the user sees** in quick install mode is:
1. Welcome (choose quick vs advanced)
2. Detection (shows what was found)
3. Review (summary before install)
4. Installing + Complete

This requires a new screen entry point or a **mode flag on Detection/Welcome** that bypasses Agents/Persona/Preset screens and auto-populates `model.Selection` with the full-cyber defaults.

**Implementation approach (lowest risk):** Add a `QuickInstall bool` field to `tui.Model`. When `QuickInstall = true`, after Detection the model sets all defaults and jumps directly to `ScreenDependencyTree` (auto-built plan) → `ScreenReview`.

### Advanced Install Flow (preserved)
Identical to current gentle-ai wizard — no changes to any screen. The welcome menu just routes to `ScreenDetection` with `QuickInstall = false`.

### Welcome Screen Redesign for ciberbal-ai

Replace current option 0 "Start installation" with two explicit options:

```
0  Quick install (everything for pentesting)
1  Advanced install (choose your setup)
2  Upgrade tools [★ badge]
3  Sync configs
4  Upgrade + Sync
5  Configure models
6  Create your own Agent
7  OpenCode SDD Profiles (N)  [conditional]
8  Manage backups
9  Quit
```

This shifts indices for existing menu routing by +1 after option 0. The simplest MVP alternative is to **keep the single "Start installation" option** and gate Quick vs Advanced at the **Detection screen level** with a mode toggle before continuing.

**Recommendation:** Add a new `ScreenInstallMode` screen between Detection and Agents that asks "Quick install / Advanced install". This avoids shifting welcome menu indices and is a single focused screen.

---

## Preset Renaming for ciberbal-ai

| gentle-ai PresetID | ciberbal-ai PresetID | Description |
|-------------------|---------------------|-------------|
| `full-gentleman` | `full-cyber` | Everything: memory, SDD, skills, docs, persona & security |
| `ecosystem-only` | `ecosystem-only` (unchanged) | Core tools only: no persona/security |
| `minimal` | `minimal` (unchanged) | Just Engram persistent memory |
| `custom` | `custom` (unchanged) | Pick individual components yourself |

> Note: PresetIDs are string constants. Renaming `full-gentleman` → `full-cyber` is a 2-file change (model/types.go + tui/screens/preset.go) but requires verifying no other package uses the string value "full-gentleman" directly.

---

## Persona Renaming for ciberbal-ai

| gentle-ai PersonaID | ciberbal-ai PersonaID | Meaning |
|--------------------|----------------------|---------|
| `gentleman` | `hacker` (or keep `gentleman`) | Default cybersecurity persona |
| `neutral` | `neutral` (unchanged) | No persona injection |
| `custom` | `custom` (unchanged) | User-defined |

> Risk: renaming `gentleman` → `hacker` touches persona injection templates in `internal/assets/` and possibly `internal/agents/`. MVP alternative: keep persona IDs but update the display label only.

---

## Implementation Scope (MVP)

### Tier 1 — Pure branding (zero logic change, safe, fast)
1. Replace logo braille art in `styles/logo.go`
2. Update `Tagline()` in `styles/styles.go` to: `"ciberbal-ai Stack {version} — One command. Any agent. Pentest-ready."`
3. Update product name strings in `app/app.go` (version print, backup path prefix)
4. Rename binary entrypoint: `cmd/gentle-ai/` → `cmd/ciberbal-ai/`
5. Update preset descriptions in `screens/preset.go` for cybersecurity framing
6. Update post-install next steps copy in `screens/complete.go`
7. Update `go.mod` module path (deferred or MVP — low risk but touches all import paths)

### Tier 2 — Quick install path (low risk, focused logic addition)
8. Add `ScreenInstallMode` screen constant (new screen between Detection and Agents)
9. Add `QuickInstall bool` to `tui.Model`
10. Add `RenderInstallMode()` in `screens/` (two options: Quick / Advanced)
11. In `model.go` `confirmSelection()`: when ScreenInstallMode + Quick selected, populate `Selection` with full-cyber defaults and jump to `ScreenDependencyTree` (skip Agents/Persona/Preset)
12. Update `router.go` to add `ScreenInstallMode` entry
13. Update `optionCount()` in model.go to handle new screen

### Tier 3 — Preset rename (medium risk, string constants)
14. Add `PresetFullCyber PresetID = "full-cyber"` to `model/types.go`
15. Update `componentsForPreset()` in `model.go`
16. Update `screens/preset.go` to use new constant
17. Verify no hardcoded "full-gentleman" string outside these files

### Tier 4 — Deferred (post-MVP)
- Color palette update (Rose Pine → dark cyberpunk palette)
- Persona rename (gentleman → hacker)
- go.mod module path rename (requires global import path replace)
- Adding cybersecurity-specific skills to skill catalog
- Localized ciberbal-ai AGENTS.md / README

---

## Risks and Edge Cases

| Risk | Severity | Mitigation |
|------|----------|------------|
| `go.mod` module rename breaks all import paths | High | Defer to post-MVP; run with original path or use a batch `sed` |
| Welcome cursor index shift if adding Quick/Advanced as separate menu items | Medium | Use `ScreenInstallMode` as a sub-screen instead of splitting welcome items |
| `PresetFullCyber` string value used in persisted state files (state.json, backup manifests) | Medium | Detect and migrate on load, or keep `full-gentleman` as internal ID and only change display label |
| Persona template files reference "gentleman" persona ID string | Low-Medium | Check `internal/assets/` and `internal/agents/` for hardcoded PersonaGentleman string |
| Tests that hardcode "full-gentleman" preset ID string | Medium | grep test files; update any literal string comparisons |
| Quick install auto-selecting all agents on a machine with many agents | Low | Use detected agents (preselectedAgents logic) same as today — already handles this |
| TUI tests for welcome screen that count menu options | Low | Update `optionCount()` and welcome screen tests if adding ScreenInstallMode |
| Backup directory path migration (~/.gentle-ai → ~/.ciberbal-ai) | Low | For existing users; add migration notice on first run if old dir exists |

---

## Recommendation

**Approach: Incremental branding + ScreenInstallMode injection**

Do NOT split "Start installation" into two welcome menu items (avoids index shifting and test breakage). Instead:

1. Run Tier 1 (pure branding) as the first implementation batch — zero logic change, fully testable, self-contained.
2. Run Tier 2 (ScreenInstallMode) as the second batch — single new screen, narrow `model.go` changes, clear test surface.
3. Defer Tier 3 (preset rename) until Tier 1+2 are stable — it requires verifying all string usages.
4. Defer Tier 4 entirely.

This gives a working `ciberbal-ai` binary with correct branding and quick-install path in two focused implementation sessions.

### Ready for Proposal
**Yes** — requirements are clear enough to propose. The Proposal should commit to:
- Binary name: `ciberbal-ai`
- Quick install preset contents: all 8 components (full-cyber = full-gentleman component set)
- Install mode screen placement: after Detection, before Agents
- Module path: defer rename to post-MVP
