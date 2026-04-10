package catalog

import "github.com/gentleman-programming/gentle-ai/internal/model"

type Component struct {
	ID          model.ComponentID
	Name        string
	Description string
}

var mvpComponents = []Component{
	{ID: model.ComponentEngram, Name: "Engram", Description: "Persistent cross-session memory across pentesting engagements and projects"},
	{ID: model.ComponentSDD, Name: "SDD", Description: "Spec-driven workflow for planning, implementing, and verifying security work"},
	{ID: model.ComponentSkills, Name: "Skills", Description: "Capability-layer skill packs spanning web, API, mobile, AD/internal, wireless, cloud, and reporting workflows"},
	{ID: model.ComponentContext7, Name: "Context7", Description: "Live framework and library docs for implementation and validation during engagements"},
	{ID: model.ComponentPersona, Name: "Persona", Description: "Operator behavior mode for communication and execution style"},
	{ID: model.ComponentPermission, Name: "Permissions", Description: "Security-first defaults and guardrails for offensive-security workflows"},
	{ID: model.ComponentGGA, Name: "GGA", Description: "AI provider routing for domain-aware workflows and model switching"},
	{ID: model.ComponentTheme, Name: "Theme", Description: "Visual theme overlay for the installer and agent ecosystem"},
}

func MVPComponents() []Component {
	components := make([]Component, len(mvpComponents))
	copy(components, mvpComponents)
	return components
}
