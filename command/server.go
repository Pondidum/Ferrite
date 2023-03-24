package command

import (
	"context"
	"ferrite/keyboard"
	"ferrite/zmk"
	"html/template"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
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

	engine := html.New("./ui", ".html")
	engine.Reload(true)
	engine.AddFunc("csv", func(elems []string) string {
		return strings.Join(elems, ",")
	})
	engine.AddFunc("css", func(in string) template.CSS {
		return template.CSS(in)
	})

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	kb, err := keyboard.ReadKeyboardInfo("./config/keyboard.json")
	if err != nil {
		return err
	}

	keymap, err := keyboard.ReadKeymap("./config/keymap.json")
	if err != nil {
		return err
	}

	possibleKeys, err := zmk.ReadKeys()
	if err != nil {
		return err
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"ZmkSymbolIndex": zmk.BuildSymbolIndex(possibleKeys),
			"Layout":         kb.Layout,
			"Keymap":         keymap,
		}, "layouts/main")
	})

	app.Static("/js", "./ui/js")
	app.Static("/css", "./ui/css")

	return app.Listen(c.addr)
}
