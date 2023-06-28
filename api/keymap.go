package api

import "ferrite/zmk"

type Keymap struct {
	Configs []Configuration `json:"configs"`
	Combos  []Combo         `json:"combos"`
	Layers  []Layer         `json:"layers"`
}

type Configuration struct {
	Name       string           `json:"name"`
	Properties map[string]Value `json:"properties"`
}

type Value struct {
	String *string `json:"string"`
	Number *int    `json:"number"`
}

type Combo struct {
	Name         string    `json:"name"`
	Timeout      int       `json:"timeout"`
	KeyPositions []int32   `json:"keyPositions"`
	Layers       []int32   `json:"layers"`
	Bindings     []Binding `json:"bindings"`
}

type Binding struct {
	Action string      `json:"action"`
	Params []Parameter `json:"params"`
}

type Parameter struct {
	Number   *int32   `json:"number"`
	KeyCodes []string `json:"keyCodes"`
}

type Layer struct {
	Name     string    `json:"name"`
	Bindings []Binding `json:"bindings"`
}

func KeymapFromZmk(f *zmk.File) Keymap {

	km := Keymap{
		Configs: make([]Configuration, len(f.Configs)),
		Combos:  make([]Combo, len(f.Device.Combos.Combos)),
		Layers:  make([]Layer, len(f.Device.Keymap.Layers)),
	}

	for i, c := range f.Configs {
		config := Configuration{
			Name:       c.Behavior,
			Properties: make(map[string]Value, len(c.Values)),
		}

		for _, opt := range c.Values {
			config.Properties[*opt.Key] = Value{String: opt.Value.String, Number: opt.Value.Number}
		}

		km.Configs[i] = config
	}

	for i, c := range f.Device.Combos.Combos {
		km.Combos[i] = Combo{
			Name:         c.Name,
			Timeout:      int(c.Timeout),
			KeyPositions: c.KeyPositions,
			Layers:       c.Layers,
			Bindings:     convertBindings(c.Bindings),
		}
	}

	for i, layer := range f.Device.Keymap.Layers {
		km.Layers[i] = Layer{
			Name:     layer.Name,
			Bindings: convertBindings(layer.Bindings),
		}
	}

	return km
}

func convertBindings(behaviours []*zmk.Behavior) []Binding {
	bindings := make([]Binding, len(behaviours))

	for i, b := range behaviours {
		bindings[i] = Binding{
			Action: b.Action,
			Params: convertParams(b.Params),
		}
	}

	return bindings
}

func convertParams(bindings []*zmk.Binding) []Parameter {
	params := make([]Parameter, len(bindings))

	for i, b := range bindings {
		params[i] = Parameter{
			Number:   b.Number,
			KeyCodes: parseKeys(b.KeyCode),
		}
	}

	return params
}

func parseKeys(input *string) []string {
	keys := []string{}

	if input == nil {
		return keys
	}

	current := []rune{}
	for _, char := range *input {

		if char == '(' {
			keys = append(keys, string(current))
			// keys.push(current + "(code)"); // modifiers are defined as "LS(code)"
			current = []rune{}
		} else if char == ')' {
			break
		} else {
			current = append(current, char)
		}
	}

	if len(current) > 0 {
		keys = append(keys, string(current))
	}

	return keys
}
