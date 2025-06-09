package gumtea

import tea "github.com/charmbracelet/bubbletea"

// Options defines configurable options for Bubble Tea programs.
type Options struct {
}

// toProgramOptions converts Options into a slice of tea.ProgramOption.
func (o *Options) toProgramOptions() []tea.ProgramOption {
	opts := []tea.ProgramOption{}
	return opts
}
