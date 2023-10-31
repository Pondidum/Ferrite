package layout

import (
	"database/sql"
	"ferrite/bindings"
	"ferrite/goes/sqlite"
)

type LayoutView struct {
	Name        string                      `json:"name"`
	Keymap      Keymap                      `json:"keymap"`
	BindingSets map[string]bindings.BindSet `json:"bindingSets"`
}

func GetLayoutRaw(store *sqlite.SqliteStore, name string) (string, error) {

	var viewJson string
	err := store.Query(func(db *sql.DB) error {
		query := `select view from auto_projections where view ->> '$.name' = ?`

		if err := db.QueryRow(query, name).Scan(&viewJson); err != nil {
			return err
		}

		return nil
	})

	return viewJson, err
}

func GetLayoutNames(store *sqlite.SqliteStore) ([]string, error) {

	names := []string{}

	err := store.Query(func(db *sql.DB) error {
		query := `
			select view ->> '$.name'
			from auto_projections
			where view_type = ?
			order by view ->> '$.name'`

		rows, err := db.Query(query, sqlite.ViewType[LayoutView]())
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var name string
			if err := rows.Scan(&name); err != nil {
				return err
			}

			names = append(names, name)
		}

		return nil
	})

	return names, err
}
