package screens

import (
	"strings"

	"github.com/gentleman-programming/ciberbal-ai/internal/model"
	"github.com/gentleman-programming/ciberbal-ai/internal/tui/styles"
)

func PersonaOptions() []model.PersonaID {
	return []model.PersonaID{model.PersonaCiberbal}
}

func RenderPersona(selected model.PersonaID, cursor int) string {
	var b strings.Builder

	b.WriteString(styles.TitleStyle.Render("Ciberbal Personality"))
	b.WriteString("\n\n")
	b.WriteString(styles.SubtextStyle.Render("One professional, neutral, technical mentor that teaches before it solves."))
	b.WriteString("\n\n")

	for idx, persona := range PersonaOptions() {
		isSelected := persona == selected
		focused := idx == cursor
		b.WriteString(renderRadio(personaLabel(persona), isSelected, focused))
	}

	b.WriteString("\n")
	b.WriteString(renderOptions([]string{"Back"}, cursor-len(PersonaOptions())))
	b.WriteString("\n")
	b.WriteString(styles.HelpStyle.Render("j/k: navigate • enter: select • esc: back"))

	return b.String()
}

func personaLabel(persona model.PersonaID) string {
	switch persona {
	case model.PersonaCiberbal, model.PersonaGentleman, model.PersonaNeutral:
		return "Ciberbal — professional technical mentor"
	case model.PersonaCustom:
		return "Custom"
	default:
		return string(persona)
	}
}
