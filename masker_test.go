package matcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMasker(t *testing.T) {
	tokens := Masker("run")
	assert.Equal(t, tokens[0].Value, "run", "they should be equal")
	assert.Equal(t, tokens[0].Required, true, "they should be equal")
	assert.Equal(t, tokens[0].Text, true, "they should be equal")
	assert.Equal(t, tokens[0].Flag, false, "they should be equal")
	assert.Equal(t, len(tokens), 1, "they should be equal")
}

func TestMaskerExtraSpaces(t *testing.T) {
	tokens := Masker("run  away  ")
	assert.Equal(t, tokens[0].Value, "run", "they should be equal")
	assert.Equal(t, tokens[0].Required, true, "they should be equal")
	assert.Equal(t, tokens[0].Text, true, "they should be equal")
	assert.Equal(t, tokens[0].Flag, false, "they should be equal")
	assert.Equal(t, tokens[1].Value, "away", "they should be equal")
	assert.Equal(t, tokens[1].Required, true, "they should be equal")
	assert.Equal(t, tokens[1].Text, true, "they should be equal")
	assert.Equal(t, tokens[1].Flag, false, "they should be equal")
	assert.Equal(t, len(tokens), 2, "they should be equal")
}

func TestMaskerQuotes(t *testing.T) {
	tokens := Masker("run \"very far\" 'away and away'")
	assert.Equal(t, tokens[0].Value, "run", "they should be equal")
	assert.Equal(t, tokens[0].Required, true, "they should be equal")
	assert.Equal(t, tokens[0].Text, true, "they should be equal")
	assert.Equal(t, tokens[0].Flag, false, "they should be equal")
	assert.Equal(t, tokens[1].Value, "very far", "they should be equal")
	assert.Equal(t, tokens[1].Required, true, "they should be equal")
	assert.Equal(t, tokens[1].Text, true, "they should be equal")
	assert.Equal(t, tokens[1].Flag, false, "they should be equal")
	assert.Equal(t, tokens[2].Value, "away and away", "they should be equal")
	assert.Equal(t, tokens[2].Required, true, "they should be equal")
	assert.Equal(t, tokens[2].Text, true, "they should be equal")
	assert.Equal(t, tokens[2].Flag, false, "they should be equal")
	assert.Equal(t, len(tokens), 3, "they should be equal")
}

func TestMaskerRules(t *testing.T) {
	tokens := Masker("run <speed> [distance] away")
	assert.Equal(t, tokens[0].Value, "run", "they should be equal")
	assert.Equal(t, tokens[0].Required, true, "they should be equal")
	assert.Equal(t, tokens[0].Text, true, "they should be equal")
	assert.Equal(t, tokens[0].Flag, false, "they should be equal")
	assert.Equal(t, tokens[1].Value, "speed", "they should be equal")
	assert.Equal(t, tokens[1].Required, true, "they should be equal")
	assert.Equal(t, tokens[1].Text, false, "they should be equal")
	assert.Equal(t, tokens[1].Flag, false, "they should be equal")
	assert.Equal(t, tokens[2].Value, "distance", "they should be equal")
	assert.Equal(t, tokens[2].Required, false, "they should be equal")
	assert.Equal(t, tokens[2].Text, false, "they should be equal")
	assert.Equal(t, tokens[2].Flag, false, "they should be equal")
	assert.Equal(t, tokens[3].Value, "away", "they should be equal")
	assert.Equal(t, tokens[3].Required, true, "they should be equal")
	assert.Equal(t, tokens[3].Text, true, "they should be equal")
	assert.Equal(t, tokens[3].Flag, false, "they should be equal")
	assert.Equal(t, len(tokens), 4, "they should be equal")
}

func TestMaskerFlags(t *testing.T) {
	tokens := Masker("run --distance=far --speed=\"super very fast\" --skip ")
	assert.Equal(t, tokens[0].Value, "run", "they should be equal")
	assert.Equal(t, tokens[0].Required, true, "they should be equal")
	assert.Equal(t, tokens[0].Text, true, "they should be equal")
	assert.Equal(t, tokens[0].Flag, false, "they should be equal")
	assert.Equal(t, tokens[1].Value, "distance=far", "they should be equal")
	assert.Equal(t, tokens[1].Required, false, "they should be equal")
	assert.Equal(t, tokens[1].Text, false, "they should be equal")
	assert.Equal(t, tokens[1].Flag, true, "they should be equal")
	assert.Equal(t, tokens[2].Value, "speed=super very fast", "they should be equal")
	assert.Equal(t, tokens[2].Required, false, "they should be equal")
	assert.Equal(t, tokens[2].Text, false, "they should be equal")
	assert.Equal(t, tokens[2].Flag, true, "they should be equal")
	assert.Equal(t, tokens[3].Value, "skip", "they should be equal")
	assert.Equal(t, tokens[3].Required, false, "they should be equal")
	assert.Equal(t, tokens[3].Text, false, "they should be equal")
	assert.Equal(t, tokens[3].Flag, true, "they should be equal")
	assert.Equal(t, len(tokens), 4, "they should be equal")
}
