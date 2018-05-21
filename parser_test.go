package matcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	tokens := Parser("run")
	assert.Equal(t, tokens[0].Value, "run", "they should be equal")
	assert.Equal(t, tokens[0].Flag, false, "they should be equal")
	assert.Equal(t, len(tokens), 1, "they should be equal")
}

func TestParserExtraSpaces(t *testing.T) {
	tokens := Parser("run  away  ")
	assert.Equal(t, tokens[0].Value, "run", "they should be equal")
	assert.Equal(t, tokens[0].Flag, false, "they should be equal")
	assert.Equal(t, tokens[1].Value, "away", "they should be equal")
	assert.Equal(t, tokens[1].Flag, false, "they should be equal")
	assert.Equal(t, len(tokens), 2, "they should be equal")
}

func TestParserQuotes(t *testing.T) {
	tokens := Parser("run \"very far\" 'away and away'")
	assert.Equal(t, tokens[0].Value, "run", "they should be equal")
	assert.Equal(t, tokens[0].Flag, false, "they should be equal")
	assert.Equal(t, tokens[1].Value, "very far", "they should be equal")
	assert.Equal(t, tokens[1].Flag, false, "they should be equal")
	assert.Equal(t, tokens[2].Value, "away and away", "they should be equal")
	assert.Equal(t, tokens[2].Flag, false, "they should be equal")
	assert.Equal(t, len(tokens), 3, "they should be equal")
}

func TestParserFlags(t *testing.T) {
	tokens := Parser("run --distance=far --speed=\"super very fast\" --skip ")
	assert.Equal(t, tokens[0].Value, "run", "they should be equal")
	assert.Equal(t, tokens[0].Flag, false, "they should be equal")
	assert.Equal(t, tokens[1].Value, "distance=far", "they should be equal")
	assert.Equal(t, tokens[1].Flag, true, "they should be equal")
	assert.Equal(t, tokens[2].Value, "speed=super very fast", "they should be equal")
	assert.Equal(t, tokens[2].Flag, true, "they should be equal")
	assert.Equal(t, tokens[3].Value, "skip=true", "they should be equal")
	assert.Equal(t, tokens[3].Flag, true, "they should be equal")
	assert.Equal(t, len(tokens), 4, "they should be equal")
}
