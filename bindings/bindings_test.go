package bindings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingDefaultBindings(t *testing.T) {

	bindings := DefaultBindings()
	assert.NotEmpty(t, bindings)

	assert.Equal(t, "LEFT_SHIFT", bindings["LEFT_SHIFT"].Name)
	assert.Equal(t, "LEFT_SHIFT", bindings["LSHIFT"].Name)
	assert.Equal(t, "LEFT_SHIFT", bindings["LSHFT"].Name)
	assert.Equal(t, "LEFT_SHIFT", bindings["LS(code)"].Name)

	assert.Equal(t, "⇧", bindings["LEFT_SHIFT"].Symbol)
	assert.Equal(t, "CTRL", bindings["LEFT_CONTROL"].Symbol)
	assert.Equal(t, "CTRL", bindings["RIGHT_CONTROL"].Symbol)
	assert.Equal(t, "⌦", bindings["DELETE"].Symbol)

}

func TestParsingGbBindings(t *testing.T) {

	bindings := GbBindings()
	assert.NotEmpty(t, bindings)

	assert.Contains(t, bindings, "GB_EXCLAMATION")
}
