package api

import "encoding/json"

type File struct {
	Includes []*Include `json:"includes"`
	Defines  []*Define  `json:"defines"`
	Configs  []*Config  `json:"configs"`
	Device   *Device    `json:"device"`
}

type Include struct {
	Value string `json:"value"`
}

type Define struct {
	Value string `json:"value"`
}

type Config struct {
	Behavior string     `json:"behavior"`
	Values   []*Options `json:"values"`
}

type Options struct {
	Key   *string `json:"key"`
	Value *Value  `json:"value"`
}

type Value struct {
	String *string `json:"string"`
	Number *int    `json:"number"`
}

type Device struct {
	Combos *Combos `json:"combos"`
	Keymap *Keymap `json:"keymap"`
}

type Combos struct {
	Compatible string   `json:"compatible"`
	Combos     []*Combo `json:"combos"`
}

type Combo struct {
	Name         string      `json:"name"`
	Timeout      int32       `json:"timeout"`
	KeyPositions []*List     `json:"keyPositions"`
	Layers       []*List     `json:"layers"`
	Bindings     []*Behavior `json:"bindings"`
}

type Keymap struct {
	Compatible string   `json:"compatible"`
	Layers     []*Layer `json:"layers"`
}

type Layer struct {
	Name           string      `json:"name"`
	Bindings       []*Behavior `json:"bindings"`
	SensorBindings []*Behavior `json:"sensorBindings"`
	EndBrace       string      `json:"endBrace"`
}

type List struct {
	Number  *int32   `json:"number"`
	KeyCode []string `json:"keyCode"`
}

func (l *List) MarshalJSON() ([]byte, error) {

	type Dto List

	dto := Dto{
		Number:  l.Number,
		KeyCode: l.KeyCode,
	}

	if dto.KeyCode == nil {
		dto.KeyCode = []string{}
	}

	return json.Marshal(dto)
}

type Behavior struct {
	Action string  `json:"action"`
	Params []*List `json:"params"`
}

func (b *Behavior) MarshalJSON() ([]byte, error) {

	type Dto Behavior

	dto := Dto{
		Action: b.Action,
		Params: b.Params,
	}

	if dto.Params == nil {
		dto.Params = []*List{}
	}

	return json.Marshal(dto)
}
