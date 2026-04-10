# installer-flow Specification

## Purpose

Defines the dual-mode (Quick/Advanced) installation flow initiated from the main welcome menu.

## Requirements

### Requirement: Welcome Menu Entry Point

The system MUST present "Start installation" as the primary entry point in the main menu for unconfigured or partially configured systems.
The system MUST preserve all existing core capabilities (memory, SDD, model config, backups, sync, profiles, create-your-own-agent) in the main menu structure.

#### Scenario: User initiates installation
- GIVEN the user is on the main welcome menu
- WHEN the user selects "Start installation"
- THEN the system MUST proceed to the install-mode decision screen
- AND the existing menu capabilities MUST remain accessible if the user cancels or completes installation

### Requirement: Install-Mode Decision Screen

The system MUST present a choice between "Quick install" and "Advanced install" after environment detection in the installation flow.

#### Scenario: User reaches decision screen
- GIVEN the user selected "Start installation"
- WHEN the environment detection completes
- THEN the system MUST display the options "Quick install" and "Advanced install"

### Requirement: Quick Install Execution

The "Quick install" option MUST automatically install a full pentesting/security-ready stack by default.
The "Quick install" option MUST NOT prompt the user to choose a pentesting specialty.

#### Scenario: User selects Quick install
- GIVEN the user is on the install-mode decision screen
- WHEN the user selects "Quick install"
- THEN the system MUST install the default pentesting stack without further specialty prompts
- AND the system MUST return to the main menu upon successful completion

### Requirement: Advanced Install Execution

The "Advanced install" option MUST preserve the granular configuration behavior of the previous installation process.

#### Scenario: User selects Advanced install
- GIVEN the user is on the install-mode decision screen
- WHEN the user selects "Advanced install"
- THEN the system MUST prompt the user for granular configuration choices (e.g., specialty, specific tools)
- AND the system MUST return to the main menu upon successful completion