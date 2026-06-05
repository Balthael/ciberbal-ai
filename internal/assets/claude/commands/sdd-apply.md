---
description: Exploitation — execute authorized tasks and collect evidence
---

If the native `sdd-apply` sub-agent is available, delegate this command to it.
Otherwise, read the skill file at `~/.claude/skills/sdd-apply/SKILL.md` FIRST, then follow its instructions exactly inline.

AUTHORIZED SCOPE GUARD: Only operate against targets you are explicitly authorized to test (HTB machines, personal labs, CTF environments, or engagement targets covered by a signed ROE). Do NOT proceed if no authorized target is defined.

CONTEXT:
- Working directory: !`pwd`
- Current project: !`basename "$(pwd)"`
- Artifact store mode: engram

TASK:
Execute the remaining incomplete exploitation tasks for the active engagement. Collect evidence for each completed step.

ENGRAM PERSISTENCE (artifact store mode: engram):
CRITICAL: mem_search returns 300-char PREVIEWS, not full content. You MUST call mem_get_observation(id) for EVERY artifact.
STEP A — SEARCH (get IDs only):
  mem_search(query: "sdd/{change-name}/spec", project: "{project}") → save spec_id
  mem_search(query: "sdd/{change-name}/design", project: "{project}") → save design_id
  mem_search(query: "sdd/{change-name}/tasks", project: "{project}") → save tasks_id
STEP A2 — CHECK PREVIOUS PROGRESS (before starting work):
  mem_search(query: "sdd/{change-name}/apply-progress", project: "{project}") → if found, save progress_id
  - Previous apply-progress (if exists): `mem_search(query: "sdd/{change-name}/apply-progress", project: "{project}")` → read and merge
STEP B — RETRIEVE FULL CONTENT (mandatory):
  mem_get_observation(id: spec_id) → full findings spec
  mem_get_observation(id: design_id) → full attack design
  mem_get_observation(id: tasks_id) → full enumeration tasks (keep tasks_id for updates)
  IF progress_id exists: mem_get_observation(id: progress_id) → read previous progress, skip completed tasks, MERGE when saving
Update tasks as you complete them:
  mem_update(id: {tasks-observation-id}, content: "{updated tasks with [x] marks}")
Save progress:
  mem_save(title: "sdd/{change-name}/apply-progress", topic_key: "sdd/{change-name}/apply-progress", type: "architecture", project: "{project}", capture_prompt: false, content: "{progress report}")
  Set capture_prompt: false when the Engram tool schema supports it; if an older schema rejects or does not expose the field, omit it rather than failing.

For each task:
1. Read the relevant findings spec scenarios (acceptance criteria / evidence requirements)
2. Read the attack design decisions (technique and tool selection)
3. Execute the authorized exploitation step
4. Collect and record evidence (output, hashes, flags, screenshots)
5. Mark the task as complete [x]

Return a structured result with: status, executive_summary, detailed_report (steps executed, evidence collected), artifacts, and next_recommended.
