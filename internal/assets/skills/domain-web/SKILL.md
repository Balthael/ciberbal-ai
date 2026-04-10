# Domain Pack — Web

## Objective

Guide browser-facing and authenticated web application assessments toward real attack-surface validation, reproducible findings, and clear remediation output.

## Checklist

- Map public, authenticated, and privileged routes before testing assumptions.
- Validate authn/authz, session handling, CSRF, input handling, and business-logic abuse paths.
- Compare intended workflow behavior against what an attacker can actually force or bypass.

## Outputs

- Prioritized list of validated web findings with reproduction steps.
- Attack paths tied to user roles, trust boundaries, and business impact.
- Notes on remediation direction for engineering teams.

## Guardrails

- Stay inside approved hosts, accounts, and test windows.
- Distinguish reflected noise from exploitable impact.
- Do not report theoretical browser issues without proof of effect.

## Evidence

- Requests, responses, screenshots, session state, and replay notes.
- Role comparisons showing what changed and why it matters.
