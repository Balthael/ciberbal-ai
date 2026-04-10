# Proposal: ciberbal-ai-pack-playbooks-depth

## Intent

Turn the current lean domain and workflow packs into richer operational playbooks without changing installation simplicity. This upgrades the existing MVP summaries to provide deep, actionable guidance (Objectives, Checklists, Outputs, Guardrails, Evidence) for agents and users performing pentesting operations.

## Scope

### In Scope
- Expand all 8 domain packs (Web, API, Mobile, AD/Internal, WiFi/Wireless, Cloud, Recon, Reporting) with rich operational structure: Objective, Checklist, Outputs, Guardrails, Evidence.
- Expand all 7 workflow phase packs (Scoping, Recon, Enumeration, Exploitation, Post-exploitation, Evidence, Reporting) with the same rich operational structure.
- Ensure the enriched assets remain statically bundled within the `full-pentest` preset without altering the installation flow.

### Out of Scope
- Dedicated runtime UX/UI for pack discovery (deferred).
- Dynamic runtime rules mapping phases to domain playbooks automatically (deferred).
- Adding new domains or new workflow phases not currently present.

## Capabilities

### New Capabilities
- None

### Modified Capabilities
- `domain-skill-packs`: Enhancing the static skill assets with deeper operational structures.
- `workflow-packs`: Enhancing the phase-based static assets with deeper operational structures.

## Approach

Update the static markdown files located in `internal/assets/skills/` (and its subdirectories) corresponding to the 15 existing domain and workflow packs. Each file will be rewritten to include standard operational sections (Objective, Checklist, Outputs, Guardrails, Evidence). No changes are required in `internal/catalog/` presets since the file paths remain the same.

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `internal/assets/skills/domains/` | Modified | Update 8 domain markdown files with deeper operational content. |
| `internal/assets/skills/workflows/` | Modified | Update 7 workflow markdown files with deeper operational content. |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Binary bloat from larger assets | Low | Markdown compresses well; the total size increase will be negligible. |
| Overly complex agent instructions | Medium | Adhere to a strict schema (Objective, Checklist, etc.) to keep instructions readable and actionable. |

## Rollback Plan

Revert the commits modifying `internal/assets/skills/domains/` and `internal/assets/skills/workflows/` files. Since no code or catalog logic changes, restoring the previous lean asset files will cleanly downgrade the system back to the MVP summaries.

## Dependencies

- Previous change: `ciberbal-ai-domain-skill-packs` (where the files and basic structure were created).

## Success Criteria

- [ ] All 8 domain packs contain `Objective`, `Checklist`, `Outputs`, `Guardrails`, and `Evidence` sections.
- [ ] All 7 workflow phase packs contain the same structured sections.
- [ ] Application compiles successfully with no changes required in the install or catalog code.