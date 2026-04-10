package catalog

import "testing"

func TestPentestingDomainsIncludesCanonicalDomains(t *testing.T) {
	domains := PentestingDomains()
	if len(domains) != 8 {
		t.Fatalf("PentestingDomains() len = %d, want 8", len(domains))
	}

	want := []string{"AD/Internal", "Mobile", "Web", "API", "WiFi/Wireless", "Cloud", "Reporting", "Recon"}
	for i, name := range want {
		if domains[i].Name != name {
			t.Fatalf("domain[%d] = %q, want %q", i, domains[i].Name, name)
		}
	}
}

func TestEngagementPhasesIncludesCanonicalPhases(t *testing.T) {
	phases := EngagementPhases()
	if len(phases) != 7 {
		t.Fatalf("EngagementPhases() len = %d, want 7", len(phases))
	}

	want := []string{"Scoping", "Recon", "Enumeration", "Exploitation", "Post-exploitation", "Evidence", "Reporting"}
	for i, name := range want {
		if phases[i].Name != name {
			t.Fatalf("phase[%d] = %q, want %q", i, phases[i].Name, name)
		}
	}
}
