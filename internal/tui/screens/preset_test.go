package screens

import (
	"strings"
	"testing"

	"github.com/gentleman-programming/ciberbal-ai/internal/model"
)

func TestPresetOptionsExposePentestingPresets(t *testing.T) {
	got := PresetOptions()
	want := []model.PresetID{
		model.PresetFullPentest,
		model.PresetEcosystemCore,
		model.PresetMinimal,
		model.PresetCustom,
	}

	if len(got) != len(want) {
		t.Fatalf("PresetOptions() len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("PresetOptions()[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestRenderPresetShowsPentestingPresetLabels(t *testing.T) {
	out := RenderPreset(model.PresetFullPentest, 0)
	for _, want := range []string{"full-pentest", "ecosystem-core", "minimal", "custom"} {
		if !strings.Contains(out, want) {
			t.Fatalf("RenderPreset output missing %q; output:\n%s", want, out)
		}
	}
	if !strings.Contains(out, "Complete pentesting ecosystem") {
		t.Fatalf("RenderPreset output missing full-pentest description; output:\n%s", out)
	}
}
