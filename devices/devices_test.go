package devices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadingDevices(t *testing.T) {

	keyboard, err := ReadDevice("ferris-sweep")
	assert.NoError(t, err)
	assert.Equal(t, "ferris-sweep", keyboard.Name)
}
