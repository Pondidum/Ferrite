package bindings

import (
	"bufio"
	"bytes"
	_ "embed"
	"encoding/json"
	"regexp"
	"sync"
)

// source: https://github.com/zmkfirmware/zmk/blob/main/docs/src/data/hid.js
//
//go:embed default_keys.json
var defaultKeysJson []byte
var defaultBinds map[string]Bind
var parseDefaultBinds sync.Once

// source: https://github.com/joelspadin/zmk-locale-generator/
//
//go:embed keys_en_gb_extended.h
var gbKeysHeader []byte
var gbBinds map[string]Bind
var parseGbBinds sync.Once

type Bind struct {
	Name        string
	Aliases     []string
	Description string
	Symbol      string
}

func DefaultBindings() map[string]Bind {
	parseDefaultBinds.Do(func() {

		source := []struct {
			Names       []string
			Description string
			Context     string
			Symbol      string
		}{}

		json.Unmarshal(defaultKeysJson, &source)

		defaultBinds = map[string]Bind{}

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
				defaultBinds[name] = bind
			}
		}

	})

	return defaultBinds
}

func GbBindings() map[string]Bind {

	parseGbBinds.Do(func() {

		rx := regexp.MustCompile(`#define (.*) \(`)
		gbBinds = map[string]Bind{}

		scanner := bufio.NewScanner(bytes.NewReader(gbKeysHeader))
		for scanner.Scan() {

			matches := rx.FindStringSubmatch(scanner.Text())
			if len(matches) == 0 {
				continue
			}

			b := Bind{
				Name: matches[1],
			}

			gbBinds[b.Name] = b
		}

	})

	return gbBinds
}
