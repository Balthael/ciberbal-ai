# Domain Pack — Cloud

## Objective

Guide cloud assessments toward validated IAM, storage, network, secret, and service-configuration findings that explain business impact clearly.

## Checklist

- Inventory accounts, projects, identities, roles, exposed services, and sensitive storage paths.
- Validate IAM abuse paths, public exposure, secret leakage, service misconfiguration, and over-permissioned identities.
- Explain whether the issue affects control plane, data plane, or both.

## Outputs

- Findings tied to identities, resources, and reachable blast radius.
- Clear attack-path explanation from initial foothold to sensitive action.
- Remediation framing aligned to least privilege and exposure reduction.

## Guardrails

- Avoid destructive changes to resources unless explicitly approved.
- Distinguish misconfiguration from proven exploitability when impact is uncertain.
- Keep account/project identifiers and secret material sanitized in outputs.

## Evidence

- Role assignments, policy excerpts, resource paths, and action proof.
- Screenshots or command output showing exposure, access, or configuration state.
