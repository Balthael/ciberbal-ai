# Usage

← [Back to README](../README.md)

---

## Persona Modes

| Persona | ID | Description |
|---------|-----|-------------|
| Gentleman | `gentleman` | Teaching-oriented mentor persona — pushes back on bad practices, explains the why |
| Neutral | `neutral` | Same teacher, same philosophy, no regional language — warm and professional |
| Custom | `custom` | Bring your own persona instructions |

---

## Interactive TUI

Just run it — the Bubbletea TUI guides you through agent selection, components, skills, and presets. The top-level menu stays the same, and `Start installation` remains a single path that provisions the full pentesting ecosystem by default:

```bash
ciberbal-ai
```

---

## CLI Commands

### install

First-time setup — detects your tools, configures agents, injects all components:

```bash
# Full ecosystem for multiple agents
ciberbal-ai install \
  --agent claude-code,opencode,gemini-cli \
  --preset full-pentest

# Minimal setup for Cursor
ciberbal-ai install \
  --agent cursor \
  --preset minimal

# Pick specific components and skills
ciberbal-ai install \
  --agent claude-code \
  --component engram,sdd,skills,context7,persona,permissions \
  --skill go-testing,skill-creator,branch-pr,issue-creation \
  --persona gentleman

# Dry-run first (preview plan without applying changes)
ciberbal-ai install --dry-run \
  --agent claude-code,opencode \
  --preset full-pentest
```

### sync

Refresh managed assets to the current version after upgrading `ciberbal-ai` or when you want your local configs aligned with the latest release. Does NOT reinstall binaries (engram, GGA) — only updates prompt content, skills, MCP configs, and SDD orchestrators.

```bash
# Sync all installed agents
ciberbal-ai sync

# Sync specific agents only
ciberbal-ai sync --agent cursor --agent windsurf

# Sync a specific component
ciberbal-ai sync --component sdd
ciberbal-ai sync --component skills
ciberbal-ai sync --component engram
```

Sync is safe and idempotent — running it twice produces no changes the second time.

### update / upgrade

Check for and install new versions of `ciberbal-ai` itself:

```bash
# Check if a newer version is available
ciberbal-ai update

# Upgrade to the latest release (downloads new binary, replaces current)
ciberbal-ai upgrade
```

After upgrading, run `ciberbal-ai sync` to refresh all managed assets to the new version's content.

### version

```bash
ciberbal-ai version
ciberbal-ai --version
ciberbal-ai -v
```

---

## CLI Flags (install)

| Flag | Description |
|------|-------------|
| `--agent`, `--agents` | Agents to configure (comma-separated) |
| `--component`, `--components` | Components to install (comma-separated) |
| `--skill`, `--skills` | Skills to install (comma-separated) |
| `--persona` | Persona mode: `gentleman`, `neutral`, `custom` |
| `--preset` | Preset: `full-pentest`, `ecosystem-core`, `minimal`, `custom` |
| `--dry-run` | Preview the install plan without applying changes |

## CLI Flags (sync)

| Flag | Description |
|------|-------------|
| `--agent`, `--agents` | Agents to sync (defaults to all installed agents) |
| `--component` | Sync a specific component only: `sdd`, `engram`, `context7`, `skills`, `gga`, `permissions`, `theme` |
| `--include-permissions` | Include permissions sync (opt-in) |
| `--include-theme` | Include theme sync (opt-in) |

---

## Typical Workflow

```bash
# First time: install everything
ciberbal-ai install --agent claude-code,cursor --preset full-pentest

# After a new release: upgrade + sync
ciberbal-ai upgrade
ciberbal-ai sync

# Adding a new agent later
ciberbal-ai install --agent windsurf --preset full-pentest
```

---

## Dependency Management

`ciberbal-ai` auto-detects prerequisites before installation and provides platform-specific guidance. It does **not** ask you to choose AD, web, API, mobile, or wireless specialties during install — those domain layers are provisioned together and used after setup.

- **Detected tools**: git, curl, node, npm, brew, go
- **Version checks**: validates minimum versions where applicable
- **Platform-aware hints**: suggests `brew install`, `apt install`, `pacman -S`, `dnf install`, or `winget install` depending on your OS
- **Node LTS alignment**: on apt/dnf systems, Node.js hints use NodeSource LTS bootstrap before package install
- **Dependency-first approach**: detects what's installed, calculates what's needed, shows the full dependency tree before installing anything, then verifies each dependency after installation
