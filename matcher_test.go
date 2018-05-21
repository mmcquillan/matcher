package matcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatcherCommand(t *testing.T) {
	match, command, values := Matcher("run", "run")
	assert.Equal(t, match, true, "they should be equal")
	assert.Equal(t, command, "run", "they should be equal")
	assert.Equal(t, len(values), 0, "they should be equal")
}

func TestMatcherCommandNot(t *testing.T) {
	match, command, values := Matcher("run", "walk")
	assert.Equal(t, match, false, "they should be equal")
	assert.Equal(t, command, "", "they should be equal")
	assert.Equal(t, len(values), 0, "they should be equal")
}

func TestMatcherCommand2(t *testing.T) {
	match, command, values := Matcher("run away", "run away")
	assert.Equal(t, match, true, "they should be equal")
	assert.Equal(t, command, "run away", "they should be equal")
	assert.Equal(t, len(values), 0, "they should be equal")
}

func TestMatcherRequired(t *testing.T) {
	match, command, values := Matcher("run <speed>", "run fast")
	assert.Equal(t, match, true, "they should be equal")
	assert.Equal(t, command, "run", "they should be equal")
	assert.Equal(t, values["speed"], "fast", "they should be equal")
	assert.Equal(t, len(values), 1, "they should be equal")
}

func TestMatcherRequired2(t *testing.T) {
	match, command, values := Matcher("run <speed> <distance>", "run fast far")
	assert.Equal(t, match, true, "they should be equal")
	assert.Equal(t, command, "run", "they should be equal")
	assert.Equal(t, values["speed"], "fast", "they should be equal")
	assert.Equal(t, values["distance"], "far", "they should be equal")
	assert.Equal(t, len(values), 2, "they should be equal")
}

func TestMatcherRequired2Not(t *testing.T) {
	match, _, _ := Matcher("run <speed> <distance>", "run fast")
	assert.Equal(t, match, false, "they should be equal")
}

func TestMatcherOptional(t *testing.T) {
	match, command, values := Matcher("run [speed]", "run fast")
	assert.Equal(t, match, true, "they should be equal")
	assert.Equal(t, command, "run", "they should be equal")
	assert.Equal(t, values["speed"], "fast", "they should be equal")
}

func TestMatcherOptional2(t *testing.T) {
	match, command, values := Matcher("run [speed] [distance]", "run fast far")
	assert.Equal(t, match, true, "they should be equal")
	assert.Equal(t, command, "run", "they should be equal")
	assert.Equal(t, values["speed"], "fast", "they should be equal")
	assert.Equal(t, values["distance"], "far", "they should be equal")
}

func TestMatcherBoth(t *testing.T) {
	match, command, values := Matcher("run <speed> [distance]", "run fast far")
	assert.Equal(t, match, true, "they should be equal")
	assert.Equal(t, command, "run", "they should be equal")
	assert.Equal(t, values["speed"], "fast", "they should be equal")
	assert.Equal(t, values["distance"], "far", "they should be equal")
}

func TestMatcherAlt1(t *testing.T) {
	match, command, values := Matcher("run <speed> [distance]", "run fast")
	assert.Equal(t, match, true, "they should be equal")
	assert.Equal(t, command, "run", "they should be equal")
	assert.Equal(t, values["speed"], "fast", "they should be equal")
	assert.Equal(t, len(values), 1, "they should be equal")
}

func TestMatcherAlt2(t *testing.T) {
	match, _, _ := Matcher("run <speed> [distance]", "run fast far jump")
	assert.Equal(t, match, false, "they should be equal")
}

func TestMatcherFlag1(t *testing.T) {
	match, command, values := Matcher("run <speed> --skip", "run fast --skip")
	assert.Equal(t, match, true, "they should be equal")
	assert.Equal(t, command, "run", "they should be equal")
	assert.Equal(t, values["speed"], "fast", "they should be equal")
	assert.Equal(t, values["skip"], "true", "they should be equal")
	assert.Equal(t, len(values), 2, "they should be equal")
}

func TestMatcherFlag2(t *testing.T) {
	match, command, values := Matcher("run [speed] --skip", "run --skip")
	assert.Equal(t, match, true, "they should be equal")
	assert.Equal(t, command, "run", "they should be equal")
	assert.Equal(t, values["skip"], "true", "they should be equal")
	assert.Equal(t, len(values), 1, "they should be equal")
}

