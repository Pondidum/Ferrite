package command

import (
	"context"
	"ferrite/api"
	"ferrite/goes/sqlite"

	"github.com/mitchellh/cli"
	"github.com/spf13/pflag"
)

func NewServerCommand(ui cli.Ui) (*ServerCommand, error) {
	cmd := &ServerCommand{
		addr: "localhost:5656",
	}
	cmd.Base = NewBase(ui, cmd)

	return cmd, nil
}

type ServerCommand struct {
	Base

	addr string
}

func (c *ServerCommand) Name() string {
	return "server"
}

func (c *ServerCommand) Synopsis() string {
	return "Runs the API and site"
}

func (c *ServerCommand) Flags() *pflag.FlagSet {
	flags := pflag.NewFlagSet(c.Name(), pflag.ContinueOnError)
	return flags
}

func (c *ServerCommand) EnvironmentVariables() map[string]string {
	return map[string]string{}
}

func (c *ServerCommand) RunContext(ctx context.Context, args []string) error {

	store, err := sqlite.CreateStore()
	if err != nil {
		return err
	}

	app, err := api.NewApiV2(store)
	if err != nil {
		return err
	}

	return app.Listen(c.addr)
}
