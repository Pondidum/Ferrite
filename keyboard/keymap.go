package keyboard

import (
	"encoding/json"
	"io/ioutil"
)

type Keymap struct {
	Layers []Layer `json:"layers"`
}

type Layer struct {
	Name     string    `json:"name"`
	Bindings []Keybind `json:"bindings"`
}

type Keybind struct {
	Type      string   `json:"type"`
	FirstKey  []string `json:"first"`
	SecondKey []string `json:"second"`
}

func ReadKeymap(jsonFile string) (*Keymap, error) {
	keyboardJs, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	km := &Keymap{}
	if err := json.Unmarshal(keyboardJs, &km); err != nil {
		return nil, err
	}

	return km, nil

}
