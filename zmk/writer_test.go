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
};

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
	};

	keymap {
		compatible = "zmk,keymap";

		default_layer {
			bindings = <
&kp Q       &kp W       &kp E       &kp R       &kp T                   &kp Y       &kp U       &kp I       &kp O       &kp P
&kp A       &mt LGUI S  &mt LALT D  &kp F       &kp G                   &kp H       &kp J       &mt LALT K  &mt LGUI L  &kp LSHIFT
&lt 6 Z     &kp X       &kp C       &kp V       &kp B                   &kp N       &kp M       &kp COMMA   &kp DOT     &kp LCTRL
                                    &lt 3 TAB   &lt 2 SPACE             &kp RET     &mo 1
			>;
		};

		layer_NUM {
			bindings = <
&kp N1        &kp N2        &kp N3        &kp N4        &kp N5                      &kp N6        &kp N7        &kp N8        &kp N9        &kp N0
&none         &kp LGUI      &kp LALT      &kp LSHIFT    &kp LCTRL                   &kp MINUS     &kp LS(MINUS) &kp EQUAL     &kp LS(EQUAL) &none
&none         &none         &none         &none         &none                       &kp PERIOD    &kp SLASH     &kp LS(N8)    &none         &none
                                          &none         &none                       &none         &none
			>;
		};

		layer_SYM {
			bindings = <
&kp LS(N1)               &kp LS(N2)               &kp LS(N3)               &kp LS(N4)               &kp LS(N5)                                        &kp LS(N6)               &kp LS(N7)               &kp LS(N8)               &kp LS(N9)               &kp LS(N0)
&kp LS(LEFT_BRACKET)     &kp LS(RIGHT_BRACKET)    &kp LEFT_BRACKET         &kp RIGHT_BRACKET        &kp NON_US_HASH                                   &kp LS(APOSTROPHE)       &kp APOSTROPHE           &kp LS(SEMICOLON)        &kp SEMICOLON            &kp LS(SLASH)
&kp LS(NON_US_BACKSLASH) &kp NON_US_BACKSLASH     &none                    &kp LS(NON_US_HASH)      &kp GRAVE                                         &kp MINUS                &kp LS(MINUS)            &kp EQUAL                &kp LS(EQUAL)            &kp SLASH
                                                                           &none                    &none                                             &kp DEL                  &kp BACKSPACE
			>;
		};

		layer_NAV {
			bindings = <
&kp ESCAPE       &none            &kp K_PREV       &kp K_PLAY_PAUSE &kp K_NEXT                        &kp PAGE_UP      &kp HOME         &kp END          &none            &kp PRINTSCREEN
&none            &kp LGUI         &kp LALT         &none            &kp K_VOL_UP                      &kp LEFT_ARROW   &kp DOWN_ARROW   &kp UP_ARROW     &kp RIGHT_ARROW  &kp LSHIFT
&none            &none            &none            &none            &kp K_VOL_DN                      &kp PAGE_DOWN    &none            &none            &none            &kp LCTRL
                                                   &none            &none                             &none            &none
			>;
		};

		layer_WM {
			bindings = <
&kp LG(N1)          &kp LG(N2)          &kp LG(N3)          &kp LG(N4)          &kp LG(N5)                              &kp LG(N6)          &kp LG(N7)          &kp LG(N8)          &kp LG(N9)          &kp LG(N0)
&kp LSHIFT          &kp LG(W)           &kp LG(E)           &kp LG(V)           &kp LG(H)                               &kp LG(LEFT_ARROW)  &kp LG(DOWN_ARROW)  &kp LG(UP_ARROW)    &kp LG(RIGHT_ARROW) &none
&kp LALT            &none               &none               &none               &none                                   &none               &none               &none               &none               &none
                                                            &kp LG(LS(Q))       &none                                   &kp LG(D)           &kp LG(ENTER)
			>;
		};

		layer_SYS {
			bindings = <
&none        &none        &none        &none        &none                     &none        &none        &none        &none        &none
&none        &none        &none        &none        &none                     &none        &none        &none        &none        &none
&none        &none        &none        &none        &bootloader               &bootloader  &none        &none        &none        &none
                                       &none        &none                     &none        &none
			>;
		};

		layer_FUN {
			bindings = <
&none      &none      &none      &none      &none                 &kp F1     &kp F2     &kp F3     &kp F4     &none
&none      &kp LGUI   &kp LALT   &kp LSHIFT &kp LCTRL             &kp F5     &kp F6     &kp F7     &kp F8     &none
&none      &none      &none      &none      &none                 &kp F9     &kp F10    &kp F11    &kp F12    &none
                                 &none      &none                 &none      &none
			>;
		};

	};
};

`, unicode.IsSpace)

func TestWriting(t *testing.T) {
	f, err := os.Open("cradio.keymap")
	assert.NoError(t, err)
	defer f.Close()

	kb, err := keyboard.ReadKeyboardInfo("../config/keyboard.json")
	assert.NoError(t, err)

	k, err := Parse(f)
	assert.NoError(t, err)

	b := &bytes.Buffer{}
	Write(b, kb, k)

	assert.Equal(t, strings.Split(expected, "\n"), strings.Split(b.String(), "\n"))

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

	var expected = strings.Split(strings.TrimLeftFunc(`
&kp Q       &kp W       &kp E       &kp R       &kp T                   &kp Y       &kp U       &kp I       &kp O       &kp P
&kp A       &mt LGUI S  &mt LALT D  &kp F       &kp G                   &kp H       &kp J       &mt LALT K  &mt LGUI L  &kp LSHIFT
&lt 6 Z     &kp X       &kp C       &kp V       &kp B                   &kp N       &kp M       &kp COMMA   &kp DOT     &kp LCTRL
                                    &lt 3 TAB   &lt 2 SPACE             &kp RET     &mo 1
`, unicode.IsSpace), "\n")

	assert.Equal(t, expected, strings.Split(b.String(), "\n"))

}

func TestParseWriteEndToEnd(t *testing.T) {

	kb, err := keyboard.ReadKeyboardInfo("../config/keyboard.json")
	assert.NoError(t, err)

	f, err := os.Open("cradio.keymap")
	assert.NoError(t, err)

	k, err := Parse(f)
	assert.NoError(t, err)
	f.Close()

	out, err := os.Create("cradio.keymap")
	assert.NoError(t, err)
	defer out.Close()

	Write(out, kb, k)
}
