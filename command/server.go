package command

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"

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

	keyboardJs, err := ioutil.ReadFile("./config/keyboard.json")
	if err != nil {
		return err
	}

	kb := map[string][]Key{}
	if err := json.Unmarshal(keyboardJs, &kb); err != nil {
		return err
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":  "Hello, World!",
			"Layout": kb["layout"],
		}, "layouts/main")
	})

	app.Static("/js", "./ui/js")
	app.Static("/css", "./ui/css")

	return app.Listen(c.addr)
}

type Key struct {
	Label string
	Row   int
	Col   int
	X     float64
	Y     float64
	R     float64
	Rx    float64
	Ry    float64
}

const DefaultSize = 65
const DefaultPadding = 5

func (k *Key) Style() string {
	x := k.X * (DefaultSize + DefaultPadding)
	y := k.Y * (DefaultSize + DefaultPadding)
	u := DefaultSize
	h := DefaultSize
	rx := (k.X - math.Max(k.Rx, k.X)) * -(DefaultSize + DefaultPadding)
	ry := (k.Y - math.Max(k.Ry, k.Y)) * -(DefaultSize + DefaultPadding)
	a := math.Max(k.R, 0)

	return fmt.Sprintf(
		"top: %vpx; left: %vpx; width: %vpx; height: %vpx; transform-origin: %vpx %vpx; transform: rotate(%vdeg)",
		y, x, u, h, rx, ry, a)

}
