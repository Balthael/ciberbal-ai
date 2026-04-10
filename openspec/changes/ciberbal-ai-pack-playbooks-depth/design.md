# Design — ciberbal-ai-pack-playbooks-depth

## Overview

This change deepens the 15 existing bundled domain/workflow packs from short summaries into structured operational playbooks.

The implementation intentionally avoids changing:

- install flow
- preset wiring
- catalog membership
- runtime UX

Only the pack content depth and its structural verification change.

## Pack schema

Each pack should use this required section order:

1. `## Objective`
2. `## Checklist`
3. `## Outputs`
4. `## Guardrails`
5. `## Evidence`

Optional later sections:

- `## Common Pitfalls`
- `## Escalation Triggers`
- `## Reporting Notes`

## Verification strategy

Add a single structure test in `internal/assets/pack_structure_test.go` that loads all 15 embedded pack files via `assets.Read()` and verifies the required headings are present.

This gives us:

- fast feedback
- structural enforcement for future edits
- spec coverage for missing mandatory sections

## Content strategy

The playbooks should be:

- operational, not marketing
- concise, not bloated
- reusable by both humans and agents
- evidence-aware and scope-aware

Each pack should provide a practical checklist and expected outputs, not just a paragraph of description.

## File impact

- `internal/assets/pack_structure_test.go` (new)
- all 8 existing `internal/assets/skills/domain-*/SKILL.md`
- all 7 existing `internal/assets/skills/workflow-*/SKILL.md`

## Risks

- too much text can reduce usability
- inconsistent wording across packs can make quality uneven
- structure-only tests do not guarantee deep content quality, but they are the right MVP enforcement layer

## Success shape

When complete:

- all 15 packs follow the same required schema
- pack content feels operational rather than placeholder-level
- tests fail if future edits remove mandatory structure
