---
description: Evidence Review — validate findings against scope and evidence requirements
---

If the native `sdd-verify` sub-agent is available, delegate this command to it.
Otherwise, read the skill file at `~/.claude/skills/sdd-verify/SKILL.md` FIRST, then follow its instructions exactly inline.

AUTHORIZED SCOPE GUARD: Only operate against targets you are explicitly authorized to test (HTB machines, personal labs, CTF environments, or engagement targets covered by a signed ROE). Do NOT proceed if no authorized target is defined.

CONTEXT:
- Working directory: !`pwd`
- Current project: !`basename "$(pwd)"`
- Artifact store mode: engram

TASK:
Review evidence for the active engagement. Read the findings spec, attack design, and enumeration tasks artifacts. Then:

ENGRAM PERSISTENCE (artifact store mode: engram):
CRITICAL: mem_search returns 300-char PREVIEWS, not full content. You MUST call mem_get_observation(id) for EVERY artifact.
STEP A — SEARCH (get IDs only):
  mem_search(query: "sdd/{change-name}/spec", project: "{project}") → save spec_id
  mem_search(query: "sdd/{change-name}/design", project: "{project}") → save design_id
  mem_search(query: "sdd/{change-name}/tasks", project: "{project}") → save tasks_id
STEP B — RETRIEVE FULL CONTENT (mandatory):
  mem_get_observation(id: spec_id) → full findings spec
  mem_get_observation(id: design_id) → full attack design
  mem_get_observation(id: tasks_id) → full enumeration tasks
Save report:
  mem_save(title: "sdd/{change-name}/verify-report", topic_key: "sdd/{change-name}/verify-report", type: "architecture", project: "{project}", capture_prompt: false, content: "{verification report}")
  Set capture_prompt: false when the Engram tool schema supports it; if an older schema rejects or does not expose the field, omit it rather than failing.

Then:
1. Check completeness — are all planned checks or exploitation tasks done?
2. Check scope — did every action remain within the ROE?
3. Check coherence — were attack design decisions followed?
4. Validate evidence quality and reproducibility
5. Build the findings/evidence compliance matrix

Return a structured verification report with: status, executive_summary, detailed_report, artifacts, and next_recommended.
