package command

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet"
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

	engine := jet.New("./ui", ".jet")
	engine.Reload(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})

	app.Static("/js", "./ui/js")
	app.Static("/css", "./ui/css")

	return app.Listen(c.addr)
}
