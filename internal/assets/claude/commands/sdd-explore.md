---
description: Recon — investigate target surface and compare attack approaches
---

If the native `sdd-explore` sub-agent is available, delegate this command to it.
Otherwise, read the skill file at `~/.claude/skills/sdd-explore/SKILL.md` FIRST, then follow its instructions exactly inline.

AUTHORIZED SCOPE GUARD: Only operate against targets you are explicitly authorized to test (HTB machines, personal labs, CTF environments, or engagement targets covered by a signed ROE). Do NOT proceed if no authorized target is defined.

CONTEXT:
- Working directory: !`pwd`
- Current project: !`basename "$(pwd)"`
- Target or recon topic: $ARGUMENTS
- Artifact store mode: engram

TASK:
Run recon for "$ARGUMENTS". Investigate the authorized target surface, identify exposed services or paths, compare attack approaches, and provide a scoped recommendation.

ENGRAM PERSISTENCE (artifact store mode: engram):
Read project context (optional):
  mem_search(query: "sdd-init/{project}", project: "{project}") → if found, mem_get_observation(id) for full content
Save exploration:
  mem_save(title: "sdd/$ARGUMENTS/explore", topic_key: "sdd/$ARGUMENTS/explore", type: "architecture", project: "{project}", capture_prompt: false, content: "{exploration}")
  Set capture_prompt: false when the Engram tool schema supports it; if an older schema rejects or does not expose the field, omit it rather than failing.

This is recon only — do NOT exploit, modify systems, or go out of scope. Research and return your analysis.

Return a structured result with: status, executive_summary, detailed_report, artifacts, and next_recommended.
