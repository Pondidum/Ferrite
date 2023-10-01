package bindings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingDefaultBindings(t *testing.T) {

	standard := DefaultBindings()
	assert.NotEmpty(t, standard.Binds)

	assert.Equal(t, "LEFT_SHIFT", standard.Binds["LEFT_SHIFT"].Name)
	assert.Equal(t, "LEFT_SHIFT", standard.Binds["LSHIFT"].Name)
	assert.Equal(t, "LEFT_SHIFT", standard.Binds["LSHFT"].Name)
	assert.Equal(t, "LEFT_SHIFT", standard.Binds["LS(code)"].Name)

	assert.Equal(t, "⇧", standard.Binds["LEFT_SHIFT"].Symbol)
	assert.Equal(t, "CTRL", standard.Binds["LEFT_CONTROL"].Symbol)
	assert.Equal(t, "CTRL", standard.Binds["RIGHT_CONTROL"].Symbol)
	assert.Equal(t, "⌦", standard.Binds["DELETE"].Symbol)

}

func TestParsingGbBindings(t *testing.T) {

	gb := GbBindings()
	assert.NotEmpty(t, gb.Binds)
	assert.Contains(t, gb.Binds, "GB_EXCLAMATION")
}
