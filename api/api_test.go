package api

import (
	"ferrite/zmk"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {

	assert.Equal(t, []string{"A"}, parseKeys("A"))
	assert.Equal(t, []string{"LS", "A"}, parseKeys("LS(A)"))
	assert.Equal(t, []string{"LS", "LC", "F"}, parseKeys("LS(LC(F))"))
}

func TestMapping(t *testing.T) {

	l, err := decode()
	assert.NoError(t, err)

	assert.Equal(t, []string{"LS", "R"}, l.KeyCode)
}

func decode() (*List, error) {

	code := "LS(R)"
	f := zmk.List{
		KeyCode: &code,
	}

	temp := map[string]any{}
	if err := mapstructure.Decode(f, &temp); err != nil {
		return nil, err
	}

	response := List{}
	d, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: KeyCodeToKeysHookFunc(),
		Result:     &response,
	})
	if err != nil {
		return nil, err
	}

	if err := d.Decode(temp); err != nil {
		return nil, err
	}

	return &response, nil
}
