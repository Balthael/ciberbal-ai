# Tasks: Deepen bundled pentesting playbooks

## Phase 1: Testing Foundation (RED)

- [x] 1.1 Create `internal/assets/pack_structure_test.go` with a table covering all 15 embedded pack paths and failing assertions for `## Objective`, `## Checklist`, `## Outputs`, `## Guardrails`, and `## Evidence`.
- [x] 1.2 In `internal/assets/pack_structure_test.go`, add fixture helpers that load packs through `assets.Read` and report exactly which mandatory heading is missing per file.

## Phase 2: Domain Pack Enrichment (GREEN)

- [x] 2.1 Rewrite `internal/assets/skills/domain-web/SKILL.md`, `domain-api/SKILL.md`, and `domain-mobile/SKILL.md` into the MVP playbook format with concise operational guidance under the five mandatory sections.
- [x] 2.2 Rewrite `internal/assets/skills/domain-ad-internal/SKILL.md`, `domain-wifi-wireless/SKILL.md`, and `domain-cloud/SKILL.md` with the same section order, tone, and evidence-focused depth.
- [x] 2.3 Rewrite `internal/assets/skills/domain-recon/SKILL.md` and `domain-reporting/SKILL.md` so their checklists and outputs stay domain-specific while matching the required structure.

## Phase 3: Workflow Pack Enrichment (GREEN)

- [x] 3.1 Rewrite `internal/assets/skills/workflow-scoping/SKILL.md`, `workflow-recon/SKILL.md`, and `workflow-enumeration/SKILL.md` into structured phase playbooks using the same mandatory headings.
- [x] 3.2 Rewrite `internal/assets/skills/workflow-exploitation/SKILL.md` and `workflow-post-exploitation/SKILL.md` with actionable guardrails and evidence expectations for high-risk phases.
- [x] 3.3 Rewrite `internal/assets/skills/workflow-evidence/SKILL.md` and `workflow-reporting/SKILL.md` so artifact handling, reporting outputs, and reviewer expectations are explicit.

## Phase 4: Verification and Refactor

- [x] 4.1 Refine `internal/assets/pack_structure_test.go` so the final assertions cover all 15 packs, keep optional sections optional, and stay readable for future pack additions.
- [x] 4.2 Run `go test ./internal/assets ./internal/components/skills ./internal/catalog` and fix any regressions tied to embedded assets, preset coverage, or bundled-pack visibility.
- [x] 4.3 Run `go test ./...` and `go vet ./...`, then verify the final diff only changes the 15 pack files plus test coverage required to enforce the new structure.
