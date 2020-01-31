package alg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatten(t *testing.T) {

	assert := assert.New(t)

	// testArray = [ 1, [ 2, [ 3 ] ], 4 ] -> [ 1, 2, 3, 4 ]
	a := []interface{}{3}
	b := []interface{}{2, a}
	input := []interface{}{1, b, 4}

	actual := Flatten(input)
	expected := []interface{}{1, 2, 3, 4}
	assert.Equal(expected, actual)

	// Check empty
	expected = []interface{}{}
	actual = Flatten(expected)
	assert.Equal(expected, actual)
}
