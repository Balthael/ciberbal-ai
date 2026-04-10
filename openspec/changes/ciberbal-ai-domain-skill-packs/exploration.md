# Exploration — ciberbal-ai-domain-skill-packs

## Goal

Turn the recently defined pentesting domain architecture into practical post-install capability packs that are bundled into the unified `full-pentest` ecosystem without introducing any install-time domain branching.

## Context

The prior change (`ciberbal-ai-pentesting-domain-architecture`) already established:

- unified install flow
- no install-time AD/web/API/mobile/WiFi selector
- canonical domains
- canonical engagement phases
- `full-pentest` as the default quick-install preset

What is still missing is the **operational layer**:

- domain skill packs
- workflow/playbook packs
- the mapping between engagement phases and domains
- the documentation explaining how users leverage those packs after install

## What already exists

The repository already has:

1. **Embedded skill delivery infrastructure**
   - bundled skills are already shipped to supported agents
2. **Catalog and preset metadata**
   - domain taxonomy exists in `internal/catalog/`
3. **Docs surfaces for user guidance**
   - README + docs can already describe capability layering
4. **SDD and workflow engine**
   - suitable for documenting domain/workflow playbooks

This means the next step does not need a new install paradigm. It needs **content architecture and packaging**.

## Recommended direction

### 1) Separate packs into two axes

#### Domain packs
- Web
- API
- Mobile
- AD/Internal
- WiFi/Wireless
- Cloud
- Recon
- Reporting

#### Workflow packs
- Scoping
- Recon
- Enumeration
- Exploitation
- Post-exploitation
- Evidence
- Reporting

This keeps the product aligned with real pentesting: domain context + engagement phase context.

### 2) Keep MVP static and bundled

The MVP should define these packs as:

- bundled skills/instructions
- playbook-style guidance
- templates/checklists
- future-extensible metadata

The MVP should **not** attempt deep autonomous behavior for each domain. That would inflate scope too early.

### 3) Bundle packs under `full-pentest`

The user requirement is still the same:

- install once
- get the full stack
- no install-time specialty branching

So `full-pentest` should bundle all domain/workflow packs by default.

### 4) Make usage discoverable after install

Users should not choose domains during install, but they still need to understand what was installed.

Likely surfaces:

- docs page(s) describing each pack
- preset review/install summary language
- future agent/skill catalog visibility

## MVP vs later

### MVP now
- define the pack taxonomy
- define where pack assets live
- create initial skill/playbook pack content
- bundle into `full-pentest`
- document how packs relate to domains and phases

### Later
- richer domain-specific automation
- dynamic pack selection or filtering at runtime
- domain-specific agent generation flows
- reporting/evidence templates integrated with generated outputs

## Likely impacted areas

- `internal/assets/skills/`
- `internal/catalog/`
- `internal/components/skills/`
- `README.md`
- `docs/`

## Risks

1. **Content bloat**
   - too many assets too quickly may dilute quality
2. **Shallow packs**
   - if packs are only names with no practical guidance, the feature will feel fake
3. **Confusion between domain packs and install choices**
   - docs must clearly state these are post-install capability layers, not setup branches

## Exploration conclusion

The right next step is to create **bundled static domain and workflow packs** as post-install capability layers under `full-pentest`, starting with a lean but real MVP that gives users practical guidance across core pentesting areas without changing the install flow.
