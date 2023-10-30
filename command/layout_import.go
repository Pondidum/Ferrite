package command

import (
	"context"
	"database/sql"
	"ferrite/goes/sqlite"
	"ferrite/layout"
	"ferrite/zmk"
	"fmt"
	"os"
	"path"
	"reflect"
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

	exists, err := hasLayoutAlready(store, layoutName)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("a layout named %s already exists", layoutName)
	}

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

func hasLayoutAlready(store *sqlite.SqliteStore, name string) (bool, error) {
	viewType := reflect.TypeOf(*new(layout.LayoutView)).Name()

	exists := false
	err := store.Query(func(db *sql.DB) error {

		query := `select count(*) from auto_projections where view_type = ? and view ->> '$.name' = ?`

		var count sql.NullInt32
		if err := db.QueryRow(query, viewType, name).Scan(&count); err != nil {
			return err
		}

		exists = count.Valid && count.Int32 > 0
		return nil

	})
	if err != nil {
		return false, err
	}

	return exists, nil

}
