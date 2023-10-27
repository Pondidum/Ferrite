package sqlite

import (
	"database/sql"
	"ferrite/goes"
	"fmt"
	"reflect"
)

type Projection struct {
	name     string
	init     func(db *sql.DB) error
	handlers map[string]func(tx *sql.Tx, descriptor goes.EventDescriptor) error
}

func NewProjection(name string, init func(db *sql.DB) error) *Projection {
	return &Projection{
		name:     name,
		init:     init,
		handlers: map[string]func(tx *sql.Tx, descriptor goes.EventDescriptor) error{},
	}
}

func Project(tx *sql.Tx, state *Projection, descriptor goes.EventDescriptor) error {

	handler, found := state.handlers[descriptor.EventType]
	if !found {
		// projections ignore events they don't know about
		return nil
	}

	if err := handler(tx, descriptor); err != nil {
		return err
	}

	return nil
}

func Register[TEvent any](projection *Projection, handler func(tx *sql.Tx, descriptor goes.EventDescriptor, event TEvent) error) {
	name := reflect.TypeOf(*new(TEvent)).Name()

	projection.handlers[name] = func(tx *sql.Tx, descriptor goes.EventDescriptor) error {

		switch e := descriptor.Event.(type) {
		case TEvent:
			return handler(tx, descriptor, e)

		case *TEvent:
			return handler(tx, descriptor, *e)

		default:
			return fmt.Errorf("unable to handle %T", e)
		}
	}

}
