package bindings

import (
	"bufio"
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"regexp"
	"sync"
)

// source: https://github.com/zmkfirmware/zmk/blob/main/docs/src/data/hid.js
//
//go:embed default_keys.json
var defaultKeysJson []byte
var defaultBinds BindSet
var parseDefaultBinds sync.Once

// source: https://github.com/joelspadin/zmk-locale-generator/
//
//go:embed keys_en_gb_extended.h
var gbKeysHeader []byte
var gbBinds BindSet
var parseGbBinds sync.Once

type BindSet struct {
	Name  string
	Binds map[string]Bind
}

func (bs *BindSet) CanonicalKey(key string) (string, bool) {
	if bind, found := bs.Binds[key]; found {
		return bind.Name, true
	}

	return key, false
}

type Bind struct {
	Name        string
	Aliases     []string
	Description string
	Symbol      string
}

func SetFromName(name string) (BindSet, error) {

	switch name {
	case "default":
		return DefaultBindings(), nil

	case "en_gb":
		return GbBindings(), nil

	default:
		return BindSet{}, fmt.Errorf("invalid bindset name")
	}
}

func DefaultBindings() BindSet {
	parseDefaultBinds.Do(func() {

		source := []struct {
			Names       []string
			Description string
			Context     string
			Symbol      string
		}{}

		json.Unmarshal(defaultKeysJson, &source)

		defaultBinds = BindSet{
			Name:  "default",
			Binds: map[string]Bind{},
		}

		for _, b := range source {

			if b.Context != "Keyboard" {
				continue
			}

			var aliases []string
			if len(b.Names) > 1 {
				aliases = b.Names[1:]
			}

			name := b.Names[0]
			symbol := b.Symbol

			if s, found := customSymbols[name]; found {
				symbol = s
			}

			bind := Bind{
				Name:        name,
				Aliases:     aliases,
				Description: b.Description,
				Symbol:      symbol,
			}

			for _, name := range b.Names {
				defaultBinds.Binds[name] = bind
			}
		}

	})

	return defaultBinds
}

func GbBindings() BindSet {

	parseGbBinds.Do(func() {

		rx := regexp.MustCompile(`#define (.*) \(`)
		gbBinds = BindSet{
			Name:  "en_gb",
			Binds: map[string]Bind{},
		}

		scanner := bufio.NewScanner(bytes.NewReader(gbKeysHeader))
		for scanner.Scan() {

			matches := rx.FindStringSubmatch(scanner.Text())
			if len(matches) == 0 {
				continue
			}

			b := Bind{
				Name: matches[1],
			}

			gbBinds.Binds[b.Name] = b
		}

	})

	return gbBinds
}
