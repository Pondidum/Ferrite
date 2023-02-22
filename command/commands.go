package command

import (
	"github.com/mitchellh/cli"
)

func Commands(ui cli.Ui) map[string]cli.CommandFactory {

	return map[string]cli.CommandFactory{
		"version": func() (cli.Command, error) {
			return NewVersionCommand(ui)
		},

		"import": func() (cli.Command, error) {
			return NewImportCommand(ui)
		},

		"server": func() (cli.Command, error) {
			return NewServerCommand(ui)
		},
	}
}
