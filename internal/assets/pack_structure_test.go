package assets

import (
	"fmt"
	"strings"
	"testing"
)

var requiredPackHeadings = []string{
	"## Objective",
	"## Checklist",
	"## Outputs",
	"## Guardrails",
	"## Evidence",
}

func TestBundledPentestingPacksContainRequiredSections(t *testing.T) {
	t.Parallel()

	paths := []string{
		"skills/domain-web/SKILL.md",
		"skills/domain-api/SKILL.md",
		"skills/domain-mobile/SKILL.md",
		"skills/domain-ad-internal/SKILL.md",
		"skills/domain-wifi-wireless/SKILL.md",
		"skills/domain-cloud/SKILL.md",
		"skills/domain-recon/SKILL.md",
		"skills/domain-reporting/SKILL.md",
		"skills/workflow-scoping/SKILL.md",
		"skills/workflow-recon/SKILL.md",
		"skills/workflow-enumeration/SKILL.md",
		"skills/workflow-exploitation/SKILL.md",
		"skills/workflow-post-exploitation/SKILL.md",
		"skills/workflow-evidence/SKILL.md",
		"skills/workflow-reporting/SKILL.md",
	}

	for _, path := range paths {
		path := path
		t.Run(path, func(t *testing.T) {
			content, err := Read(path)
			if err != nil {
				t.Fatalf("Read(%q) error = %v", path, err)
			}

			for _, heading := range requiredPackHeadings {
				assertPackHasHeading(t, path, content, heading)
			}
		})
	}
}

func TestMissingMandatorySectionsInDomainPackAreFlagged(t *testing.T) {
	content := `# Domain Pack — Web

## Objective
Goal

## Checklist
- one

## Outputs
- one

## Evidence
- one`

	missing := missingRequiredHeadings(content)
	if len(missing) != 1 || missing[0] != "## Guardrails" {
		t.Fatalf("missingRequiredHeadings(domain) = %v, want [## Guardrails]", missing)
	}
}

func TestMissingMandatorySectionsInWorkflowPackAreFlagged(t *testing.T) {
	content := `# Workflow Pack — Recon

## Objective
Goal

## Checklist
- one

## Outputs
- one

## Guardrails
- one`

	missing := missingRequiredHeadings(content)
	if len(missing) != 1 || missing[0] != "## Evidence" {
		t.Fatalf("missingRequiredHeadings(workflow) = %v, want [## Evidence]", missing)
	}
}

func assertPackHasHeading(t *testing.T, path, content, heading string) {
	t.Helper()
	if !strings.Contains(content, heading) {
		t.Fatalf("%s missing required heading %s\n%s", path, heading, previewPack(content))
	}
}

func missingRequiredHeadings(content string) []string {
	missing := make([]string, 0)
	for _, heading := range requiredPackHeadings {
		if !strings.Contains(content, heading) {
			missing = append(missing, heading)
		}
	}
	return missing
}

func previewPack(content string) string {
	trimmed := strings.TrimSpace(content)
	if len(trimmed) > 220 {
		trimmed = trimmed[:220] + "..."
	}
	return fmt.Sprintf("pack preview: %q", trimmed)
}
