package matcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	tokens := Tokenize("run")
	assert.Equal(t, tokens[0], "run", "they should be equal")
	assert.Equal(t, len(tokens), 1, "they should be equal")
}

func TestTokenizeExtraSpaces(t *testing.T) {
	tokens := Tokenize("run  away  ")
	assert.Equal(t, tokens[0], "run", "they should be equal")
	assert.Equal(t, tokens[1], "away", "they should be equal")
	assert.Equal(t, len(tokens), 2, "they should be equal")
}

func TestTokenizeQuotes(t *testing.T) {
	tokens := Tokenize("run \"very far\" 'away and away'")
	assert.Equal(t, tokens[0], "run", "they should be equal")
	assert.Equal(t, tokens[1], "very far", "they should be equal")
	assert.Equal(t, tokens[2], "away and away", "they should be equal")
	assert.Equal(t, len(tokens), 3, "they should be equal")
}

func TestTokenizeRules(t *testing.T) {
	tokens := Tokenize("run <speed> [distance] away")
	assert.Equal(t, tokens[0], "run", "they should be equal")
	assert.Equal(t, tokens[1], "<speed>", "they should be equal")
	assert.Equal(t, tokens[2], "[distance]", "they should be equal")
	assert.Equal(t, tokens[3], "away", "they should be equal")
	assert.Equal(t, len(tokens), 4, "they should be equal")
}

func TestTokenizeFlags(t *testing.T) {
	tokens := Tokenize("run --distance=far --speed=\"super very fast\" --skip ")
	assert.Equal(t, tokens[0], "run", "they should be equal")
	assert.Equal(t, tokens[1], "--distance=far", "they should be equal")
	assert.Equal(t, tokens[2], "--speed=super very fast", "they should be equal")
	assert.Equal(t, tokens[3], "--skip", "they should be equal")
	assert.Equal(t, len(tokens), 4, "they should be equal")
}
