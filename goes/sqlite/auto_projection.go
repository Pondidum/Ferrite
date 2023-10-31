package sqlite

import (
	"database/sql"
	"encoding/json"
	"ferrite/goes"
	"reflect"
)

func createAutoProjectionTable(db *sql.DB) error {

	createTable := `
		create table if not exists auto_projections(
			aggregate_id text primary key,
			view_type text not null,
			view text not null
		);`

	_, err := db.Exec(createTable)
	return err

}

func writeAutoProjection(tx *sql.Tx, state *goes.AggregateState) error {
	view := goes.Project(state)
	if view == nil {
		return nil
	}

	viewType := reflect.TypeOf(view).Name()
	viewJson, err := json.Marshal(view)
	if err != nil {
		return err
	}

	updateView := `
		insert into auto_projections (aggregate_id, view_type, view)
		values (?, ?, ?)
		on conflict(aggregate_id)
		do update set view=excluded.view`

	if _, err := tx.Exec(updateView, state.ID(), viewType, viewJson); err != nil {
		return err
	}

	return nil
}

func ViewById[TView any](store *SqliteStore, aggregate_id string) (*TView, error) {
	var view TView

	err := store.Query(func(db *sql.DB) error {

		query := `select view from auto_projections where aggregate_id = ?`
		viewJson := ""

		if err := db.QueryRow(query, aggregate_id).Scan(&viewJson); err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(viewJson), &view); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &view, nil
}

func ViewByProperty[TView any](store *SqliteStore, path string, value any) (*TView, error) {
	viewType := reflect.TypeOf(*new(TView)).Name()
	var view TView

	err := store.Query(func(db *sql.DB) error {

		query := `select view from auto_projections where view_type = ? and view ->> ? = ?`
		viewJson := ""

		if err := db.QueryRow(query, viewType, path, value).Scan(&viewJson); err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(viewJson), &view); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &view, nil
}

func ViewType[TView any]() string {
	return reflect.TypeOf(*new(TView)).Name()
}
