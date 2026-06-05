---
description: Guided engagement walkthrough — authorized target or lab
---

If the native `sdd-onboard` sub-agent is available, delegate this command to it.
Otherwise, read the skill file at `~/.claude/skills/sdd-onboard/SKILL.md` FIRST, then follow its instructions exactly inline.

AUTHORIZED SCOPE GUARD: Only operate against targets you are explicitly authorized to test (HTB machines, personal labs, CTF environments, or engagement targets covered by a signed ROE). Do NOT proceed if no authorized target is defined.

CONTEXT:
- Working directory: !`pwd`
- Current project: !`basename "$(pwd)"`
- Artifact store mode: engram

TASK:
Guide the user through a complete authorized engagement workflow using their target or lab. This is real scoped work with real artifacts, not a toy example. The goal is to teach by doing — walk through recon, scope/ROE, findings spec, attack design, enumeration tasks, exploitation, evidence review, and reporting archive.

ENGRAM PERSISTENCE (artifact store mode: engram):
Save onboarding progress as you go:
  mem_save(title: "sdd-onboard/{project}", topic_key: "sdd-onboard/{project}", type: "architecture", project: "{project}", capture_prompt: false, content: "{onboarding state}")
  Set capture_prompt: false when the Engram tool schema supports it; if an older schema rejects or does not expose the field, omit it rather than failing.
topic_key enables upserts — re-running updates, not duplicates.

Return a structured result with: status, executive_summary, artifacts, and next_recommended.
