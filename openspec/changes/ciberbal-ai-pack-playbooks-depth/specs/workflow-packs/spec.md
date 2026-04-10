# Delta for workflow-packs

## ADDED Requirements

### Requirement: Structured Operational Content in Workflow Packs

All workflow phase packs MUST contain standardized operational sections to provide actionable guidance during an engagement. The mandatory sections MUST include: Objective, Checklist, Outputs, Guardrails, and Evidence. Optional sections MAY include Common Pitfalls, Escalation Triggers, and Reporting Notes.

#### Scenario: Validating workflow pack structure

- GIVEN a workflow phase pack (e.g., Recon, Exploitation, Scoping)
- WHEN the pack is processed or viewed
- THEN it MUST contain the Objective, Checklist, Outputs, Guardrails, and Evidence sections
- AND it MAY contain Common Pitfalls, Escalation Triggers, or Reporting Notes

#### Scenario: Missing mandatory sections in a workflow pack

- GIVEN a workflow phase pack
- WHEN the pack is missing one or more mandatory sections (e.g., lacks Evidence)
- THEN the system or reviewer SHOULD flag the pack as non-compliant with the required operational structure