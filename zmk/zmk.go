package zmk

import (
	_ "embed"
	"encoding/json"
	"strings"
)

// source: https://github.com/zmkfirmware/zmk/blob/main/docs/src/data/hid.js
//
//go:embed keys.json
var keyJson []byte

func ReadKeys() ([]KeyCode, error) {

	var keys []KeyCode
	if err := json.Unmarshal(keyJson, &keys); err != nil {
		return nil, err
	}

	return keys, nil
}

var customSymbols = map[string]string{
	"NUMBER_1": "1",
	"NUMBER_2": "2",
	"NUMBER_3": "3",
	"NUMBER_4": "4",
	"NUMBER_5": "5",
	"NUMBER_6": "6",
	"NUMBER_7": "7",
	"NUMBER_8": "8",
	"NUMBER_9": "9",
	"NUMBER_0": "0",

	"N1": "1",
	"N2": "2",
	"N3": "3",
	"N4": "4",
	"N5": "5",
	"N6": "6",
	"N7": "7",
	"N8": "8",
	"N9": "9",
	"N0": "0",

	"LEFT_CONTROL": "CTRL",
	"MINUS":        "-",
	"EQUAL":        "=",

	"DEL":       "⌦",
	"DELETE":    "⌦",
	"BACKSPACE": "⌫",

	"NON_US_HASH": "#",
	"GRAVE":       "`",

	"K_PLAY_PAUSE": "⏯",
	"K_PREV":       "⏮",
	"K_NEXT":       "⏭",

	"K_VOL_UP": "🔊",
	"K_VOL_DN": "🔉",
}

func BuildSymbolIndex(keys []KeyCode) map[string]string {

	index := make(map[string]string, len(keys))

	for _, key := range keys {
		for _, name := range key.Names {

			name = strings.TrimSuffix(name, "(code)")

			if sym, found := customSymbols[name]; found {
				index[name] = sym
			} else {
				index[name] = key.String()
			}

		}
	}

	return index
}

type KeyCode struct {
	Names         []string
	Symbol        string
	Description   string
	Context       string
	Clarify       bool
	Documentation string
	OS            Os
	Footnotes     map[string][]string
}

func (kc KeyCode) String() string {
	if kc.Symbol != "" {
		return kc.Symbol
	}

	if len(kc.Names) > 0 {
		return kc.Names[0]
	}

	return ""
}

type Os struct {
	Windows bool
	Linux   bool
	Android bool
	Macos   bool
	Ios     bool
}
