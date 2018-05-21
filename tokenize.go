package matcher

import ()

// Tokenize func
func Tokenize(input string) (tokens []string) {
	token := []int32{}
	isQuote := false
	for i, c := range input {
		if c == 34 || c == 39 { // quotes
			isQuote = !isQuote
		}
		if isQuote {
			if c != 34 && c != 39 {
				token = append(token, c)
			}
		} else {
			switch c {
			case 9, 10, 11, 12, 13, 32:
				if len(token) > 0 {
					tokens = append(tokens, string(token))
				}
				token = []int32{}
			default:
				if c != 34 && c != 39 {
					token = append(token, c)
				}
			}
		}
		if i == len(input)-1 && len(token) > 0 {
			tokens = append(tokens, string(token))
		}
	}
	return tokens
}
