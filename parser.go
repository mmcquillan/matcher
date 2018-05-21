package matcher

import (
	"strings"
)

// Arg struct
type Arg struct {
	Pos   int
	Value string
	Flag  bool
}

// Parser function
func Parser(input string) (args []Arg) {
	ts := Tokenize(input)
	args = make([]Arg, len(ts))
	for i, t := range ts {
		if strings.HasPrefix(t, "--") {
			if !strings.Contains(t, "=") {
				t += "=true"
			}
			args[i] = Arg{
				Pos:   i,
				Value: t[2:],
				Flag:  true,
			}
		} else {
			args[i] = Arg{
				Pos:   i,
				Value: t,
				Flag:  false,
			}
		}
	}
	return args
}
