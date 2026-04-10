# Tasks: Domain Skill Packs

## Phase 1: Foundation + RED tests

- [x] 1.1 Add failing coverage in `internal/assets/assets_test.go` and `internal/components/skills/presets_test.go` for 8 domain packs, 7 workflow packs, updated embedded asset counts, and `full-pentest` membership.
- [x] 1.2 Extend `internal/model/types.go` with `SkillID` constants for `domain-*` and `workflow-*` packs; register them in `internal/catalog/skills.go` with clear categories/priorities.
- [x] 1.3 Refactor `internal/components/skills/presets.go` so `full-pentest` installs new `domainSkills` + `workflowSkills`, while `minimal` and `ecosystem-core` keep current scope.

## Phase 2: MVP domain pack content

- [x] 2.1 Create `internal/assets/skills/domain-web/SKILL.md`, `domain-api/SKILL.md`, `domain-mobile/SKILL.md`, and `domain-ad-internal/SKILL.md` with lean operator playbooks, guardrails, and engagement prompts.
- [x] 2.2 Create `internal/assets/skills/domain-wifi-wireless/SKILL.md`, `domain-cloud/SKILL.md`, `domain-recon/SKILL.md`, and `domain-reporting/SKILL.md` with the same MVP structure.

## Phase 3: MVP workflow packs + wiring

- [x] 3.1 Create `internal/assets/skills/workflow-scoping/SKILL.md`, `workflow-recon/SKILL.md`, `workflow-enumeration/SKILL.md`, and `workflow-exploitation/SKILL.md` with phase-specific checklists and expected outputs.
- [x] 3.2 Create `internal/assets/skills/workflow-post-exploitation/SKILL.md`, `workflow-evidence/SKILL.md`, and `workflow-reporting/SKILL.md` with reproducibility, evidence, and reporting guidance.
- [x] 3.3 Expand `internal/components/skills/inject_test.go` to verify representative domain/workflow packs write correctly for supported agents and stay idempotent.

## Phase 4: Catalog/docs + verification

- [x] 4.1 Update `internal/catalog/skills_test.go` and any affected catalog assertions so preset/catalog coverage catches duplicate or missing domain/workflow IDs.
- [x] 4.2 Update `docs/components.md`, `docs/pentesting-domains.md`, and `README.md` to explain post-install domain/workflow packs, `full-pentest` bundling, and the domain-to-phase usage model.
- [x] 4.3 Verify the MVP against the spec with `go test ./...` and `go vet ./...`, then confirm `full-pentest` bundles all 15 new packs without install-time branching.
