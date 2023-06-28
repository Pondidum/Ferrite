package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {

	ptr := func(s string) *string { return &s }

	assert.Equal(t, []string{"A"}, parseKeys(ptr("A")))
	assert.Equal(t, []string{"LS", "A"}, parseKeys(ptr("LS(A)")))
	assert.Equal(t, []string{"LS", "LC", "F"}, parseKeys(ptr("LS(LC(F))")))
}
