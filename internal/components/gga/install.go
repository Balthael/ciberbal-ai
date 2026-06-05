package gga

import (
	"github.com/gentleman-programming/ciberbal-ai/internal/installcmd"
	"github.com/gentleman-programming/ciberbal-ai/internal/model"
	"github.com/gentleman-programming/ciberbal-ai/internal/system"
)

func InstallCommand(profile system.PlatformProfile) ([][]string, error) {
	return installcmd.NewResolver().ResolveComponentInstall(profile, model.ComponentGGA)
}

func ShouldInstall(enabled bool) bool {
	return enabled
}
