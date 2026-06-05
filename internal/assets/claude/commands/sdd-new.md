---
description: Start a new engagement workflow — run recon then scope proposal
---

Follow the Ciberbal engagement workflow inline using the instructions already installed in `~/.claude/CLAUDE.md`.
The Claude Code session model is controlled by Claude Code; Ciberbal AI only configures models for Agent tool calls to phase sub-agents.

AUTHORIZED SCOPE GUARD: Only operate against targets you are explicitly authorized to test (HTB machines, personal labs, CTF environments, or engagement targets covered by a signed ROE). Do NOT proceed if no authorized target is defined.

WORKFLOW:

1. Launch `sdd-explore` to run recon and map the attack surface for this engagement
2. Present the recon summary to the user
3. Launch `sdd-propose` to create a scope/ROE proposal based on the recon
4. Present the proposal summary and ask the user if they want to continue with findings spec and attack design

CONTEXT:

- Working directory: !`pwd`
- Current project: !`basename "$(pwd)"`
- Engagement name: $ARGUMENTS
- Execution mode: ask/cache per orchestrator
- Artifact store mode: ask/cache per orchestrator
- Delivery strategy: ask/cache per orchestrator

ENGRAM NOTE:
Sub-agents handle persistence automatically. Each phase saves its artifact to engram with topic_key "sdd/$ARGUMENTS/{type}".

Read the orchestrator instructions to coordinate this engagement. Do NOT execute phase work inline when a native sub-agent is available.
