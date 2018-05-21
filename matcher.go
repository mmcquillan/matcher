package matcher

import (
	"strings"
)

// Matcher func
func Matcher(mask string, input string) (match bool, command string, values map[string]string) {

	// initialize values
	match = false
	command = ""
	values = make(map[string]string)

	// tokenize input
	inputTokens := Parser(input)
	maskTokens := Masker(mask)

	// process input flags first
	var inputs []string
	for _, token := range inputTokens {
		if token.Flag {
			flag := strings.Split(token.Value, "=")
			values[flag[0]] = flag[1]
		} else {
			inputs = append(inputs, token.Value)
		}
	}

	// process flags not passed in
	for _, token := range maskTokens {
		if token.Flag {
			if _, chk := values[token.Value]; !chk {
				values[token.Value] = "false"
			}
		}
	}

	// compare input to mask
	match = true
	for i, m := range maskTokens {
		if match && i < len(inputs) {
			if m.Text {
				if m.Value == inputs[i] {
					command += inputs[i] + " "
					match = true
				} else {
					match = false
				}
			} else if m.Remainder {
				values[m.Value] = strings.Join(inputs[i:], " ")
				match = true
			} else {
				values[m.Value] = inputs[i]
				match = true
			}
		} else if match {
			if m.Text || m.Required {
				match = false
			} else if !m.Required || m.Remainder {
				match = true
			}
		} else {
			match = false
		}
	}

	// if no remainder and more inputs than masks
	if len(maskTokens) < len(inputs) {
		if !maskTokens[len(maskTokens)-1].Remainder {
			match = false
		}
	}

	return match, strings.TrimSpace(command), values
}
