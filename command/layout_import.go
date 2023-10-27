package command

import (
	"context"
	"ferrite/goes/sqlite"
	"ferrite/layout"
	"ferrite/zmk"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/spf13/pflag"
)

func NewLayoutImportCommand(ui cli.Ui) (*LayoutImportCommand, error) {
	cmd := &LayoutImportCommand{}
	cmd.Base = NewBase(ui, cmd)

	return cmd, nil
}

type LayoutImportCommand struct {
	Base

	name string
}

func (c *LayoutImportCommand) Name() string {
	return "layout import"
}

func (c *LayoutImportCommand) Synopsis() string {
	return "Import a layout file"
}

func (c *LayoutImportCommand) Flags() *pflag.FlagSet {
	flags := pflag.NewFlagSet(c.Name(), pflag.ContinueOnError)

	flags.StringVar(&c.name, "name", "", "specify a name for the layout")

	return flags
}

func (c *LayoutImportCommand) EnvironmentVariables() map[string]string {
	return map[string]string{}
}

func (c *LayoutImportCommand) RunContext(ctx context.Context, args []string) error {

	p := args[0]
	r, err := os.Open(p)
	if err != nil {
		return err
	}

	f, err := zmk.Parse(r)
	if err != nil {
		return err
	}

	layoutName := c.name
	if layoutName == "" {
		layoutName = path.Base(strings.TrimSuffix(p, path.Ext(p)))
	}

	// protect against duplicate layout names later
	c.Ui.Output(fmt.Sprintf("Creating new layout '%s'", layoutName))

	store, err := sqlite.CreateStore()
	if err != nil {
		return err
	}
	defer store.Close()

	// store.RegisterProjection(layout.NewLayoutsProjection())

	// exists, err := layout.QueryHasLayout(store, layoutName)
	// if err != nil {
	// 	return err
	// }

	// if exists {
	// 	return fmt.Errorf("a layout called %s already exists", layoutName)
	// }

	l := layout.CreateLayout(layoutName)

	c.Ui.Info("Importing existing layout file")
	if err := l.ImportFrom(f); err != nil {
		return err
	}

	if err := store.Save(l.State); err != nil {
		return err
	}

	c.Ui.Info("Done.")
	return nil
}
