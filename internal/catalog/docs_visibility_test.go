package catalog

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPentestingDomainsDocListsBundledDomainPacks(t *testing.T) {
	content, err := os.ReadFile(filepath.Join("..", "..", "docs", "pentesting-domains.md"))
	if err != nil {
		t.Fatalf("ReadFile(pentesting-domains.md) error = %v", err)
	}

	text := string(content)
	for _, want := range []string{
		"## Bundled packs in `full-pentest`",
		"8 domain packs",
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("pentesting-domains.md missing %q", want)
		}
	}

	for _, domain := range []string{"Web", "API", "Mobile", "AD/Internal", "WiFi/Wireless", "Cloud", "Recon", "Reporting"} {
		if !strings.Contains(text, domain) {
			t.Fatalf("pentesting-domains.md missing domain %q", domain)
		}
	}
}

func TestPentestingDomainsDocListsWorkflowPhaseGuidance(t *testing.T) {
	content, err := os.ReadFile(filepath.Join("..", "..", "docs", "pentesting-domains.md"))
	if err != nil {
		t.Fatalf("ReadFile(pentesting-domains.md) error = %v", err)
	}

	text := string(content)
	for _, want := range []string{
		"## Canonical engagement phases",
		"Scoping",
		"Enumeration",
		"Post-exploitation",
		"7 workflow packs",
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("pentesting-domains.md missing %q", want)
		}
	}
}