func TestMatcherNoFlag(t *testing.T) {
	match, command, values := Matcher("run <speed> --skip", "run fast")
	assert.Equal(t, match, true, "they should be equal")
	assert.Equal(t, command, "run", "they should be equal")
	assert.Equal(t, values["speed"], "fast", "they should be equal")
	assert.Equal(t, values["skip"], "false", "they should be equal")
	assert.Equal(t, len(values), 2, "they should be equal")
}

func TestMatcherExtraFlag(t *testing.T) {
	match, command, values := Matcher("run <speed>", "run fast --jump")
	assert.Equal(t, match, true, "they should be equal")
	assert.Equal(t, command, "run", "they should be equal")
	assert.Equal(t, values["speed"], "fast", "they should be equal")
	assert.Equal(t, values["jump"], "true", "they should be equal")
	assert.Equal(t, len(values), 2, "they should be equal")
}

func TestMatcherExtraFlagValue(t *testing.T) {
	match, command, values := Matcher("run <speed>", "run fast --jump=xyz")
	assert.Equal(t, match, true, "they should be equal")
	assert.Equal(t, command, "run", "they should be equal")
	assert.Equal(t, values["speed"], "fast", "they should be equal")
	assert.Equal(t, values["jump"], "xyz", "they should be equal")
	assert.Equal(t, len(values), 2, "they should be equal")
}

func TestMatcherRemaining(t *testing.T) {
	match, command, values := Matcher("run <speed> [song...]", "run fast Welcome to the Jungle")
	assert.Equal(t, match, true, "they should be equal")
	assert.Equal(t, command, "run", "they should be equal")
	assert.Equal(t, values["speed"], "fast", "they should be equal")
	assert.Equal(t, values["song"], "Welcome to the Jungle", "they should be equal")
	assert.Equal(t, len(values), 2, "they should be equal")
}

func TestMatcherNoRemaining(t *testing.T) {
	match, command, values := Matcher("run <speed> [song...]", "run fast")
	assert.Equal(t, match, true, "they should be equal")
	assert.Equal(t, command, "run", "they should be equal")
	assert.Equal(t, values["speed"], "fast", "they should be equal")
	assert.Equal(t, len(values), 1, "they should be equal")
}

func TestMatcherQuotedReq1(t *testing.T) {
	match, command, values := Matcher("run <speed> [distance]", "run \"super fast\"")
	assert.Equal(t, match, true, "they should be equal")
	assert.Equal(t, command, "run", "they should be equal")
	assert.Equal(t, values["speed"], "super fast", "they should be equal")
	assert.Equal(t, len(values), 1, "they should be equal")
}

func TestMatcherQuotedReq2(t *testing.T) {
	match, command, values := Matcher("run <speed> [distance]", "run 'super fast'")
	assert.Equal(t, match, true, "they should be equal")
	assert.Equal(t, command, "run", "they should be equal")
	assert.Equal(t, values["speed"], "super fast", "they should be equal")
	assert.Equal(t, len(values), 1, "they should be equal")
}

func TestMatcherQuotedOpt1(t *testing.T) {
	match, command, values := Matcher("run <speed> [distance]", "run fast \"super far\"")
	assert.Equal(t, match, true, "they should be equal")
	assert.Equal(t, command, "run", "they should be equal")
	assert.Equal(t, values["speed"], "fast", "they should be equal")
	assert.Equal(t, values["distance"], "super far", "they should be equal")
	assert.Equal(t, len(values), 2, "they should be equal")
}

func TestMatcherQuotedOpt2(t *testing.T) {
	match, command, values := Matcher("run <speed> [distance]", "run fast 'super far'")
	assert.Equal(t, match, true, "they should be equal")
	assert.Equal(t, command, "run", "they should be equal")
	assert.Equal(t, values["speed"], "fast", "they should be equal")
	assert.Equal(t, values["distance"], "super far", "they should be equal")
	assert.Equal(t, len(values), 2, "they should be equal")
}

func TestMatcherQuotedFlag1(t *testing.T) {
	match, command, values := Matcher("run --skip", "run --skip=\"just for fun\"")
	assert.Equal(t, match, true, "they should be equal")
	assert.Equal(t, command, "run", "they should be equal")
	assert.Equal(t, values["skip"], "just for fun", "they should be equal")
	assert.Equal(t, len(values), 1, "they should be equal")
}

func TestMatcherQuotedFlag2(t *testing.T) {
	match, command, values := Matcher("run --skip", "run --skip='just for fun'")
	assert.Equal(t, match, true, "they should be equal")
	assert.Equal(t, command, "run", "they should be equal")
	assert.Equal(t, values["skip"], "just for fun", "they should be equal")
	assert.Equal(t, len(values), 1, "they should be equal")
}
