package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestExecuteCommandQuietModeIncludesCapturedOutputOnFailure(t *testing.T) {
	restore := SetCommandOutputStreaming(false)
	defer restore()

	err := executeCommand("bash", "-c", "echo boom && exit 1")
	if err == nil {
		t.Fatal("executeCommand() error = nil, want non-nil")
	}

	if !strings.Contains(err.Error(), "boom") {
		t.Fatalf("executeCommand() error = %q, want captured output", err.Error())
	}
}

func TestSetCommandOutputStreamingRestore(t *testing.T) {
	streamCommandOutput = true
	restore := SetCommandOutputStreaming(false)

	if streamCommandOutput {
		t.Fatal("streamCommandOutput should be false after SetCommandOutputStreaming(false)")
	}

	restore()
	if !streamCommandOutput {
		t.Fatal("restore should reset streamCommandOutput to previous value")
	}
}

func TestSetCommandIORestore(t *testing.T) {
	stdin := bytes.NewBufferString("input\n")
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	previousStdin := commandStdin
	previousStdout := commandStdout
	previousStderr := commandStderr
	restore := SetCommandIO(stdin, stdout, stderr)

	if commandStdin != stdin {
		t.Fatal("commandStdin was not replaced")
	}
	if commandStdout != stdout {
		t.Fatal("commandStdout was not replaced")
	}
	if commandStderr != stderr {
		t.Fatal("commandStderr was not replaced")
	}

	restore()
	if commandStdin != previousStdin {
		t.Fatal("commandStdin was not restored")
	}
	if commandStdout != previousStdout {
		t.Fatal("commandStdout was not restored")
	}
	if commandStderr != previousStderr {
		t.Fatal("commandStderr was not restored")
	}
}

func TestExecuteCommandStreamingUsesConfiguredStdin(t *testing.T) {
	restoreStreaming := SetCommandOutputStreaming(true)
	defer restoreStreaming()

	previousStdin := commandStdin
	commandStdin = bytes.NewBufferString("ok\n")
	defer func() { commandStdin = previousStdin }()

	if err := executeCommand("bash", "-c", `read value; test "$value" = "ok"`); err != nil {
		t.Fatalf("executeCommand() error = %v, want nil", err)
	}
}

func TestExecuteCommandStreamingUsesConfiguredIO(t *testing.T) {
	restoreStreaming := SetCommandOutputStreaming(true)
	defer restoreStreaming()

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	restoreIO := SetCommandIO(bytes.NewBufferString("ok\n"), stdout, stderr)
	defer restoreIO()

	if err := executeCommand("bash", "-c", `read value; printf "%s" "$value"; printf "err" >&2`); err != nil {
		t.Fatalf("executeCommand() error = %v, want nil", err)
	}
	if got := stdout.String(); got != "ok" {
		t.Fatalf("stdout = %q, want ok", got)
	}
	if got := stderr.String(); got != "err" {
		t.Fatalf("stderr = %q, want err", got)
	}
}

func TestExecuteCommandQuietModeUsesConfiguredStdin(t *testing.T) {
	restoreStreaming := SetCommandOutputStreaming(false)
	defer restoreStreaming()

	previousStdin := commandStdin
	commandStdin = bytes.NewBufferString("ok\n")
	defer func() { commandStdin = previousStdin }()

	if err := executeCommand("bash", "-c", `read value; test "$value" = "ok"`); err != nil {
		t.Fatalf("executeCommand() error = %v, want nil", err)
	}
}

func TestCommandRequiresTerminal(t *testing.T) {
	tests := []struct {
		name    string
		command string
		args    []string
		want    bool
	}{
		{
			name:    "direct sudo command",
			command: "sudo",
			args:    []string{"apt-get", "install", "-y", "git"},
			want:    true,
		},
		{
			name:    "absolute sudo path",
			command: "/usr/bin/sudo",
			args:    []string{"apt-get", "install", "-y", "git"},
			want:    true,
		},
		{
			name:    "direct su command",
			command: "/usr/bin/su",
			args:    []string{"root", "-c", "apt-get install -y git"},
			want:    true,
		},
		{
			name:    "direct doas command",
			command: "doas",
			args:    []string{"apt-get", "install", "-y", "git"},
			want:    true,
		},
		{
			name:    "sudo inside shell pipeline",
			command: "bash",
			args:    []string{"-c", "curl -fsSL https://example.test/setup.sh | sudo -E bash -"},
			want:    true,
		},
		{
			name:    "su inside shell command",
			command: "sh",
			args:    []string{"-c", "su root -c 'apt-get install -y git'"},
			want:    true,
		},
		{
			name:    "doas inside shell command",
			command: "zsh",
			args:    []string{"-c", "doas apt-get install -y git"},
			want:    true,
		},
		{
			name:    "absolute sudo inside shell pipeline",
			command: "bash",
			args:    []string{"-c", "curl https://example.test/setup.sh | /usr/bin/sudo bash -"},
			want:    true,
		},
		{
			name:    "quiet non-interactive shell command",
			command: "bash",
			args:    []string{"-c", "echo ok"},
			want:    false,
		},
		{
			name:    "plain command with sudo in an argument",
			command: "printf",
			args:    []string{"sudo"},
			want:    false,
		},
		{
			name:    "url ending with su is not a shell command",
			command: "bash",
			args:    []string{"-c", "curl https://example.test/issues/su"},
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commandRequiresTerminal(tt.command, tt.args); got != tt.want {
				t.Fatalf("commandRequiresTerminal(%q, %v) = %v, want %v", tt.command, tt.args, got, tt.want)
			}
		})
	}
}

func TestShouldStreamCommandOutputForInteractiveCommandInQuietMode(t *testing.T) {
	restoreStreaming := SetCommandOutputStreaming(false)
	defer restoreStreaming()

	if !shouldStreamCommandOutput("sudo", []string{"apt-get", "install", "-y", "git"}) {
		t.Fatal("shouldStreamCommandOutput() = false, want true for sudo in quiet mode")
	}

	if shouldStreamCommandOutput("bash", []string{"-c", "echo ok"}) {
		t.Fatal("shouldStreamCommandOutput() = true, want false for non-interactive command in quiet mode")
	}
}
