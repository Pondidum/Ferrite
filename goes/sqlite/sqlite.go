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
	projections []*Projection
}

func CreateStore() (*SqliteStore, error) {

	store := &SqliteStore{}

	db, err := store.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	createTables := `
CREATE TABLE IF NOT EXISTS events(
	event_id integer primary key autoincrement,
	aggregate_id text not null,
	sequence integer not null,
	timestamp timestamp not null,
	event_type text not null,
	event_data text not null,
	constraint aggregate_sequence unique(aggregate_id, sequence) on conflict rollback
);`
	_, err = db.Exec(createTables)
	if err != nil {
		return nil, err
	}

	if err := createAutoProjectionTable(db); err != nil {
		return nil, err
	}

	return store, nil
}

func (store *SqliteStore) open() (*sql.DB, error) {
	return sql.Open("sqlite3", "ferrite.db")
}

func (store *SqliteStore) RegisterProjection(p *Projection) error {

	db, err := store.open()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := p.init(db); err != nil {
		return err
	}

	store.projections = append(store.projections, p)

	return nil
}

func (store *SqliteStore) Save(state *goes.AggregateState) error {

	db, err := store.open()
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.BeginTx(context.Background(), nil)
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

	if err := writeAutoProjection(tx, state); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (store *SqliteStore) Load(state *goes.AggregateState, id uuid.UUID) error {

	db, err := store.open()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("select sequence, timestamp, event_type, event_data from events where aggregate_id = ?", id)
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

	db, err := store.open()
	if err != nil {
		return err
	}
	defer db.Close()

	return query(db)
}
