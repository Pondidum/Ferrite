package layout

import (
	"ferrite/bindings"
	"ferrite/goes"
	"ferrite/zmk"
	"fmt"

	"github.com/google/uuid"
)

func BlankLayout() *Layout {
	l := &Layout{
		State:       goes.NewAggregateState(),
		bindingSets: map[string]bindings.BindSet{},
	}

	goes.Register(l.State, l.onLayoutCreated)
	goes.Register(l.State, l.onLayoutImported)
	goes.Register(l.State, l.onBindingChanged)

	goes.RegisterAutoProjection(l.State, l.toView)

	return l
}

func CreateLayout(name string) *Layout {
	layout := BlankLayout()
	goes.Apply(layout.State, LayoutCreated{
		ID:   uuid.New(),
		Name: name,
	})
	return layout
}

type Layout struct {
	State *goes.AggregateState

	Name string

	keymap      Keymap
	bindingSets map[string]bindings.BindSet
}

func (l *Layout) toView() LayoutView {
	return LayoutView{
		Name:        l.Name,
		Keymap:      l.keymap,
		BindingSets: l.bindingSets,
	}
}

func (l *Layout) onLayoutCreated(e LayoutCreated) {
	goes.SetID(l.State, e.ID)
	l.Name = e.Name
}

func (l *Layout) ImportFrom(file *zmk.File) error {
	// a zmk.File doesn't serialize, but our object model does!
	mapper := ZmkMapper{}
	km, bs := mapper.KeymapFromZmk(file)

	bindingSets := make([]string, len(bs))
	for i, set := range bs {
		bindingSets[i] = set.Name
	}

	return goes.Apply(l.State, LayoutImported{
		Keymap:      km,
		BindingSets: bindingSets,
	})
}

func (l *Layout) onLayoutImported(e LayoutImported) {
	l.keymap = e.Keymap

	for _, name := range e.BindingSets {
		if _, found := l.bindingSets[name]; !found {
			set, _ := bindings.SetFromName(name)
			l.bindingSets[set.Name] = set
		}
	}
}

func (l *Layout) BindKey(layerIndex int, key int, bind Binding) error {
	if layerIndex < 0 || layerIndex >= len(l.keymap.Layers) {
		return fmt.Errorf("invalid layer, valid range is 0 to %v", len(l.keymap.Layers)-1)
	}

	layer := l.keymap.Layers[layerIndex]
	if key < 0 || key >= len(layer.Bindings) {
		return fmt.Errorf("invalid key index, valid range is 0 to %v", len(layer.Bindings)-1)
	}

	e := BindingChanged{
		Layer:  layerIndex,
		Key:    key,
		Action: bind.Action,
		Params: make([]Parameter, len(bind.Params)),
	}

	for i, p := range bind.Params {
		param := Parameter{
			Number: p.Number,
		}

		if p.KeyCode != nil {
			code := canonicalKey(l.bindingSets, *p.KeyCode)
			param.KeyCode = &code
		}

		for _, key := range p.Modifiers {
			param.Modifiers = append(param.Modifiers, canonicalKey(l.bindingSets, key))
		}

		e.Params[i] = param
	}

	return goes.Apply(l.State, e)
}

func (l *Layout) onBindingChanged(e BindingChanged) {
	l.keymap.Layers[e.Layer].Bindings[e.Key] = Binding{
		Action: e.Action,
		Params: e.Params,
	}
}

func canonicalKey(sets map[string]bindings.BindSet, key string) string {
	for _, set := range sets {
		if k, found := set.CanonicalKey(key); found {
			return k
		}
	}

	return key
}

type LayoutCreated struct {
	ID   uuid.UUID
	Name string
}

type LayoutImported struct {
	Keymap      Keymap
	BindingSets []string
}

type BindingChanged struct {
	Layer  int
	Key    int
	Action string
	Params []Parameter
}
