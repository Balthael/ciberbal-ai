# Domain Pack — Mobile

## Objective

Guide iOS and Android assessments toward meaningful client, transport, storage, and backend interaction findings instead of device-only trivia.

## Checklist

- Review local storage, secrets handling, jailbreak/root assumptions, and certificate trust behavior.
- Validate network protections, API interaction, and app-to-backend trust boundaries.
- Distinguish what is purely local compromise from what creates broader account or server impact.

## Outputs

- Mobile-specific findings tied to user risk and backend implications.
- Notes on what required device compromise versus what worked with ordinary user context.
- Reproduction path including tooling and environment assumptions.

## Guardrails

- Do not overstate local-only issues as remote compromise.
- Keep test devices, app builds, and instrumentation setup documented.
- Preserve user data and credentials carefully during dynamic analysis.

## Evidence

- Screenshots, intercepted flows, storage artifacts, and instrumentation output.
- Device state assumptions and app version/build identifiers.
