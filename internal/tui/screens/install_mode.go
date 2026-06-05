package screens

import (
	"strings"

	"github.com/gentleman-programming/ciberbal-ai/internal/tui/styles"
)

// InstallModeOptions returns the two install-mode choices presented after
// environment detection. Index 0 is Quick, index 1 is Advanced.
func InstallModeOptions() []string {
	return []string{
		"Quick install",
		"Advanced install",
	}
}

// RenderInstallMode renders the install-mode decision screen.
// cursor=0 highlights "Quick install"; cursor=1 highlights "Advanced install".
func RenderInstallMode(cursor int) string {
	var b strings.Builder

	b.WriteString(styles.TitleStyle.Render("How would you like to install?"))
	b.WriteString("\n\n")

	b.WriteString(renderOptions(InstallModeOptions(), cursor))

	b.WriteString("\n")
	b.WriteString(styles.SubtextStyle.Render("Quick install: full pentesting stack with sensible defaults"))
	b.WriteString("\n")
	b.WriteString(styles.SubtextStyle.Render("Advanced install: choose agents, persona, preset and components"))
	b.WriteString("\n\n")
	b.WriteString(styles.HelpStyle.Render("j/k: navigate • enter: select • esc: back"))

	return b.String()
}
