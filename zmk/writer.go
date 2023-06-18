package zmk

import (
	"bytes"
	"ferrite/keyboard"
	"fmt"
	"io"
	"strconv"
)

func Write(w io.Writer, f *File) {

	writeIncludes(w, f.Includes)

	// defines?
	writeConfigs(w, f.Configs)

	writeDevice(w, f.Device)
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
		io.WriteString(w, "}\n")
	}
	io.WriteString(w, "\n")
}

func writeDevice(w io.Writer, device *Device) {

	io.WriteString(w, "/ {\n\n")

	writeCombos(w, device.Combos)

	io.WriteString(w, "\n")

	writeKeymap(w, device.Keymap)

	io.WriteString(w, "}\n")
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
		io.WriteString(w, fmt.Sprintf("\t\t\tkey-positions = <%v>;\n", renderList(combo.KeyPositions)))
		io.WriteString(w, fmt.Sprintf("\t\t\tlayers = <%v>;\n", renderList(combo.Layers)))
		io.WriteString(w, fmt.Sprintf("\t\t\tbindings = <%v>;\n", renderBehaviors(combo.Bindings)))

		io.WriteString(w, "\t\t};\n")
	}

	io.WriteString(w, "\t}\n")
}

func writeKeymap(w io.Writer, keymap *Keymap) {
	io.WriteString(w, "\tkeymap {\n")
	io.WriteString(w, fmt.Sprintf("\t\tcompatible = \"%s\";\n", keymap.Compatible))
	io.WriteString(w, "\n")

	for _, layer := range keymap.Layers {
		io.WriteString(w, fmt.Sprintf("\t\t%s {\n", layer.Name))
		io.WriteString(w, "\t\t\tbindings = <\n")
		io.WriteString(w, "\t\t\t>;\n")
		io.WriteString(w, "\t\t}\n")
		io.WriteString(w, "\n")
	}

	io.WriteString(w, "\t}\n")
}

func renderList(list []*List) string {
	b := &bytes.Buffer{}

	for i, l := range list {
		io.WriteString(b, renderListItem(l))

		if i < len(list)-1 {
			io.WriteString(b, " ")
		}
	}

	return b.String()
}

func renderListItem(l *List) string {
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

	return fmt.Sprintf("&%s %s", behavior.Action, renderList(behavior.Params))
}

func renderBindings(w io.Writer, kb *keyboard.Keyboard, bindings []*Behavior) {

	keys := make([]string, len(bindings))
	maxChars := 0

	for i, b := range bindings {
		rendered := renderBehavior(b)

		keys[i] = rendered
		if chars := len(rendered); chars > maxChars {
			maxChars = chars
		}
	}

	// separator
	sep := " "

	currentRow := 0
	line := bytes.Buffer{}

	for i, key := range keys {
		row := kb.Layout[i].Row

		if row > currentRow {
			w.Write(bytes.TrimSpace(line.Bytes()))
			io.WriteString(w, "\n")

			line.Reset()
			currentRow = row
		}

		line.WriteString(fmt.Sprintf("%-*s%s", maxChars, key, sep))
	}

	w.Write(bytes.TrimSpace(line.Bytes()))
	io.WriteString(w, "\n")

}
