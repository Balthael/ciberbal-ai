<div align="center">

<img width="3276" height="1280" alt="image" src="https://github.com/user-attachments/assets/3a3e4ae1-b9f4-4ce9-8fd0-3833812beb99" />

<h1>ciberbal-ai</h1>

<p><strong>One command. Any agent. Any OS. A cybersecurity and pentesting AI ecosystem -- configured and ready.</strong></p>

<p>
<a href="https://github.com/Gentleman-Programming/gentle-ai/releases"><img src="https://img.shields.io/github/v/release/Gentleman-Programming/gentle-ai" alt="Release"></a>
<a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License: MIT"></a>
<img src="https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go&logoColor=white" alt="Go 1.24+">
<img src="https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-lightgrey" alt="Platform">
</p>

</div>

---

## What It Does

This is NOT just an AI agent installer. Most agents are easy to install already. `ciberbal-ai` is an **ecosystem configurator** for cybersecurity and pentesting workflows -- it takes the AI agent(s) you already use and upgrades them with persistent memory, Spec-Driven Development workflow, curated skills, MCP servers, model routing, and a security-first operating style.

**Before**: "I installed Claude Code / OpenCode / Cursor, but it's just a chatbot."

**After**: Your agent now has memory, methodology, skills, workflow, MCP tools, and a cybersecurity-oriented setup that is ready for real pentesting work.

### 8 Supported Agents

| Agent | Delegation Model | Key Feature |
|-------|:---:|---|
| **Claude Code** | Full (Task tool) | Sub-agents, output styles |
| **OpenCode** | Full (multi-mode overlay) | Per-phase model routing |
| **Gemini CLI** | Full (experimental) | Custom agents in `~/.gemini/agents/` |
| **Cursor** | Full (native subagents) | 9 SDD agents in `~/.cursor/agents/` |
| **VS Code Copilot** | Full (runSubagent) | Parallel execution |
| **Codex** | Solo-agent | CLI-native, TOML config |
| **Windsurf** | Solo-agent | Plan Mode, Code Mode, native workflows |
| **Antigravity** | Solo-agent + Mission Control | Built-in Browser/Terminal sub-agents |

> **Note**: This working copy is a derivative adaptation of `gentle-ai` focused on cybersecurity and pentesting. Some package, release, and install paths still reference the original project until the full public rename is completed.

---

## Quick Start

> **Current state**: this local copy is ready to test, but some install commands below still reference original `gentle-ai` distribution paths because the Go module and release infrastructure have not been fully renamed yet.

### macOS / Linux

```bash
curl -fsSL https://raw.githubusercontent.com/Gentleman-Programming/gentle-ai/main/scripts/install.sh | bash
```

### Windows

```powershell
scoop bucket add gentleman https://github.com/Gentleman-Programming/scoop-bucket
scoop install gentle-ai
```

Or via PowerShell script: `irm https://raw.githubusercontent.com/Gentleman-Programming/gentle-ai/main/scripts/install.ps1 | iex`

### After install: project-level setup

Once your agents are configured, open your AI agent in a project and run these two commands to register the project context:

| Command | What it does | When to re-run |
|---------|-------------|----------------|
| `/sdd-init` | Detects stack, testing capabilities, activates Strict TDD Mode if available | When your project adds/removes test frameworks, or first time in a new project |
| `skill-registry` | Scans installed skills and project conventions, builds the registry | After installing/removing skills, or first time in a new project |

These are **not required** for basic usage. The SDD orchestrator runs `/sdd-init` automatically if it detects no context. But if something changed in your project (new test runner, new dependencies), re-running them manually ensures the agents have up-to-date context.

### Default operating model

`ciberbal-ai` is designed to support a **fast full-stack install** for cybersecurity and pentesting:

- Quick install configures the full ecosystem by default
- no specialty prompt is required during the quick path
- Advanced install preserves granular configuration when you need it
- the long-term target is one setup usable across web, API, mobile, cloud, AD/internal, reporting, and general security workflows

---

## Install

> **Important**: until the public rename is completed, command examples may still use `gentle-ai` names, package paths, or release locations. For local testing of this working copy, prefer building or running from source.

### Local development / testing

```bash
go test ./...
go run ./cmd/ciberbal-ai
```

### Homebrew (macOS / Linux)

```bash
brew tap Gentleman-Programming/homebrew-tap
brew install gentle-ai
```

### Go install (any platform with Go 1.24+)

```bash
go install github.com/gentleman-programming/gentle-ai/cmd/gentle-ai@latest
```

### Scoop (Windows)

```powershell
scoop bucket add gentleman https://github.com/Gentleman-Programming/scoop-bucket
scoop install gentle-ai
```

**Migrating from PowerShell installer to Scoop?** Remove the old binary first:

```powershell
Remove-Item "$env:LOCALAPPDATA\gentle-ai" -Recurse -Force
# Then install via Scoop as shown above
```

### Windows (PowerShell — alternative)

```powershell
# Option 1: PowerShell installer (downloads binary from GitHub Releases)
irm https://raw.githubusercontent.com/Gentleman-Programming/gentle-ai/main/scripts/install.ps1 | iex

# Option 2: Go install (requires Go 1.24+)
go install github.com/gentleman-programming/gentle-ai/cmd/gentle-ai@latest
```

### From releases

Download the binary for your platform from [GitHub Releases](https://github.com/Gentleman-Programming/gentle-ai/releases).

---

## Backups

Every install, sync, and upgrade automatically snapshots your config files. Backups are **compressed** (tar.gz), **deduplicated** (identical configs are not re-backed up), and **auto-pruned** (keeps the 5 most recent). Pin important backups via the TUI (`p` key) to protect them from pruning.

> **Current note**: backup storage paths are still inherited from the original project in some internal code paths and will be renamed in a later cleanup pass.

See [Backup & Rollback Guide](docs/rollback.md) for details.

---

## Documentation

| Topic | Description |
|-------|-------------|
| [Intended Usage](docs/intended-usage.md) | How ciberbal-ai is meant to be used — the mental model |
| [Agents](docs/agents.md) | Supported agents, feature matrix, config paths, and per-agent notes |
| [Components, Skills & Presets](docs/components.md) | All components, GGA behavior, skill catalog, and preset definitions |
| [Usage](docs/usage.md) | Persona modes, interactive TUI, CLI flags, and dependency management |
| [Backup & Rollback](docs/rollback.md) | Backup retention, compression, dedup, pinning, and restore |
| [Platforms](docs/platforms.md) | Supported platforms, Windows notes, security verification, config paths |
| [Architecture & Development](docs/architecture.md) | Codebase layout, testing, and relationship to Gentleman.Dots |

---

## Contributors

This project exists because of the community. See [CONTRIBUTORS.md](CONTRIBUTORS.md) for the full list.

<a href="https://github.com/Gentleman-Programming/gentle-ai/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=Gentleman-Programming/gentle-ai" />
</a>

---

## Credits

`ciberbal-ai` is based on [`gentle-ai`](https://github.com/Gentleman-Programming/gentle-ai), originally created by **Gentleman Programming**.

This project adapts that foundation for cybersecurity and pentesting workflows, preserving the original MIT license while extending the platform with a security-focused install flow, branding, and future domain-specific skills.

## License and Attribution

This repository includes work derived from `gentle-ai`, Copyright (c) 2025 Gentleman Programming, used under the MIT License.

The original copyright notice and permission notice are preserved in [`LICENSE`](LICENSE), as required by the license terms.

---

<div align="center">
<a href="LICENSE"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License: MIT"></a>
</div>
