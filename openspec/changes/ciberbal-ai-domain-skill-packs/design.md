# Design — ciberbal-ai-domain-skill-packs

## Overview

This change adds the first practical post-install capability packs for the pentesting ecosystem without changing the unified installation model.

The design introduces two bundled pack families:

- **Domain packs** for Web, API, Mobile, AD/Internal, WiFi/Wireless, Cloud, Recon, and Reporting
- **Workflow packs** for Scoping, Recon, Enumeration, Exploitation, Post-exploitation, Evidence, and Reporting

All of them are bundled into `full-pentest`.

## Design decisions

### 1) Packs are implemented as bundled skill assets

The repository already has a working embedded-skill distribution mechanism. The lowest-risk MVP is to represent domain/workflow packs as additional embedded `SKILL.md` assets under `internal/assets/skills/`.

### 2) `full-pentest` expands, other presets stay narrow

- `minimal` remains SDD-only
- `ecosystem-core` remains shared workflow/core only
- `full-pentest` expands to include domain + workflow packs

This preserves install simplicity while making the full preset materially more useful.

### 3) Packs are guidance-first, not deep automation

The MVP content should provide:

- operator framing
- playbook/checklist guidance
- evidence/reporting reminders

It should not try to encode every edge case of each domain.

## File impact

- `internal/model/types.go`
- `internal/catalog/skills.go`
- `internal/components/skills/presets.go`
- `internal/assets/skills/*`
- `internal/assets/assets_test.go`
- `internal/components/skills/inject_test.go`
- `internal/catalog/skills_test.go`
- docs surfaces explaining bundled packs

## Sequencing

1. Add new skill IDs and catalog metadata
2. Bundle them into `full-pentest`
3. Add embedded pack assets
4. Extend tests for asset loading, preset membership, and injection
5. Update docs

## Risks

- Bundle growth over time
- Packs becoming too generic if not refined iteratively
- Confusion if docs imply install-time selection rather than post-install capability layering
