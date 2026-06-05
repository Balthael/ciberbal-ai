---
description: Continue next engagement phase in the dependency chain
agent: ciberbal
---

Follow the engagement orchestrator workflow to continue the active engagement.

AUTHORIZED SCOPE GUARD: Only operate against targets you are explicitly authorized to test (HTB machines, personal labs, or engagement targets covered by a signed ROE). Do NOT proceed if no authorized target is defined.

WORKFLOW:
1. Check which artifacts already exist for the active engagement (scope/ROE proposal, findings spec, attack design, enumeration tasks)
2. Determine the next phase needed based on the dependency graph:
   proposal → [spec ∥ design] → tasks → apply → verify → archive
3. Launch the appropriate sub-agent(s) for the next phase
4. Present the result and ask the user to proceed

CONTEXT:
- Working directory: !`echo -n "$(pwd)"`
- Current project: !`echo -n "$(basename $(pwd))"`
- Engagement name: $ARGUMENTS
- Artifact store mode: engram

ENGRAM NOTE:
To check which artifacts exist, search: mem_search(query: "sdd/$ARGUMENTS/", project: "{project}") to list all artifacts for this engagement.
Sub-agents handle persistence automatically with topic_key "sdd/$ARGUMENTS/{type}".

Read the orchestrator instructions to coordinate this workflow. Do NOT execute phase work inline — delegate to sub-agents.
