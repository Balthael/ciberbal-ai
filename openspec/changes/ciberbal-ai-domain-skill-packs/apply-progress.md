## Apply Progress

**Change**: `ciberbal-ai-domain-skill-packs`
**Mode**: Strict TDD

### Completed work

- Added 8 domain pack `SkillID`s and 7 workflow pack `SkillID`s
- Registered pack IDs in the catalog
- Expanded `full-pentest` bundling to include all domain/workflow packs
- Added embedded `SKILL.md` assets for all 15 packs
- Extended asset, preset, catalog, injection, and golden tests
- Updated docs to explain bundled packs under `full-pentest`

### TDD Cycle Evidence

| Task | Test File | Layer | Safety Net | RED | GREEN | TRIANGULATE | REFACTOR |
|------|-----------|-------|------------|-----|-------|-------------|----------|
| 1.1 Skill IDs + preset wiring | `internal/catalog/skills_test.go`, `internal/components/skills/presets_test.go`, `internal/components/golden_test.go` | Unit + Integration | ✅ Passed | ✅ Written | ✅ Passed | ✅ 3 cases | ➖ None needed |
| 2.1 Embedded domain/workflow packs | `internal/assets/assets_test.go`, `internal/components/skills/inject_test.go` | Unit + Integration | ✅ Passed | ✅ Written | ✅ Passed | ✅ 2 cases | ➖ None needed |
| 4.2 Docs discovery alignment | `internal/catalog/docs_visibility_test.go` | Unit | N/A (new) | ✅ Written | ✅ Passed | ✅ 2 cases | ➖ None needed |

### Files Changed

- `internal/model/types.go`
- `internal/catalog/skills.go`
- `internal/catalog/skills_test.go`
- `internal/catalog/docs_visibility_test.go`
- `internal/components/skills/presets.go`
- `internal/components/skills/presets_test.go`
- `internal/components/skills/inject_test.go`
- `internal/assets/assets_test.go`
- `internal/assets/skills/domain-*/SKILL.md`
- `internal/assets/skills/workflow-*/SKILL.md`
- `testdata/golden/skills-presets.json`
- `README.md`
- `docs/components.md`
- `docs/pentesting-domains.md`

### Validation commands

- `go test ./...`
- `go vet ./...`
