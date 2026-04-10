package catalog

import (
	"strings"
	"testing"
)

func TestMVPComponentsReturnsCopy(t *testing.T) {
	first := MVPComponents()
	second := MVPComponents()

	if len(first) == 0 || len(second) == 0 {
		t.Fatal("MVPComponents() returned empty slices")
	}

	first[0].Name = "Mutated"
	if second[0].Name == "Mutated" {
		t.Fatal("MVPComponents() should return a copy, but mutation leaked")
	}
}

func TestMVPComponentsDescribePentestingCapabilityLayers(t *testing.T) {
	components := MVPComponents()

	joined := make([]string, 0, len(components))
	for _, component := range components {
		joined = append(joined, component.Description)
	}
	allDescriptions := strings.Join(joined, " ")

	for _, want := range []string{"pentesting", "web", "API", "mobile", "reporting"} {
		if !strings.Contains(allDescriptions, want) {
			t.Fatalf("component descriptions missing %q: %s", want, allDescriptions)
		}
	}
}
