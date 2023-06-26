package zmk

import (
	"bytes"
	"ferrite/keyboard"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

func Write(w io.Writer, keyboard *keyboard.Keyboard, f *File) {

	writeIncludes(w, f.Includes)

	// defines?
	writeConfigs(w, f.Configs)

	writeDevice(w, keyboard, f.Device)
}

func writeIncludes(w io.Writer, includes []*Include) {

	for _, inc := range includes {
		io.WriteString(w, fmt.Sprintf("#include <%s>\n", inc.Value))
	}
	io.WriteString(w, "\n")

}

func writeConfigs(w io.Writer, configs []*Config) {

	for _, conf := range configs {
		io.WriteString(w, fmt.Sprintf("&%s {\n", conf.Behavior))

		for _, value := range conf.Values {

			if value.Value.Number != nil {
				io.WriteString(w, fmt.Sprintf("\t%s = <%v>;\n", *value.Key, *value.Value.Number))
			}

			if value.Value.String != nil {
				io.WriteString(w, fmt.Sprintf("\t%s = \"%s\";\n", *value.Key, *value.Value.String))
			}

		}
		io.WriteString(w, "};\n")
	}
	io.WriteString(w, "\n")
}

func writeDevice(w io.Writer, keyboard *keyboard.Keyboard, device *Device) {

	io.WriteString(w, "/ {\n\n")

	writeCombos(w, device.Combos)

	io.WriteString(w, "\n")

	writeKeymap(w, keyboard, device.Keymap)

	io.WriteString(w, "};\n")
	io.WriteString(w, "\n")

}

func writeCombos(w io.Writer, combos *Combos) {
	if combos == nil {
		return
	}

	io.WriteString(w, "\tcombos {\n")
	io.WriteString(w, fmt.Sprintf("\t\tcompatible = \"%s\";\n", combos.Compatible))

	for _, combo := range combos.Combos {
		io.WriteString(w, fmt.Sprintf("\t\t%s {\n", combo.Name))

		io.WriteString(w, fmt.Sprintf("\t\t\ttimeout-ms = <%d>;\n", combo.Timeout))
		io.WriteString(w, fmt.Sprintf("\t\t\tkey-positions = <%v>;\n", renderIntList(combo.KeyPositions)))
		io.WriteString(w, fmt.Sprintf("\t\t\tlayers = <%v>;\n", renderIntList(combo.Layers)))
		io.WriteString(w, fmt.Sprintf("\t\t\tbindings = <%v>;\n", renderBehaviors(combo.Bindings)))

		io.WriteString(w, "\t\t};\n")
	}

	io.WriteString(w, "\t};\n")
}

func renderIntList(l []int32) string {

	vals := make([]string, len(l))
	for i, val := range l {
		vals[i] = strconv.FormatInt(int64(val), 10)
	}
	return strings.Join(vals, " ")
}

func writeKeymap(w io.Writer, keyboard *keyboard.Keyboard, keymap *Keymap) {
	io.WriteString(w, "\tkeymap {\n")
	io.WriteString(w, fmt.Sprintf("\t\tcompatible = \"%s\";\n", keymap.Compatible))
	io.WriteString(w, "\n")

	for _, layer := range keymap.Layers {
		io.WriteString(w, fmt.Sprintf("\t\t%s {\n", layer.Name))
		io.WriteString(w, "\t\t\tbindings = <\n")

		renderBindings(w, keyboard, layer.Bindings)

		io.WriteString(w, "\t\t\t>;\n")
		io.WriteString(w, "\t\t};\n")
		io.WriteString(w, "\n")
	}

	io.WriteString(w, "\t};\n")
}

func renderBindingList(list []*Binding) string {
	b := &bytes.Buffer{}

	for i, l := range list {
		io.WriteString(b, renderListItem(l))

		if i < len(list)-1 {
			io.WriteString(b, " ")
		}
	}

	return b.String()
}

func renderListItem(l *Binding) string {
	if l.Number != nil {
		return strconv.FormatInt(int64(*l.Number), 10)
	}

	if l.KeyCode != nil {
		return *l.KeyCode
	}

	return ""

}

func renderBehaviors(behaviors []*Behavior) string {
	b := &bytes.Buffer{}

	for i, l := range behaviors {
		io.WriteString(b, renderBehavior(l))

		if i < len(behaviors)-1 {
			io.WriteString(b, " ")
		}
	}

	return b.String()
}

func renderBehavior(behavior *Behavior) string {

	return fmt.Sprintf("&%s %s", behavior.Action, renderBindingList(behavior.Params))
}

func renderBindings(w io.Writer, kb *keyboard.Keyboard, bindings []*Behavior) {

	keys := make([]string, len(bindings))
	runesPerColumn := 0

	for i, b := range bindings {
		rendered := renderBehavior(b)

		keys[i] = rendered
		if chars := len(rendered); chars > runesPerColumn {
			runesPerColumn = chars
		}
	}

	rowCount := 0
	colCount := 0

	for _, key := range kb.Layout {
		if key.Col > colCount {
			colCount = key.Col
		}
		if key.Row > rowCount {
			rowCount = key.Row
		}
	}

	rowCount += 1
	colCount += 1

	// separator
	sep := " "
	runesPerColumn += len(sep)

	// row + column matrix
	rows := make([][]rune, rowCount)
	for r := range rows {
		rows[r] = repeat(runesPerColumn*colCount, ' ')
	}

	for i, key := range keys {
		r := kb.Layout[i].Row
		c := kb.Layout[i].Col

		for j, char := range key {

			rows[r][(c*runesPerColumn)+j] = char
		}
	}

	for _, row := range rows {
		io.WriteString(w, strings.TrimRightFunc(string(row), unicode.IsSpace))
		io.WriteString(w, "\n")
	}
}

func repeat(length int, r rune) []rune {
	row := make([]rune, length)
	for i := range row {
		row[i] = r
	}

	return row
}
