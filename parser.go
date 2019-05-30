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
func Parser(input string) (args []string, shortFlags []Flag, longFlags []Flag) {

	// setup vars and tokenize
	ts := Tokenize(input)
	args = make([]string, 0)
	shortFlags = make([]Flag, 0)
	longFlags = make([]Flag, 0)

	// handle short vars

	// handle long vars
	for _, t := range ts {
		if strings.HasPrefix(t, "--") {
			if strings.Contains(t, "=") {
				longFlags = append(longFlags, Flag{
					Name:  strings.Split(t, "=")[0][2:],
					Value: strings.Split(t, "=")[1],
				})
			} else {
				longFlags = append(longFlags, Flag{
					Name:  t[2:],
					Value: "true",
				})
			}
		} else if strings.HasPrefix(t, "-") {
			if strings.Contains(t, "=") {
				shortFlags = append(shortFlags, Flag{
					Name:  strings.Split(t, "=")[0][1:],
					Value: strings.Split(t, "=")[1],
				})
			} else {
				shortFlags = append(shortFlags, Flag{
					Name:  t[1:],
					Value: "true",
				})
			}
		} else {
			args = append(args, t)
		}
	}
	return args, shortFlags, longFlags
}
