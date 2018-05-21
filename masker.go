package matcher

import (
	"strings"
)

// Mask struct
type Mask struct {
	Value     string
	Required  bool
	Text      bool
	Flag      bool
	Remainder bool
}

// Masker function
func Masker(input string) (tokens []Mask) {
	ts := Tokenize(input)
	tokens = make([]Mask, len(ts))
	for i, t := range ts {
		l := len(t)
		if strings.HasPrefix(t, "[") && strings.HasSuffix(t, "...]") {
			tokens[i] = Mask{
				Value:     t[1 : l-4],
				Required:  false,
				Text:      false,
				Flag:      false,
				Remainder: true,
			}
		} else if strings.HasPrefix(t, "<") && strings.HasSuffix(t, "...>") {
			tokens[i] = Mask{
				Value:     t[1 : l-4],
				Required:  true,
				Text:      false,
				Flag:      false,
				Remainder: true,
			}
		} else if strings.HasPrefix(t, "[") && strings.HasSuffix(t, "]") {
			tokens[i] = Mask{
				Value:     t[1 : l-1],
				Required:  false,
				Text:      false,
				Flag:      false,
				Remainder: false,
			}
		} else if strings.HasPrefix(t, "<") && strings.HasSuffix(t, ">") {
			tokens[i] = Mask{
				Value:     t[1 : l-1],
				Required:  true,
				Text:      false,
				Flag:      false,
				Remainder: false,
			}
		} else if strings.HasPrefix(t, "--") {
			tokens[i] = Mask{
				Value:     t[2:],
				Required:  false,
				Text:      false,
				Flag:      true,
				Remainder: false,
			}
		} else if strings.HasPrefix(t, "<--") && strings.HasSuffix(t, ">") {
			tokens[i] = Mask{
				Value:     t[3 : l-1],
				Required:  true,
				Text:      false,
				Flag:      true,
				Remainder: false,
			}
		} else if strings.HasPrefix(t, "[--") && strings.HasSuffix(t, "]") {
			tokens[i] = Mask{
				Value:     t[3 : l-1],
				Required:  false,
				Text:      false,
				Flag:      true,
				Remainder: false,
			}
		} else {
			tokens[i] = Mask{
				Value:     t,
				Required:  true,
				Text:      true,
				Flag:      false,
				Remainder: false,
			}
		}
	}
	return tokens
}
