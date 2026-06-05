---
description: Reporting Archive — close engagement and preserve artifacts
---

If the native `sdd-archive` sub-agent is available, delegate this command to it.
Otherwise, read the skill file at `~/.claude/skills/sdd-archive/SKILL.md` FIRST, then follow its instructions exactly inline.

AUTHORIZED SCOPE GUARD: Only archive artifacts for authorized targets (HTB machines, personal labs, CTF environments, or engagement targets covered by a signed ROE). Do NOT include out-of-scope material.

CONTEXT:
- Working directory: !`pwd`
- Current project: !`basename "$(pwd)"`
- Artifact store mode: engram

TASK:
Archive the active engagement. Read the evidence review report first to confirm the engagement is ready to close. Then:

ENGRAM PERSISTENCE (artifact store mode: engram):
CRITICAL: mem_search returns 300-char PREVIEWS, not full content. You MUST call mem_get_observation(id) for EVERY artifact.
STEP A — SEARCH (get IDs only):
  mem_search(query: "sdd/{change-name}/proposal", project: "{project}") → save proposal_id
  mem_search(query: "sdd/{change-name}/spec", project: "{project}") → save spec_id
  mem_search(query: "sdd/{change-name}/design", project: "{project}") → save design_id
  mem_search(query: "sdd/{change-name}/tasks", project: "{project}") → save tasks_id
  mem_search(query: "sdd/{change-name}/verify-report", project: "{project}") → save verify_id
STEP B — RETRIEVE FULL CONTENT (mandatory):
  mem_get_observation(id: proposal_id) → full scope/ROE proposal
  mem_get_observation(id: spec_id) → full findings spec
  mem_get_observation(id: design_id) → full attack design
  mem_get_observation(id: tasks_id) → full enumeration tasks
  mem_get_observation(id: verify_id) → full evidence review report
Record all observation IDs in the archive report for traceability.
Save:
  mem_save(title: "sdd/{change-name}/archive-report", topic_key: "sdd/{change-name}/archive-report", type: "architecture", project: "{project}", capture_prompt: false, content: "{archive report with observation IDs}")
  Set capture_prompt: false when the Engram tool schema supports it; if an older schema rejects or does not expose the field, omit it rather than failing.

Then:
1. Compile engagement artifacts and evidence references
2. Move the engagement folder to archive with date prefix when using file-backed artifacts
3. Verify the archive/reporting bundle is complete

Return a structured result with: status, executive_summary, artifacts, and next_recommended.
