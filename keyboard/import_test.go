package keyboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeycodeParsing(t *testing.T) {

	cases := map[string]Keybind{
		"&kp Q":         {Type: "kp", FirstKey: []string{"Q"}},
		"&none":         {Type: "none", FirstKey: []string{}},
		"&kp LS(N1)":    {Type: "kp", FirstKey: []string{"LS", "N1"}},
		"&kp LG(LS(Q))": {Type: "kp", FirstKey: []string{"LG", "LS", "Q"}},
		"&mt LGUI L":    {Type: "mt", FirstKey: []string{"LGUI"}, SecondKey: []string{"L"}},
	}

	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			assert.Equal(t, expected, parseKeySequence(input))
		})
	}
}

func TestImporting(t *testing.T) {

	keymap, err := Import("keymap.json")
	assert.NoError(t, err)

	assert.Len(t, keymap.Layers, 7)

	assert.Equal(t, "default", keymap.Layers[0].Name)
	assert.Len(t, keymap.Layers[0].Bindings, 34)

	assert.Equal(t, "NUM", keymap.Layers[1].Name)
	assert.Len(t, keymap.Layers[1].Bindings, 34)

	assert.Equal(t, "SYM", keymap.Layers[2].Name)
	assert.Len(t, keymap.Layers[2].Bindings, 34)

	assert.Equal(t, "NAV", keymap.Layers[3].Name)
	assert.Len(t, keymap.Layers[3].Bindings, 34)

	assert.Equal(t, "WM", keymap.Layers[4].Name)
	assert.Len(t, keymap.Layers[4].Bindings, 34)

	assert.Equal(t, "SYS", keymap.Layers[5].Name)
	assert.Len(t, keymap.Layers[5].Bindings, 34)

	assert.Equal(t, "FUN", keymap.Layers[6].Name)
	assert.Len(t, keymap.Layers[6].Bindings, 34)

}
