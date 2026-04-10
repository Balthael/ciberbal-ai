# Domain Pack — AD/Internal

## Objective

Guide Active Directory and internal assessments toward identity-centric attack paths, privilege validation, and blast-radius understanding.

## Checklist

- Map domains, trusts, identities, privileges, reachable services, and administrative boundaries.
- Validate credential exposure, privilege escalation paths, delegation abuse, and lateral movement opportunities.
- Tie technical findings back to control failure and operational impact.

## Outputs

- Identity/path-based findings showing how access expands.
- Reachability notes explaining pivoting potential and affected systems.
- Prioritized hardening gaps with remediation framing.

## Guardrails

- Stay inside approved hosts, accounts, segments, and escalation boundaries.
- Avoid persistence or disruptive changes unless explicitly authorized.
- Record exactly what trust assumptions were required for each path.

## Evidence

- Hostnames, identities, privileges, command output, and graph/path evidence.
- Scope notes showing where a path stopped and why.
