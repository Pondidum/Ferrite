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
	Number    *int32   `json:"number"`
	KeyCode   *string  `json:"keyCode"`
	Modifiers []string `json:"modifiers"`
}

type Layer struct {
	Name     string    `json:"name"`
	Bindings []Binding `json:"bindings"`
}

func KeymapFromZmk(zmkKeys map[string]zmk.KeyCode, f *zmk.File) Keymap {

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
			Bindings:     convertBindings(zmkKeys, c.Bindings),
		}
	}

	for i, layer := range f.Device.Keymap.Layers {
		km.Layers[i] = Layer{
			Name:     layer.Name,
			Bindings: convertBindings(zmkKeys, layer.Bindings),
		}
	}

	return km
}

func convertBindings(zmkKeys map[string]zmk.KeyCode, behaviours []*zmk.Behavior) []Binding {
	bindings := make([]Binding, len(behaviours))

	for i, b := range behaviours {
		bindings[i] = Binding{
			Action: b.Action,
			Params: convertParams(zmkKeys, b.Params),
		}
	}

	return bindings
}

func convertParams(zmkKeys map[string]zmk.KeyCode, bindings []*zmk.Binding) []Parameter {
	params := make([]Parameter, len(bindings))

	for i, b := range bindings {

		key, modifiers := parseKeys(b.KeyCode)
		param := Parameter{
			Number:    b.Number,
			Modifiers: canonical(zmkKeys, modifiers),
		}

		if key != "" {
			param.KeyCode = &zmkKeys[key].Names[0]
		}

		params[i] = param
	}

	return params
}

func canonical(zmkKeys map[string]zmk.KeyCode, mods []string) []string {

	canonical := make([]string, len(mods))
	for i, code := range mods {
		canonical[i] = zmkKeys[code].Names[0]
	}
	return canonical
}

func parseKeys(input *string) (string, []string) {
	keys := []string{}

	if input == nil {
		return "", keys
	}

	current := []rune{}
	for _, char := range *input {

		if char == '(' {
			// modifiers are defined as "LS(code)"
			keys = append(keys, string(current)+"(code)")
			current = []rune{}
		} else if char == ')' {
			break
		} else {
			current = append(current, char)
		}
	}

	return string(current), keys
}
