package matcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMasker(t *testing.T) {
	tokens := Masker("run")
	assert.Equal(t, tokens[0].Name, "run", "they should be equal")
	assert.Equal(t, tokens[0].Required, true, "they should be equal")
	assert.Equal(t, tokens[0].Type, "text", "they should be equal")
	assert.Equal(t, tokens[0].ShortFlag, false, "they should be equal")
	assert.Equal(t, tokens[0].LongFlag, false, "they should be equal")
	assert.Equal(t, len(tokens), 1, "they should be equal")
}

func TestMaskerExtraSpaces(t *testing.T) {
	tokens := Masker("run  away  ")
	assert.Equal(t, tokens[0].Name, "run", "they should be equal")
	assert.Equal(t, tokens[0].Required, true, "they should be equal")
	assert.Equal(t, tokens[0].Type, "text", "they should be equal")
	assert.Equal(t, tokens[0].ShortFlag, false, "they should be equal")
	assert.Equal(t, tokens[0].LongFlag, false, "they should be equal")
	assert.Equal(t, tokens[1].Name, "away", "they should be equal")
	assert.Equal(t, tokens[1].Required, true, "they should be equal")
	assert.Equal(t, tokens[1].Type, "text", "they should be equal")
	assert.Equal(t, tokens[1].ShortFlag, false, "they should be equal")
	assert.Equal(t, tokens[1].LongFlag, false, "they should be equal")
	assert.Equal(t, len(tokens), 2, "they should be equal")
}

func TestMaskerFlags(t *testing.T) {
	tokens := Masker("run <--fast> [-x]")
	assert.Equal(t, tokens[0].Name, "run", "they should be equal")
	assert.Equal(t, tokens[0].Required, true, "they should be equal")
	assert.Equal(t, tokens[0].Type, "text", "they should be equal")
	assert.Equal(t, tokens[0].ShortFlag, false, "they should be equal")
	assert.Equal(t, tokens[0].LongFlag, false, "they should be equal")
	assert.Equal(t, tokens[1].Name, "fast", "they should be equal")
	assert.Equal(t, tokens[1].Required, true, "they should be equal")
	assert.Equal(t, tokens[1].Type, "string", "they should be equal")
	assert.Equal(t, tokens[1].ShortFlag, false, "they should be equal")
	assert.Equal(t, tokens[1].LongFlag, true, "they should be equal")
	assert.Equal(t, tokens[2].Name, "x", "they should be equal")
	assert.Equal(t, tokens[2].Required, false, "they should be equal")
	assert.Equal(t, tokens[2].Type, "string", "they should be equal")
	assert.Equal(t, tokens[2].ShortFlag, true, "they should be equal")
	assert.Equal(t, tokens[2].LongFlag, false, "they should be equal")
	assert.Equal(t, len(tokens), 3, "they should be equal")
}

func TestMaskerQuotes(t *testing.T) {
	tokens := Masker("run \"very far\" 'away and away'")
	assert.Equal(t, tokens[0].Name, "run", "they should be equal")
	assert.Equal(t, tokens[0].Required, true, "they should be equal")
	assert.Equal(t, tokens[0].Type, "text", "they should be equal")
	assert.Equal(t, tokens[0].ShortFlag, false, "they should be equal")
	assert.Equal(t, tokens[0].LongFlag, false, "they should be equal")
	assert.Equal(t, tokens[1].Name, "very far", "they should be equal")
	assert.Equal(t, tokens[1].Required, true, "they should be equal")
	assert.Equal(t, tokens[1].Type, "text", "they should be equal")
	assert.Equal(t, tokens[1].ShortFlag, false, "they should be equal")
	assert.Equal(t, tokens[1].LongFlag, false, "they should be equal")
	assert.Equal(t, tokens[2].Name, "away and away", "they should be equal")
	assert.Equal(t, tokens[2].Required, true, "they should be equal")
	assert.Equal(t, tokens[2].Type, "text", "they should be equal")
	assert.Equal(t, tokens[2].ShortFlag, false, "they should be equal")
	assert.Equal(t, tokens[2].LongFlag, false, "they should be equal")
	assert.Equal(t, len(tokens), 3, "they should be equal")
}

func TestMaskerRules(t *testing.T) {
	tokens := Masker("run <speed> [distance] away")
	assert.Equal(t, tokens[0].Name, "run", "they should be equal")
	assert.Equal(t, tokens[0].Required, true, "they should be equal")
	assert.Equal(t, tokens[0].Type, "text", "they should be equal")
	assert.Equal(t, tokens[0].ShortFlag, false, "they should be equal")
	assert.Equal(t, tokens[0].LongFlag, false, "they should be equal")
	assert.Equal(t, tokens[1].Name, "speed", "they should be equal")
	assert.Equal(t, tokens[1].Required, true, "they should be equal")
	assert.Equal(t, tokens[1].Type, "string", "they should be equal")
	assert.Equal(t, tokens[1].ShortFlag, false, "they should be equal")
	assert.Equal(t, tokens[1].LongFlag, false, "they should be equal")
	assert.Equal(t, tokens[2].Name, "distance", "they should be equal")
	assert.Equal(t, tokens[2].Required, false, "they should be equal")
	assert.Equal(t, tokens[2].Type, "string", "they should be equal")
	assert.Equal(t, tokens[2].ShortFlag, false, "they should be equal")
	assert.Equal(t, tokens[2].LongFlag, false, "they should be equal")
	assert.Equal(t, tokens[3].Name, "away", "they should be equal")
	assert.Equal(t, tokens[3].Required, true, "they should be equal")
	assert.Equal(t, tokens[3].Type, "text", "they should be equal")
	assert.Equal(t, tokens[3].ShortFlag, false, "they should be equal")
	assert.Equal(t, tokens[3].LongFlag, false, "they should be equal")
	assert.Equal(t, len(tokens), 4, "they should be equal")
}

func TestMaskerTypes(t *testing.T) {
	tokens := Masker("run <speed(string)> [distance(int)] [enthusiasm(list:low,high)] [--jump(bool)]")
	assert.Equal(t, tokens[0].Name, "run", "they should be equal")
	assert.Equal(t, tokens[0].Type, "text", "they should be equal")
	assert.Equal(t, tokens[1].Name, "speed", "they should be equal")
	assert.Equal(t, tokens[1].Type, "string", "they should be equal")
	assert.Equal(t, tokens[1].Valid, "*", "they should be equal")
	assert.Equal(t, tokens[2].Name, "distance", "they should be equal")
	assert.Equal(t, tokens[2].Type, "int", "they should be equal")
	assert.Equal(t, tokens[3].Name, "enthusiasm", "they should be equal")
	assert.Equal(t, tokens[3].Type, "list", "they should be equal")
	assert.Equal(t, tokens[3].Valid, "low,high", "they should be equal")
	assert.Equal(t, tokens[4].Name, "jump", "they should be equal")
	assert.Equal(t, tokens[4].Type, "bool", "they should be equal")
}
