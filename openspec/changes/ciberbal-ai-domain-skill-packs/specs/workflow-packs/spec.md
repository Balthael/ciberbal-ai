# Workflow Packs Specification

## Purpose

Defines the phase-specific workflow packs that guide engagement execution across different pentesting phases, providing playbooks, templates, and skills tailored to each phase.

## Requirements

### Requirement: Bundled Workflow Packs

The system MUST provide bundled workflow packs aligned with standard engagement phases. The supported phases MUST include: Scoping, Recon, Enumeration, Exploitation, Post-exploitation, Evidence, and Reporting.

#### Scenario: Execute phase-specific workflow

- GIVEN the pentesting ecosystem is installed
- WHEN a user engages in a specific pentesting phase
- THEN the system MUST provide workflow guidance, skills, and playbooks for Scoping, Recon, Enumeration, Exploitation, Post-exploitation, Evidence, and Reporting

#### Scenario: Missing phase execution

- GIVEN the pentesting ecosystem is installed
- WHEN a user tries to invoke a workflow phase that does not exist
- THEN the system SHOULD return an error indicating the phase is unrecognized

### Requirement: Default Bundling in full-pentest

The system MUST bundle all workflow packs into the `full-pentest` preset by default, alongside the domain skill packs.

#### Scenario: Workflow availability post-install

- GIVEN the system is installed via the `full-pentest` preset
- WHEN the user begins an engagement
- THEN all workflow phase packs MUST be automatically available
- AND the installation process MUST NOT require manual selection of workflow packs

### Requirement: Post-Install Discovery

The system MUST document the available workflow packs and provide visibility into how they integrate with domain skill packs during an engagement.

#### Scenario: Understanding phase guidance

- GIVEN the ecosystem is installed
- WHEN the user reviews the documentation
- THEN the documentation MUST detail the workflow packs and their expected usage mapped to engagement phases
