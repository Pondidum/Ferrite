package layout

import (
	"database/sql"
	"ferrite/bindings"
)

type LayoutView struct {
	Name        string                      `json:"name"`
	Keymap      Keymap                      `json:"keymap"`
	BindingSets map[string]bindings.BindSet `json:"bindingSets"`
}

// type Layouts struct{}

// func NewLayoutsProjection() *sqlite.Projection {

// 	l := &Layouts{}

// 	projection := sqlite.NewProjection("layouts", l.init)
// 	sqlite.Register(projection, l.onLayoutCreated)

// 	return projection
// }

// func (l *Layouts) init(db *sql.DB) error {

// 	createTable := `
// 	CREATE TABLE IF NOT EXISTS projection_layouts(
// 		id integer primary key autoincrement,
// 		layout_id text not null,
// 		name text not null,
// 		constraint layout_names unique(name) on conflict rollback
// 	);
// 	`

// 	_, err := db.Exec(createTable)
// 	return err
// }

// func (l *Layouts) onLayoutCreated(tx *sql.Tx, descriptor goes.EventDescriptor, event LayoutCreated) error {

// 	insert := `insert into projection_layouts (layout_id, name) values (?, ?)`

// 	_, err := tx.Exec(insert, descriptor.AggregateID, event.Name)
// 	return err
// }

// func QueryHasLayout(store *sqlite.SqliteStore, layoutName string) (bool, error) {

// 	var count sql.NullInt64
// 	if err := store.DB().QueryRow("select count(*) from projection_layouts where name = ?", layoutName).Scan(&count); err != nil {
// 		if err != sql.ErrNoRows {
// 			return false, err
// 		}
// 	}

// 	return count.Valid && count.Int64 > 0, nil
// }

// func QueryLayoutID(store *sqlite.SqliteStore, layoutName string) (uuid.UUID, error) {

// 	var id string
// 	if err := store.DB().QueryRow("select layout_id from projection_layouts where name = ?", layoutName).Scan(&id); err != nil {
// 		if err != sql.ErrNoRows {
// 			return uuid.Nil, err
// 		}
// 	}

// 	layoutId, err := uuid.Parse(id)
// 	if err != nil {
// 		return uuid.Nil, err
// 	}

// 	return layoutId, nil
// }

func GetLayoutRaw(name string, viewJson *string) func(db *sql.DB) error {

	return func(db *sql.DB) error {
		query := `select view from auto_projections where view ->> '$.name' = ?`

		if err := db.QueryRow(query, name).Scan(viewJson); err != nil {
			return err
		}

		return nil
	}
}
