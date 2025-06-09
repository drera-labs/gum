package gumtea

import tea "github.com/charmbracelet/bubbletea"

// Options defines configurable options for Bubble Tea programs.
type Options struct {
	// FPS sets the frames per second; if 0, default FPS is used.
	FPS int `help:"Set the frames per second" env:"FPS" default:"0"`
}

// toProgramOptions converts Options into a slice of tea.ProgramOption.
func (o *Options) toProgramOptions() []tea.ProgramOption {
	opts := []tea.ProgramOption{}
	if o.FPS > 0 {
		opts = append(opts, tea.WithFPS(o.FPS))
	}
	return opts
}
