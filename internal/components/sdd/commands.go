package sdd

type OpenCodeCommand struct {
	Name        string
	Description string
	Body        string
}

func OpenCodeCommands() []OpenCodeCommand {
	return []OpenCodeCommand{
		{Name: "sdd-init", Description: "Initialize engagement context", Body: "/sdd-init"},
		{Name: "sdd-new", Description: "Start a new engagement workflow", Body: "/sdd-new ${engagement-name}"},
		{Name: "sdd-continue", Description: "Continue next engagement phase", Body: "/sdd-continue ${engagement-name}"},
		{Name: "sdd-ff", Description: "Fast-forward engagement planning phases", Body: "/sdd-ff ${engagement-name}"},
		{Name: "sdd-apply", Description: "Execute exploitation tasks", Body: "/sdd-apply ${engagement-name}"},
		{Name: "sdd-verify", Description: "Review evidence and findings", Body: "/sdd-verify ${engagement-name}"},
		{Name: "sdd-archive", Description: "Archive engagement and compile report", Body: "/sdd-archive ${engagement-name}"},
		{Name: "sdd-onboard", Description: "Guided engagement walkthrough", Body: "/sdd-onboard"},
	}
}
