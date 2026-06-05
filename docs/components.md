# Components, Skills & Presets

← [Back to README](../README.md)

---

## Components

| Component | ID | Description |
|-----------|-----|-------------|
| Engram | `engram` | Persistent cross-session memory via MCP — auto-detection of project name, full-text search, git sync, project consolidation. See the upstream [Engram repo](https://github.com/Gentleman-Programming/engram) |
| SDD | `sdd` | Spec-Driven Development workflow (9 phases) — the agent handles SDD organically when the task warrants it, or when you ask; you don't need to learn the commands |
| Skills | `skills` | Curated capability-layer skill library spanning pentesting domains and delivery workflows |
| Context7 | `context7` | MCP server for live framework/library documentation |
| Persona | `persona` | Gentleman, neutral, or custom behavior mode |
| Permissions | `permissions` | Security-first defaults and guardrails |
| GGA | `gga` | Gentleman Guardian Angel — AI provider switcher |
| Theme | `theme` | Gentleman Kanagawa theme overlay |

## GGA Behavior

`ciberbal-ai --component gga` installs/provisions the `gga` binary globally on your machine.

It does **not** run project-level hook setup automatically (`gga init` / `gga install`) because that should be an explicit decision per repository.

After global install, enable GGA per project with:

```bash
gga init
gga install
```

---

## Skills

### Included Skills (installed by ciberbal-ai)

14 skill files organized by category, embedded in the binary and injected into your agent's configuration:

#### SDD (Spec-Driven Development)

| Skill | ID | Description |
|-------|-----|-------------|
| SDD Init | `sdd-init` | Bootstrap SDD context in a project |
| SDD Explore | `sdd-explore` | Investigate codebase before committing to a change |
| SDD Propose | `sdd-propose` | Create change proposal with intent, scope, approach |
| SDD Spec | `sdd-spec` | Write specifications with requirements and scenarios |
| SDD Design | `sdd-design` | Technical design with architecture decisions |
| SDD Tasks | `sdd-tasks` | Break down a change into implementation tasks |
| SDD Apply | `sdd-apply` | Implement tasks following specs and design |
| SDD Verify | `sdd-verify` | Validate implementation matches specs |
| SDD Archive | `sdd-archive` | Sync delta specs to main specs and archive |
| Judgment Day | `judgment-day` | Parallel adversarial review — two independent judges review the same target |

#### Foundation

| Skill | ID | Description |
|-------|-----|-------------|
| Go Testing | `go-testing` | Go testing patterns including Bubbletea TUI testing |
| Skill Creator | `skill-creator` | Create new AI agent skills following the Agent Skills spec |
| Branch & PR | `branch-pr` | PR creation workflow with conventional commits, branch naming, and issue-first enforcement |
| Issue Creation | `issue-creation` | Issue filing workflow with bug report and feature request templates |

These foundation skills are installed by default with both `full-pentest` and `ecosystem-core` presets.

### Coding Skills (separate repository)

For framework-specific development skills (React 19, Angular, TypeScript, Tailwind 4, Zod 4, Playwright, etc.), use an external/community skill repository such as [Gentleman-Programming/Gentleman-Skills](https://github.com/Gentleman-Programming/Gentleman-Skills). These are separate from Ciberbal's bundled pentest and workflow skills.

---

## Presets

| Preset | ID | What's Included |
|--------|-----|-----------------|
| Full Pentest | `full-pentest` | Default full-stack install: all core components, all bundled skills, and all pentesting capability layers provisioned together |
| Ecosystem Core | `ecosystem-core` | Shared workflow/core capability layers without the full persona/tooling set |
| Minimal | `minimal` | Engram + SDD skills only |
| Custom | `custom` | You pick components, skills, and persona individually |

## Domain-aware capability layers

`ciberbal-ai` is moving toward a model where install-time setup stays simple, but the installed ecosystem includes capability layers for:

- Web
- API
- Mobile
- AD/Internal
- WiFi/Wireless
- Cloud
- Recon
- Reporting

These are not mutually exclusive install branches. They are part of the same pentesting ecosystem and will increasingly shape skills, docs, prompts, and workflows.

## Bundled domain and workflow packs

Under `full-pentest`, ciberbal-ai now bundles lightweight post-install packs for:

- Domains: Web, API, Mobile, AD/Internal, WiFi/Wireless, Cloud, Recon, Reporting
- Workflows: Scoping, Recon, Enumeration, Exploitation, Post-exploitation, Evidence, Reporting

These packs are installed together as guidance layers. They are meant to orient the operator after setup, not to force specialty selection during installation.
