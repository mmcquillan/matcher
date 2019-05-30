package matcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	args, shortFlags, longFlags := Parser("run")
	assert.Equal(t, args[0], "run", "they should be equal")
	assert.Equal(t, len(args), 1, "they should be equal")
	assert.Equal(t, len(shortFlags), 0, "they should be equal")
	assert.Equal(t, len(longFlags), 0, "they should be equal")
}

func TestParserExtraSpaces(t *testing.T) {
	args, shortFlags, longFlags := Parser("run  away  ")
	assert.Equal(t, args[0], "run", "they should be equal")
	assert.Equal(t, args[1], "away", "they should be equal")
	assert.Equal(t, len(args), 2, "they should be equal")
	assert.Equal(t, len(shortFlags), 0, "they should be equal")
	assert.Equal(t, len(longFlags), 0, "they should be equal")
}

func TestParserQuotes(t *testing.T) {
	args, shortFlags, longFlags := Parser("run \"very far\" 'away and away'")
	assert.Equal(t, args[0], "run", "they should be equal")
	assert.Equal(t, args[1], "very far", "they should be equal")
	assert.Equal(t, args[2], "away and away", "they should be equal")
	assert.Equal(t, len(args), 3, "they should be equal")
	assert.Equal(t, len(shortFlags), 0, "they should be equal")
	assert.Equal(t, len(longFlags), 0, "they should be equal")
}

func TestParserFlags(t *testing.T) {
	args, shortFlags, longFlags := Parser("run --distance=far --speed=\"super very fast\" --skip -x ")
	assert.Equal(t, args[0], "run", "they should be equal")
	assert.Equal(t, len(args), 1, "they should be equal")
	assert.Equal(t, longFlags[0].Name, "distance", "they should be equal")
	assert.Equal(t, longFlags[0].Value, "far", "they should be equal")
	assert.Equal(t, longFlags[1].Name, "speed", "they should be equal")
	assert.Equal(t, longFlags[1].Value, "super very fast", "they should be equal")
	assert.Equal(t, longFlags[2].Name, "skip", "they should be equal")
	assert.Equal(t, longFlags[2].Value, "true", "they should be equal")
	assert.Equal(t, shortFlags[0].Name, "x", "they should be equal")
	assert.Equal(t, shortFlags[0].Value, "true", "they should be equal")
	assert.Equal(t, len(shortFlags), 1, "they should be equal")
	assert.Equal(t, len(longFlags), 3, "they should be equal")
}
