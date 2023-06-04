package zmk

import (
	"fmt"
	"io"
)

func Write(w io.Writer, f *File) {

	for _, inc := range f.Includes {
		io.WriteString(w, fmt.Sprintf("#include <%s>\n", inc.Value))
	}
	io.WriteString(w, "\n")

	// defines?

	for _, conf := range f.Configs {
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
