# Proposal: ciberbal-ai-menu-and-install-flow

## Intent

Adapt the existing gentle-ai CLI tool into a cybersecurity/pentesting specialized tool called `ciberbal-ai`. The core objective is to streamline the installation and configuration flow for pentesters by introducing a Quick/Advanced installation mode, while preserving the foundational architecture, memory, agent creation, and capabilities of gentle-ai.

## Scope

### In Scope
- Rename branding from `gentle-ai` to `ciberbal-ai`.
- Preserve the main welcome menu structure and its core capabilities (memory, SDD, model config, backups, profiles, sync/upgrade, agent creation).
- Add an "Install Mode" selection screen (Quick vs. Advanced) inside the "Start installation" flow.
- Implement "Quick install" to automatically install all essential cybersecurity/pentesting dependencies without asking for a specialty.
- Retain "Advanced install" with the granular configuration options currently present.
- Support interactive mode planning and hybrid artifact storage.

### Out of Scope
- Removing or heavily modifying the existing agent execution engine.
- Adding new non-cybersecurity capabilities to the core framework.
- Changes to the underlying Git bash (`gga`) integration.

## Capabilities

### New Capabilities
- `installer-flow`: Defines the new dual-mode (Quick/Advanced) installation flow and default cybersecurity dependencies for the Quick mode.
- `ciberbal-branding`: Defines the rebranding from gentle-ai to ciberbal-ai across the welcome menu and CLI outputs.

### Modified Capabilities
- None

## Approach

We will fork or update the main CLI entry points to reflect the `ciberbal-ai` branding. We will refactor the installation sequence initiated from the welcome menu to first prompt for "Quick" or "Advanced" mode. If "Quick" is selected, we will bypass the granular capability prompts and execute a predefined list of pentesting dependency installations. The underlying configuration storage, profiles, and SDD integrations will remain untouched.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `cmd/` | Modified | CLI entrypoints renamed to ciberbal-ai |
| `internal/tui/menu/` | Modified | Welcome menu branding updates |
| `internal/installer/` | Modified | Introduction of Quick/Advanced flow |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Quick install misses tools | Medium | Maintain an easily updateable list of default tools for the Quick mode. |
| Existing profiles break | Low | Configuration formats remain identical; only default values change during Quick install. |

## Rollback Plan

Revert the commits introducing the new install flow and branding changes. Since the underlying configuration format is not changing structurally, existing user profiles will remain compatible with the previous version.

## Dependencies

- None

## Success Criteria

- [ ] Running the CLI shows `ciberbal-ai` branding in the main welcome menu.
- [ ] Selecting "Start installation" prompts for "Quick" or "Advanced" mode.
- [ ] "Quick install" completes without further prompting and installs pentesting defaults.
- [ ] "Advanced install" allows granular configuration as before.