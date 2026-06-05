---
description: Continue next engagement phase in the dependency chain
---

Follow the Ciberbal engagement workflow inline using the instructions already installed in `~/.claude/CLAUDE.md`.
The Claude Code session model is controlled by Claude Code; Ciberbal AI only configures models for Agent tool calls to phase sub-agents.

AUTHORIZED SCOPE GUARD: Only operate against targets you are explicitly authorized to test (HTB machines, personal labs, CTF environments, or engagement targets covered by a signed ROE). Do NOT proceed if no authorized target is defined.

WORKFLOW:

1. Check which artifacts already exist for the active engagement (scope/ROE proposal, findings spec, attack design, enumeration tasks)
2. Determine the next phase needed based on the dependency graph:
   scope proposal → [findings spec ∥ attack design] → enumeration tasks → exploitation → evidence review → reporting archive
3. Launch the appropriate sub-agent(s) for the next phase
4. Present the result and ask the user to proceed

CONTEXT:

- Working directory: !`pwd`
- Current project: !`basename "$(pwd)"`
- Engagement name: $ARGUMENTS
- Execution mode: ask/cache per orchestrator
- Artifact store mode: ask/cache per orchestrator
- Delivery strategy: ask/cache per orchestrator

ENGRAM NOTE:
To check which artifacts exist, search: mem_search(query: "sdd/$ARGUMENTS/", project: "{project}") to list all artifacts for this change.
Sub-agents handle persistence automatically with topic_key "sdd/$ARGUMENTS/{type}".

Read the orchestrator instructions to coordinate this workflow. Do NOT execute phase work inline when a native sub-agent is available.
