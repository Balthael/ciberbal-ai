---
description: Start a new engagement workflow — run recon then scope proposal
agent: ciberbal
---

Follow the engagement orchestrator workflow for starting a new pentest engagement named "$ARGUMENTS".

AUTHORIZED SCOPE GUARD: Only operate against targets you are explicitly authorized to test (HTB machines, personal labs, or engagement targets covered by a signed ROE). Do NOT proceed if no authorized target is defined.

WORKFLOW:
1. Launch sdd-explore sub-agent to run recon and map the attack surface for this engagement
2. Present the recon summary to the user
3. Launch sdd-propose sub-agent to create a scope/ROE proposal based on the recon
4. Present the proposal summary and ask the user if they want to continue with findings spec and attack design

CONTEXT:
- Working directory: !`echo -n "$(pwd)"`
- Current project: !`echo -n "$(basename $(pwd))"`
- Engagement name: $ARGUMENTS
- Artifact store mode: engram

ENGRAM NOTE:
Sub-agents handle persistence automatically. Each phase saves its artifact to engram with topic_key "sdd/$ARGUMENTS/{type}".

Read the orchestrator instructions to coordinate this workflow. Do NOT execute phase work inline — delegate to sub-agents.
