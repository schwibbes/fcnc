package task

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProp(t *testing.T) {

	line := "  a= b\t"
	res := parseProperty(line)
	assert.Equal(t, "a", res[0])
	assert.Equal(t, "b", res[1])
}

func TestCleanup(t *testing.T) {
	expr := " a b c d "
	res := cleanup(expr)
	assert.Equal(t, "a b c d", res)
}
