package command

import (
	"context"
	"encoding/json"
	"ferrite/goes/sqlite"
	"ferrite/layout"

	"github.com/mitchellh/cli"
	"github.com/spf13/pflag"
)

func NewLayoutShowCommand(ui cli.Ui) (*LayoutShowCommand, error) {
	cmd := &LayoutShowCommand{}
	cmd.Base = NewBase(ui, cmd)

	return cmd, nil
}

type LayoutShowCommand struct {
	Base
}

func (c *LayoutShowCommand) Name() string {
	return "layout show"
}

func (c *LayoutShowCommand) Synopsis() string {
	return "View a layout"
}

func (c *LayoutShowCommand) Flags() *pflag.FlagSet {
	flags := pflag.NewFlagSet(c.Name(), pflag.ContinueOnError)

	return flags
}

func (c *LayoutShowCommand) EnvironmentVariables() map[string]string {
	return map[string]string{}
}

func (c *LayoutShowCommand) RunContext(ctx context.Context, args []string) error {

	layoutName := args[0]

	store, err := sqlite.CreateStore()
	if err != nil {
		return err
	}
	defer store.Close()

	view, err := sqlite.ViewByProperty[layout.LayoutView](store, "$.name", layoutName)
	if err != nil {
		return err
	}

	js, err := json.MarshalIndent(view, "", "  ")
	if err != nil {
		return err
	}

	c.Ui.Output(string(js))
	return nil
}
