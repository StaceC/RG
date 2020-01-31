package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {

	assert := assert.New(t)
	input := "I AM IN TROUBLE"
	expected := "../.-|--/..|-./-|.-.|---|..-|-...|.-..|."
	actual, _ := Encode(input, false)
	assert.Equal(expected, actual)

	expected = "2/1A|B/2|A1/A|1A1|C|2A|A3|1A2|1"
	actual, _ = Encode(input, true)
	assert.Equal(expected, actual)
}

func TestObfuscate(t *testing.T) {
	assert := assert.New(t)

	// Dots
	input := "..."
	expected := "3"
	actual, _ := Obfuscate(input)
	assert.Equal(expected, actual)

	// Dashes
	input = "---"
	expected = "C"
	actual, _ = Obfuscate(input)
	assert.Equal(expected, actual)

	// Mix
	input = ".--.-..."
	expected = "1B1A3"
	actual, _ = Obfuscate(input)
	assert.Equal(expected, actual)

	// Empty
	input = ""
	expected = ""
	actual, _ = Obfuscate(input)
	assert.Equal(expected, actual)

	// Invalid input
	var err error
	input = "blah"
	expected = ""
	actual, err = Obfuscate(input)
	assert.Equal(expected, actual)
	assert.NotNil(err)

	// Invalid Rune
	input = "--.a-"
	expected = ""
	actual, err = Obfuscate(input)
	assert.Equal(expected, actual)
	assert.NotNil(err)
}
