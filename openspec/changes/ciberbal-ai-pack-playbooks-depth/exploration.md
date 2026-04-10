# Exploration — ciberbal-ai-pack-playbooks-depth

## Goal

Deepen the existing bundled domain and workflow packs so they stop feeling like lightweight placeholders and start acting like practical pentesting playbooks, while preserving the unified install model.

## Current state

The previous change (`ciberbal-ai-domain-skill-packs`) already delivered:

- 8 bundled domain packs
- 7 bundled workflow packs
- `full-pentest` bundling for all packs
- docs explaining post-install capability discovery

However, the verify/archive reports explicitly recorded that the pack content is still **too lean**. The current assets are enough to prove structure and bundling, but not enough to feel like operational guidance.

## What richer packs should contain

The next useful step is to standardize each pack around sections like:

- Objective
- Checklist
- Outputs
- Guardrails
- Evidence

Optional later sections can include:

- Common Pitfalls
- Escalation Triggers
- Reporting Notes

## Why this matters

Without richer content, the current packs are mostly labels plus short blurbs. That is enough for product structure, but not enough for real operator value.

Richer packs improve:

- prompt usefulness
- consistency across engagements
- evidence hygiene
- reporting quality
- domain-to-phase execution discipline

## MVP depth vs later depth

### MVP now
- rewrite all 15 packs to the richer structured format
- keep content concise but operational
- standardize section order and tone

### Later
- dynamic runtime pack discovery/UX
- deeper cross-links between domains and phases
- richer examples/templates/artifacts per pack
- domain-specific reporting templates

## Risks

1. **Too verbose**
   - packs could become bloated and hard to use
2. **Inconsistent structure**
   - if sections drift across files, maintainability drops
3. **Operational shallowness hidden by longer text**
   - more text is not automatically better; every section must still be useful

## Recommended direction

Use the existing static assets and deepen them in place, keeping:

- bundling unchanged
- preset logic unchanged
- install flow unchanged

This is the lowest-risk way to add real user value immediately.
