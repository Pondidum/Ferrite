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
