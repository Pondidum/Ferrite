package sqlite

import (
	"context"
	"database/sql"
	"encoding/json"
	"ferrite/goes"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteStore struct {
	db *sql.DB

	projections []*Projection
}

func CreateStore() (*SqliteStore, error) {

	db, err := sql.Open("sqlite3", "ferrite.db")
	if err != nil {
		return nil, err
	}

	createTables := `
CREATE TABLE IF NOT EXISTS events(
	event_id integer primary key autoincrement,
	aggregate_id text not null,
	sequence integer not null,
	timestamp timestamp not null,
	event_type text not null,
	event_data text not null,
	constraint aggregate_sequence unique(aggregate_id, sequence) on conflict rollback
);

create table if not exists auto_projections(
	aggregate_id text primary key,
	view text not null
);
`

	_, err = db.Exec(createTables)
	if err != nil {
		return nil, err
	}

	return &SqliteStore{
		db: db,
	}, nil
}

func (store *SqliteStore) Close() error {
	return store.db.Close()
}

func (store *SqliteStore) DB() *sql.DB {
	return store.db
}

func (store *SqliteStore) RegisterProjection(p *Projection) error {

	if err := p.init(store.db); err != nil {
		return err
	}

	store.projections = append(store.projections, p)

	return nil
}

func (store *SqliteStore) Save(state *goes.AggregateState) error {

	tx, err := store.db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var dbSequence sql.NullInt64
	if err := tx.QueryRow("select max(sequence) from events where aggregate_id = ?", state.ID()).Scan(&dbSequence); err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}

	sequence := goes.Sequence(state)

	if dbSequence.Valid && dbSequence.Int64 > int64(sequence) {
		return fmt.Errorf("aggregate has new events in the database. db: %v, memory: %v", dbSequence, sequence)
	}

	insert := `insert into events (aggregate_id, sequence, timestamp, event_type, event_data) values (?, ?, ?, ?, ?)`

	err = goes.SaveEvents(state, func(e goes.EventDescriptor) error {

		eventJson, err := json.Marshal(e.Event)
		if err != nil {
			return err
		}

		if _, err := tx.Exec(insert, e.AggregateID, e.Sequence, e.Timestamp, e.EventType, eventJson); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	if view := goes.Project(state); view != nil {

		viewJson, err := json.Marshal(view)
		if err != nil {
			return err
		}

		updateView := `insert into auto_projections (aggregate_id, view) values (?, ?) on conflict(aggregate_id) do update set view=excluded.view`
		if _, err := tx.Exec(updateView, state.ID(), viewJson); err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (store *SqliteStore) Load(state *goes.AggregateState, id uuid.UUID) error {

	rows, err := store.db.Query("select sequence, timestamp, event_type, event_data from events where aggregate_id = ?", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	events := []goes.EventDescriptor{}
	for rows.Next() {

		e := goes.EventDescriptor{
			AggregateID: state.ID(),
		}

		var eventJson []byte

		if err := rows.Scan(&e.Sequence, &e.Timestamp, &e.EventType, &eventJson); err != nil {
			return err
		}

		if e.Event, err = goes.NewEvent(state, e.EventType); err != nil {
			return err
		}

		if err := json.Unmarshal(eventJson, &e.Event); err != nil {
			return err
		}

		events = append(events, e)
	}

	return goes.LoadEvents(state, events)
}

func (store *SqliteStore) Query(query func(db *sql.DB) error) error {
	return query(store.db)
}
