package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {

	ptr := func(s string) *string { return &s }

	cases := []struct {
		input     string
		key       string
		modifiers []string
	}{
		{input: "A", key: "A", modifiers: []string{}},
		{input: "LS(A)", key: "A", modifiers: []string{"LS(code)"}},
		{input: "LS(LC(F))", key: "F", modifiers: []string{"LS(code)", "LC(code)"}},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			k, m := parseKeys(ptr(tc.input))

			assert.Equal(t, tc.key, k)
			assert.Equal(t, tc.modifiers, m)
		})
	}
}
