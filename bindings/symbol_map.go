package bindings

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

	"LEFT_CONTROL":  "CTRL",
	"RIGHT_CONTROL": "CTRL",

	"MINUS": "-",
	"EQUAL": "=",

	"DELETE":    "⌦",
	"BACKSPACE": "⌫",
	"TAB":       "⇥",

	"NON_US_HASH": "#",
	"GRAVE":       "`",

	"K_PLAY_PAUSE": "⏯",
	"K_PREV":       "⏮",
	"K_NEXT":       "⏭",

	"K_VOL_UP": "🔊",
	"K_VOL_DN": "🔉",
}

// func ApplySymbols(keys []KeyCode) {

// 	for i, key := range keys {
// 		if key.Symbol == "" {

// 			for _, name := range key.Names {
// 				if sym, found := customSymbols[name]; found {
// 					keys[i].Symbol = sym
// 					break
// 				}
// 			}

// 		}
// 	}
// }
