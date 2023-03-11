package keyboard

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
	"strings"
)

func Import(originalPath string) (*Keymap, error) {

	content, err := ioutil.ReadFile(originalPath)
	if err != nil {
		return nil, err
	}

	type dto struct {
		LayerNames []string   `json:"layer_names"`
		Layers     [][]string `json:"layers"`
	}

	original := &dto{}
	if err := json.Unmarshal(content, original); err != nil {
		return nil, err
	}

	keymap := &Keymap{
		Layers: make([]Layer, len(original.LayerNames)),
	}

	for i := 0; i < len(original.LayerNames); i++ {
		layer := Layer{
			Name:     original.LayerNames[i],
			Bindings: make([]Keybind, len(original.Layers[i])),
		}

		for k, code := range original.Layers[i] {
			layer.Bindings[k] = parseKeycode(code)
		}

		keymap.Layers[i] = layer
	}

	return keymap, nil

}

var keyTypeRx = regexp.MustCompile(`^(&.+?)\b`)
var paramsRx = regexp.MustCompile(`\((.+)\)`)

func parseKeycode(keycode string) Keybind {
	keyType := keyTypeRx.FindStringSubmatch(keycode)

	codes := []string{}

	keybind := strings.TrimSpace(keyTypeRx.ReplaceAllString(keycode, ""))

	value := paramsRx.ReplaceAllString(keybind, "")
	if value != "" {
		codes = append(codes, strings.Split(value, " ")...)
	}

	rem := paramsRx.FindStringSubmatch(keybind)

	for len(rem) > 0 {
		value = paramsRx.ReplaceAllString(rem[1], "")
		if value != "" {
			codes = append(codes, value)
		}
		rem = paramsRx.FindStringSubmatch(rem[1])
	}

	return Keybind{
		Type:  strings.TrimLeft(keyType[0], "&"),
		Codes: codes,
	}
}
