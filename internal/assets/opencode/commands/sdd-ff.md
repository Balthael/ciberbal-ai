---
description: Fast-forward engagement planning phases — scope through enumeration tasks
agent: ciberbal
---

Follow the engagement orchestrator workflow to fast-forward all planning phases for engagement "$ARGUMENTS".

AUTHORIZED SCOPE GUARD: Only operate against targets you are explicitly authorized to test (HTB machines, personal labs, or engagement targets covered by a signed ROE). Do NOT proceed if no authorized target is defined.

WORKFLOW:
Run these sub-agents in sequence:
1. sdd-propose — define scope and Rules of Engagement
2. sdd-spec — document findings spec and evidence requirements
3. sdd-design — plan the attack design and technique selection
4. sdd-tasks — break down into actionable enumeration tasks

Present a combined summary after ALL phases complete (not between each one).

CONTEXT:
- Working directory: !`echo -n "$(pwd)"`
- Current project: !`echo -n "$(basename $(pwd))"`
- Engagement name: $ARGUMENTS
- Artifact store mode: engram

ENGRAM NOTE:
Sub-agents handle persistence automatically. Each phase saves its artifact to engram with topic_key "sdd/$ARGUMENTS/{type}" where type is: proposal, spec, design, tasks.

Read the orchestrator instructions to coordinate this workflow. Do NOT execute phase work inline — delegate to sub-agents.
