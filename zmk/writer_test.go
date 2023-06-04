package zmk

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

var expected = strings.TrimLeftFunc(`
#include <behaviors.dtsi>
#include <dt-bindings/zmk/keys.h>
#include <dt-bindings/zmk/bt.h>

&mt {
	tapping-term-ms = <200>;
	flavor = "tap-preferred";
}

/ {

	combos {
		compatible = "zmk,combos";
		combo_system {
			timeout-ms = <50>;
			key-positions = <2 3 4>;
			layers = <0 5>;
			bindings = <&tog 5>;
		};
		combo_wm {
			timeout-ms = <50>;
			key-positions = <19 29>;
			layers = <0>;
			bindings = <&mo 4>;
		};
	}
}

`, unicode.IsSpace)

func TestWriting(t *testing.T) {
	f, err := os.Open("cradio.keymap")
	assert.NoError(t, err)
	defer f.Close()

	k, err := Parse(f)
	assert.NoError(t, err)

	b := &bytes.Buffer{}
	Write(b, k)

	assert.Equal(t, expected, b.String())

}
