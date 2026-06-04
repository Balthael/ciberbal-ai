package app

import (
	"fmt"
	"io"
)

func printHelp(w io.Writer, version string) {
	fmt.Fprintf(w, `ciberbal-ai — AI Gentle Stack (%s)

USAGE
  ciberbal-ai                     Launch interactive TUI
  ciberbal-ai <command> [flags]

COMMANDS
  install             Configure AI coding agents on this machine
  sync                Sync agent configs and skills to current version
  update              Check for available updates
  upgrade             Apply updates to managed tools
  restore             Restore a config backup
  skill-registry      Manage the local skill registry index
  version             Print version

FLAGS
  --help, -h    Show this help

Run 'ciberbal-ai help' for this message.
Documentation: https://github.com/Gentleman-Programming/gentle-ai
`, version)
}
