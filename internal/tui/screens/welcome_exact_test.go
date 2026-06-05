package screens_test

import (
	"reflect"
	"testing"

	"github.com/gentleman-programming/ciberbal-ai/internal/tui/screens"
)

func TestWelcomeOptionsMatchExpectedMainMenuWithoutProfiles(t *testing.T) {
	got := screens.WelcomeOptions(nil, true, false, 0, true)
	want := []string{
		"Start installation",
		"Upgrade tools (up to date)",
		"Sync configs",
		"Upgrade + Sync",
		"Configure models",
		"Create your own Agent",
		"Manage backups",
		"Quit",
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("WelcomeOptions() = %#v, want %#v", got, want)
	}
}

func TestWelcomeOptionsMatchExpectedMainMenuWithProfiles(t *testing.T) {
	got := screens.WelcomeOptions(nil, true, true, 0, true)
	want := []string{
		"Start installation",
		"Upgrade tools (up to date)",
		"Sync configs",
		"Upgrade + Sync",
		"Configure models",
		"Create your own Agent",
		"Engagement Profiles",
		"Manage backups",
		"Quit",
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("WelcomeOptions() = %#v, want %#v", got, want)
	}
}
