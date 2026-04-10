# Domain Skill Packs Specification

## Purpose

Defines the structure, bundling, and visibility of domain-specific skill packs for the pentesting ecosystem. This ensures that domain capabilities are available post-installation without requiring the user to make mutually exclusive choices during the install process.

## Requirements

### Requirement: Bundled Domain Skill Packs

The system MUST provide bundled skill packs for specific pentesting domains. The supported domains MUST include: Web, API, Mobile, AD/Internal, WiFi/Wireless, Cloud, Recon, and Reporting.

#### Scenario: Verify domain skill packs are available

- GIVEN the pentesting ecosystem is installed
- WHEN a user attempts to use a domain-specific capability
- THEN the system MUST provide access to skills, prompts, and playbooks for Web, API, Mobile, AD/Internal, WiFi/Wireless, Cloud, Recon, and Reporting

#### Scenario: Handling unsupported domains

- GIVEN the pentesting ecosystem is installed
- WHEN a user attempts to access a domain skill pack not in the supported list (e.g., Hardware)
- THEN the system SHOULD gracefully indicate that the specific domain pack is not currently bundled

### Requirement: Default Bundling in full-pentest

The system MUST bundle all domain skill packs into the `full-pentest` preset by default, ensuring they are automatically deployed during a standard installation.

#### Scenario: Installing full-pentest preset

- GIVEN a user installs the system using the `full-pentest` preset
- WHEN the installation completes
- THEN all domain skill packs MUST be present and ready for use
- AND the system MUST NOT prompt the user to select specific domains during the installation

### Requirement: Post-Install Discovery and Visibility

The system MUST provide documentation and visibility mechanisms for users to discover the installed domain skill packs after installation.

#### Scenario: Discovering installed capabilities

- GIVEN the `full-pentest` preset has been installed
- WHEN the user consults the documentation or capability visibility surfaces
- THEN the system MUST clearly list the available domain skill packs and explain how to leverage them post-install
