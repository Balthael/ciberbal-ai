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

| Task | RED evidence | GREEN evidence | TRIANGULATE | SAFETY NET | Files Changed | Notes |
|------|--------------|----------------|-------------|------------|---------------|-------|
| Skill IDs + preset wiring | Existing preset/catalog tests failed until new IDs and bundling were added | `go test ./...` passed after updating model/catalog/presets and golden data | Added explicit tests for full-pentest domain/workflow membership | Preset/catalog/golden tests guard regressions in bundling | `internal/model/types.go`, `internal/catalog/skills.go`, `internal/components/skills/presets.go`, `internal/catalog/skills_test.go`, `internal/components/skills/presets_test.go`, `testdata/golden/skills-presets.json` | No install-time branching added |
| Embedded pack assets | Asset-count/readability expectations required new embedded files | `go test ./...` passed after all 15 pack assets were added | Injection tests verify representative domain/workflow packs write correctly for agents | Asset tests + inject tests guard missing or misnamed files | `internal/assets/assets_test.go`, `internal/assets/skills/*`, `internal/components/skills/inject_test.go` | MVP content is intentionally lean |
| Docs alignment | Docs needed to reflect bundled packs under the unified install model | `go test ./...` and `go vet ./...` passed after doc updates | Docs now reinforce the domain+workflow model introduced in prior architecture work | Existing architecture/install docs plus current tests reduce drift risk | `README.md`, `docs/components.md`, `docs/pentesting-domains.md` | Still no domain picker in install flow |

### Validation commands

- `go test ./...`
- `go vet ./...`
