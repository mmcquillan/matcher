package matcher

import (
	"unicode"
)

// Tokenize func
func Tokenize(input string) (tokens []string) {
	runeInput := []rune(input)

	token := []rune{}
	isQuote := false
	for i, r := range runeInput {
		if r == rune('\'') || r == rune('"') { // quotes
			isQuote = !isQuote
		}
		if isQuote {
			if r != rune('\'') && r != rune('"') {
				token = append(token, r)
			}
		} else {
			if unicode.IsSpace(r) {
				if len(token) > 0 {
					tokens = append(tokens, string(token))
				}
				token = []rune{}
			} else {
				if r != rune('\'') && r != rune('"') {
					token = append(token, r)
				}
			}
		}
		if i == len(runeInput)-1 && len(token) > 0 {
			tokens = append(tokens, string(token))
		}
	}
	return tokens
}
