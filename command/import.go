package command

import (
	"context"
	"encoding/json"
	"ferrite/keyboard"
	"fmt"
	"io/ioutil"

	"github.com/mitchellh/cli"
	"github.com/spf13/pflag"
)

func NewImportCommand(ui cli.Ui) (*ImportCommand, error) {
	cmd := &ImportCommand{}
	cmd.Base = NewBase(ui, cmd)

	return cmd, nil
}

type ImportCommand struct {
	Base
}

func (c *ImportCommand) Name() string {
	return "import"
}

func (c *ImportCommand) Synopsis() string {
	return "import a nickcoutsos' keymap file"
}

func (c *ImportCommand) Flags() *pflag.FlagSet {
	flags := pflag.NewFlagSet(c.Name(), pflag.ContinueOnError)

	return flags
}

func (c *ImportCommand) EnvironmentVariables() map[string]string {
	return map[string]string{}
}

func (c *ImportCommand) RunContext(ctx context.Context, args []string) error {

	if len(args) != 1 {
		return fmt.Errorf("this command takes 1 argument: path to a keymap.json")
	}

	keymap, err := keyboard.Import(args[0])
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(keymap, "", "  ")
	if err != nil {
		return err
	}

	output := "./config/keymap.json"

	if err := ioutil.WriteFile(output, b, 0666); err != nil {
		return err
	}

	c.Ui.Output(fmt.Sprintf("Successfully imported to %s", output))
	return nil
}
