package zmk

import (
	_ "embed"
	"encoding/json"
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

type KeyCode struct {
	Names         []string
	Description   string
	Context       string
	Clarify       bool
	Documentation string
	OS            Os
	Footnotes     map[string][]string
}

type Os struct {
	Windows bool
	Linux   bool
	Android bool
	Macos   bool
	Ios     bool
}
