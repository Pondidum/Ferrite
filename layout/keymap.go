package layout

import (
	"ferrite/bindings"
	"ferrite/zmk"
	"strconv"
	"strings"
)

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

func (p *Parameter) String() string {
	if p.Number != nil {
		return strconv.Itoa(int(*p.Number))
	}

	return strings.Join(append(p.Modifiers, *p.KeyCode), "")
}

type Layer struct {
	Name     string    `json:"name"`
	Bindings []Binding `json:"bindings"`
}

type ZmkMapper struct {
}

func (m *ZmkMapper) KeymapFromZmk(f *zmk.File) (Keymap, []bindings.BindSet) {

	km := Keymap{
		Configs: make([]Configuration, len(f.Configs)),
		Combos:  make([]Combo, len(f.Device.Combos.Combos)),
		Layers:  make([]Layer, len(f.Device.Keymap.Layers)),
	}

	bindingSets := []bindings.BindSet{bindings.DefaultBindings()}

	for _, include := range f.Includes {
		if strings.HasPrefix(include.External, "keys_en_gb") {
			bindingSets = append(bindingSets, bindings.GbBindings())
		}
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
			Bindings:     m.convertBindings(c.Bindings),
		}
	}

	for i, layer := range f.Device.Keymap.Layers {
		km.Layers[i] = Layer{
			Name:     layer.Name,
			Bindings: m.convertBindings(layer.Bindings),
		}
	}

	return km, bindingSets
}

func (m *ZmkMapper) convertBindings(behaviours []*zmk.Behavior) []Binding {
	bindings := make([]Binding, len(behaviours))

	for i, b := range behaviours {
		bindings[i] = Binding{
			Action: b.Action,
			Params: m.convertParams(b.Params),
		}
	}

	return bindings
}

func (m *ZmkMapper) convertParams(bindings []*zmk.Binding) []Parameter {
	params := make([]Parameter, len(bindings))

	for i, b := range bindings {

		key, modifiers := m.parseKeys(b.KeyCode)
		param := Parameter{
			Number:    b.Number,
			Modifiers: zmk.Canonicalise(modifiers),
		}

		if key != "" {
			canonicalKey := zmk.Canonical(key)
			param.KeyCode = &canonicalKey
		}

		params[i] = param
	}

	return params
}

func (m *ZmkMapper) parseKeys(input *string) (string, []string) {
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
