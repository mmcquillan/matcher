package matcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	args, flags := Parser("run")
	assert.Equal(t, args[0], "run", "they should be equal")
	assert.Equal(t, len(args), 1, "they should be equal")
	assert.Equal(t, len(flags), 0, "they should be equal")
}

func TestParserExtraSpaces(t *testing.T) {
	args, flags := Parser("run  away  ")
	assert.Equal(t, args[0], "run", "they should be equal")
	assert.Equal(t, args[1], "away", "they should be equal")
	assert.Equal(t, len(args), 2, "they should be equal")
	assert.Equal(t, len(flags), 0, "they should be equal")
}

func TestParserQuotes(t *testing.T) {
	args, flags := Parser("run \"very far\" 'away and away'")
	assert.Equal(t, args[0], "run", "they should be equal")
	assert.Equal(t, args[1], "very far", "they should be equal")
	assert.Equal(t, args[2], "away and away", "they should be equal")
	assert.Equal(t, len(args), 3, "they should be equal")
	assert.Equal(t, len(flags), 0, "they should be equal")
}

func TestParserFlags(t *testing.T) {
	args, flags := Parser("run --distance=far --speed=\"super very fast\" --skip ")
	assert.Equal(t, args[0], "run", "they should be equal")
	assert.Equal(t, len(args), 1, "they should be equal")
	assert.Equal(t, flags[0].Name, "distance", "they should be equal")
	assert.Equal(t, flags[0].Value, "far", "they should be equal")
	assert.Equal(t, flags[1].Name, "speed", "they should be equal")
	assert.Equal(t, flags[1].Value, "super very fast", "they should be equal")
	assert.Equal(t, flags[2].Name, "skip", "they should be equal")
	assert.Equal(t, flags[2].Value, "true", "they should be equal")
	assert.Equal(t, len(flags), 3, "they should be equal")
}
