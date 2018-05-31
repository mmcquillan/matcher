package matcher

import (
	"strings"
)

// Flag struct
type Flag struct {
	Name  string
	Value string
}

// Parser function
func Parser(input string) (args []string, flags []Flag) {
	ts := Tokenize(input)
	args = make([]string, 0)
	flags = make([]Flag, 0)
	for _, t := range ts {
		if strings.HasPrefix(t, "--") {
			if strings.Contains(t, "=") {
				flags = append(flags, Flag{
					Name:  strings.Split(t, "=")[0][2:],
					Value: strings.Split(t, "=")[1],
				})
			} else {
				flags = append(flags, Flag{
					Name:  t[2:],
					Value: "true",
				})
			}
		} else {
			args = append(args, t)
		}
	}
	return args, flags
}
