# Domain Pack — API

## Objective

Guide API assessments toward validated auth, authorization, object access, schema abuse, and business-logic findings with replayable proof.

## Checklist

- Inventory REST, GraphQL, gRPC, and internal service endpoints in scope.
- Test authn/authz, BOLA/BFLA-style access, mass assignment, rate limiting, and workflow abuse.
- Confirm whether the issue is protocol-level, implementation-level, or business-logic-level.

## Outputs

- Endpoint-focused finding list with request/response proof.
- Trust-boundary notes explaining who can abuse what.
- Clear severity rationale tied to data exposure or action impact.

## Guardrails

- Avoid destructive mutations unless they are explicitly approved.
- Separate malformed-request noise from real authorization defects.
- Keep replay artifacts safe and sanitized when secrets are involved.

## Evidence

- Full sample requests, tokens/scopes used, and observed responses.
- Before/after comparisons proving unauthorized access or workflow bypass.
