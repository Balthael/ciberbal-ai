---
description: Fast-forward engagement planning phases — scope through enumeration tasks
---

Follow the Ciberbal engagement workflow inline using the instructions already installed in `~/.claude/CLAUDE.md`.
The Claude Code session model is controlled by Claude Code; Gentle AI only configures models for Agent tool calls to phase sub-agents.

AUTHORIZED SCOPE GUARD: Only operate against targets you are explicitly authorized to test (HTB machines, personal labs, CTF environments, or engagement targets covered by a signed ROE). Do NOT proceed if no authorized target is defined.

WORKFLOW:
Run these sub-agents in sequence:

1. `sdd-propose` — create the scope/ROE proposal
2. `sdd-spec` — write findings specifications and evidence requirements
3. `sdd-design` — create the attack design
4. `sdd-tasks` — break down into enumeration tasks and actionable checks

Present a combined summary after ALL phases complete (not between each one).

CONTEXT:

- Working directory: !`pwd`
- Current project: !`basename "$(pwd)"`
- Engagement name: $ARGUMENTS
- Execution mode: ask/cache per orchestrator
- Artifact store mode: ask/cache per orchestrator
- Delivery strategy: ask/cache per orchestrator

ENGRAM NOTE:
Sub-agents handle persistence automatically. Each phase saves its artifact to engram with topic_key "sdd/$ARGUMENTS/{type}" where type is: proposal, spec, design, tasks.

Read the orchestrator instructions to coordinate this workflow. Do NOT execute phase work inline when a native sub-agent is available.
