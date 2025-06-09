package gumtea

import tea "github.com/charmbracelet/bubbletea"

// NewProgram wraps tea.NewProgram, injecting both baseOpts and options from o.
func NewProgram(model tea.Model, o Options, baseOpts ...tea.ProgramOption) *tea.Program {
	opts := append(baseOpts, o.toProgramOptions()...)
	return tea.NewProgram(model, opts...)
}
