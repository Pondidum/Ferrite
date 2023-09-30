package devices

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed ferris-sweep.json
var ferrisSweepJson []byte

func AllDevices() []string {
	return []string{
		"ferris-sweep",
	}
}

func ReadDevice(name string) (*Keyboard, error) {
	switch name {
	case "ferris-sweep":
		return parseKeyboard("ferris-sweep", ferrisSweepJson)

	default:
		return nil, fmt.Errorf("unknown device: %s", name)
	}
}

func parseKeyboard(name string, b []byte) (*Keyboard, error) {

	kb := &Keyboard{}
	if err := json.Unmarshal(b, &kb); err != nil {
		return nil, err
	}

	kb.Name = name

	return kb, nil
}
