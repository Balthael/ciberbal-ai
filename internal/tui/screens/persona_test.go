package screens

import (
	"strings"
	"testing"

	"github.com/gentleman-programming/ciberbal-ai/internal/model"
)

func TestPersonaOptionsOnlyExposeCiberbal(t *testing.T) {
	options := PersonaOptions()

	if len(options) != 1 {
		t.Fatalf("PersonaOptions() returned %d options, want 1: %v", len(options), options)
	}
	if options[0] != model.PersonaCiberbal {
		t.Fatalf("PersonaOptions()[0] = %q, want %q", options[0], model.PersonaCiberbal)
	}
}

func TestRenderPersonaUsesSingleProfessionalCiberbalPersonality(t *testing.T) {
	out := RenderPersona(model.PersonaCiberbal, 0)

	want := []string{
		"Ciberbal Personality",
		"One professional, neutral, technical mentor that teaches before it solves.",
		"Ciberbal — professional technical mentor",
	}
	for _, text := range want {
		if !strings.Contains(out, text) {
			t.Fatalf("RenderPersona() missing %q; got:\n%s", text, out)
		}
	}

	for _, old := range []string{"Choose your Persona", "(*) gentleman", "( ) gentleman", "(*) neutral", "( ) neutral", "(*) custom", "( ) custom"} {
		if strings.Contains(out, old) {
			t.Fatalf("RenderPersona() contains legacy persona option %q; got:\n%s", old, out)
		}
	}
}
