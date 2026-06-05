package screens

import (
	"strings"
	"testing"
)

func TestRenderCompleteSuccessShowsGGANotesWhenInstalled(t *testing.T) {
	out := RenderComplete(CompletePayload{
		ConfiguredAgents:    1,
		InstalledComponents: 1,
		GGAInstalled:        true,
	})

	if !strings.Contains(out, "GGA (per project)") {
		t.Fatalf("missing GGA section: %q", out)
	}
	if !strings.Contains(out, "gga init") || !strings.Contains(out, "gga install") {
		t.Fatalf("missing GGA repo commands: %q", out)
	}
}

func TestRenderCompleteSuccessHidesGGANotesWhenNotInstalled(t *testing.T) {
	out := RenderComplete(CompletePayload{
		ConfiguredAgents:    1,
		InstalledComponents: 1,
		GGAInstalled:        false,
	})

	if strings.Contains(out, "GGA (per project)") {
		t.Fatalf("unexpected GGA section: %q", out)
	}
}

// TestRenderCompleteFailedRetryTextUsesCiberbalBranding verifies that the
// "retry" instruction in the failure path says "ciberbal-ai" not "gentle-ai".
func TestRenderCompleteFailedRetryTextUsesCiberbalBranding(t *testing.T) {
	out := RenderComplete(CompletePayload{
		FailedSteps: []FailedStep{{ID: "step-x", Error: "oops"}},
	})

	if strings.Contains(out, "gentle-ai again") {
		t.Fatalf("retry text still says old branding 'gentle-ai': %q", out)
	}
	if !strings.Contains(out, "ciberbal-ai again") {
		t.Fatalf("retry text missing 'ciberbal-ai': %q", out)
	}
}
