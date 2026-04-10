# Proposal: ciberbal-ai-domain-skill-packs

## Intent

Define the next practical layer for ciberbal-ai by introducing domain skill packs and workflow packs for Web, API, Mobile, AD/Internal, WiFi/Wireless, Cloud, Recon, and Reporting. This gives users a full, domain-aware security toolkit by default without forcing an install-time choice, fitting perfectly into the unified single-pass installation model.

## Scope

### In Scope
- Define bundled skill sets, prompts, and playbooks for each of the 8 domains (Web, API, Mobile, AD/Internal, WiFi/Wireless, Cloud, Recon, Reporting).
- Ensure all packs are included in the default `full-pentest` preset.
- Implement workflow packs aligned with engagement phases (Scoping, Recon, Enumeration, Exploitation, Post-exploitation, Evidence, Reporting).

### Out of Scope
- Prompting the user to select specific domains during the installation flow.
- Deep, dynamic agentic behavior for every edge case within a domain; we are focusing on providing the structural packs and initial skill definitions.
- Changes to the top-level CLI menu structure.

## Capabilities

### New Capabilities
- `domain-skill-packs`: Bundled capability layers providing skills and playbooks for specific pentesting domains (Web, API, Mobile, AD/Internal, WiFi/Wireless, Cloud, Recon, Reporting).
- `workflow-packs`: Bundled capability layers providing phase-specific workflows (Scoping, Recon, Enumeration, Exploitation, Post-exploitation, Evidence, Reporting) to support engagement execution.

### Modified Capabilities
- None

## Approach

Create the content for the domain and workflow packs. Since the domain taxonomy was established in `ciberbal-ai-pentesting-domain-architecture`, we will now populate these domains with actual capability sets (skills, templates, and agent instructions). These packs will be statically defined and bundled into the default `full-pentest` preset, ensuring they are automatically deployed during a standard installation.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `internal/catalog/` | Modified | Update presets to bundle the new skill packs into the default pentesting ecosystem. |
| `internal/assets/skills/` | New | Add skill definitions, playbooks, and prompts for each domain and workflow phase. |
| `docs/` | Modified | Update documentation to detail the contents and usage of the new domain and workflow packs. |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Bundle size bloat | Medium | Keep initial skill packs lean and focused; rely on dynamic prompting where possible rather than massive static assets. |
| User overwhelmed by options post-install | Low | Clear documentation mapping engagement phases to the relevant domain skills. |

## Rollback Plan

Remove the added skill packs from `internal/assets/skills/` and revert the `internal/catalog/` changes that bundle them into the presets. The system will fall back to the base architecture defined in the previous phase without the new capabilities.

## Dependencies

- Existing Engram integration and SDD workflow engine.
- `ciberbal-ai-pentesting-domain-architecture` (taxonomy and presets).

## Success Criteria

- [ ] All 8 domain skill packs are defined and available post-installation.
- [ ] Workflow phase packs are accessible.
- [ ] The `full-pentest` preset bundles all packs without prompting the user during installation.
