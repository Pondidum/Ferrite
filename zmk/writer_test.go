package zmk

import (
	"bytes"
	"ferrite/keyboard"
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

	keymap {
		compatible = "zmk,keymap";

		default_layer {
			bindings = <
			>;
		}

		layer_NUM {
			bindings = <
			>;
		}

		layer_SYM {
			bindings = <
			>;
		}

		layer_NAV {
			bindings = <
			>;
		}

		layer_WM {
			bindings = <
			>;
		}

		layer_SYS {
			bindings = <
			>;
		}

		layer_FUN {
			bindings = <
			>;
		}

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

func TestWriteBindings(t *testing.T) {
	f, err := os.Open("cradio.keymap")
	assert.NoError(t, err)
	defer f.Close()

	conf, err := Parse(f)
	assert.NoError(t, err)

	layer := conf.Device.Keymap.Layers[0]
	bindings := layer.Bindings

	kb, err := keyboard.ReadKeyboardInfo("../config/keyboard.json")
	assert.NoError(t, err)
	b := &bytes.Buffer{}
	renderBindings(b, kb, bindings)

	var expected = strings.TrimLeftFunc(`
&kp Q       &kp W       &kp E       &kp R       &kp T       &kp Y       &kp U       &kp I       &kp O       &kp P
&kp A       &mt LGUI S  &mt LALT D  &kp F       &kp G       &kp H       &kp J       &mt LALT K  &mt LGUI L  &kp LSHIFT
&lt 6 Z     &kp X       &kp C       &kp V       &kp B       &kp N       &kp M       &kp COMMA   &kp DOT     &kp LCTRL
&lt 3 TAB   &lt 2 SPACE &kp RET     &mo 1
`, unicode.IsSpace)

	assert.Equal(t, expected, b.String())

}
