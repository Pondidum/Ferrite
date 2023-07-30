package zmk

import (
	_ "embed"
	"encoding/json"
)

type KeyMap = map[string]KeyCode

// source: https://github.com/zmkfirmware/zmk/blob/main/docs/src/data/hid.js
//
//go:embed keys.json
var keyJson []byte

var keyMap KeyMap

func ReadKeys() (map[string]KeyCode, error) {

	if keyMap == nil {

		var keys []KeyCode
		if err := json.Unmarshal(keyJson, &keys); err != nil {
			return nil, err
		}

		ApplySymbols(keys)

		m := make(map[string]KeyCode, len(keys))

		for _, key := range keys {
			for _, name := range key.Names {
				m[name] = key
			}
		}

		keyMap = m
	}

	return keyMap, nil
}

func Canonical(key string) string {
	return keyMap[key].Names[0]
}

func Canonicalise(keys []string) []string {

	canonical := make([]string, len(keys))
	for i, code := range keys {
		canonical[i] = Canonical(code)
	}
	return canonical
}

type KeyCode struct {
	Names         []string            `json:"names"`
	Symbol        string              `json:"symbol"`
	Description   string              `json:"description"`
	Context       string              `json:"context"`
	Clarify       bool                `json:"clarify"`
	Documentation string              `json:"documentation"`
	OS            Os                  `json:"os"`
	Footnotes     map[string][]string `json:"footnotes"`
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
	Windows bool `json:"windows"`
	Linux   bool `json:"linux"`
	Android bool `json:"android"`
	Macos   bool `json:"macos"`
	Ios     bool `json:"ios"`
}
