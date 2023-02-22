package keyboard

type Keymap struct {
	Layers []Layer `json:"layers"`
}

type Layer struct {
	Name     string    `json:"name"`
	Bindings []Keybind `json:"bindings"`
}

type Keybind struct {
	Type  string   `json:"type"`
	Codes []string `json:"codes"`
}
