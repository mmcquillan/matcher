package matcher

import (
	"strconv"
	"strings"
)

// Matcher func
func Matcher(mask string, input string) (match bool, command string, values map[string]string) {

	// initialize
	match = true
	command = ""
	values = make(map[string]string)
	args, flags := Parser(input)
	masks := Masker(mask)
	pos := 0

	// initial flag pass
	for _, flag := range flags {
		allFlags := false
		matchFlag := false
		for _, mask := range masks {
			if mask.Flag {
				if mask.Name == flag.Name {
					matchFlag = true
				}
				if mask.Name == "" {
					allFlags = true
				}
			}
		}
		if !matchFlag && allFlags {
			values[flag.Name] = flag.Value
		} else if !allFlags && !matchFlag {
			match = false
		}
	}

	// loop over the mask
	for _, mask := range masks {
		if match {
			if mask.Flag {
				pos--
				flagMatch := false
				for _, flag := range flags {
					if mask.Name == flag.Name {

						// check for valid values
						valid := false
						for _, v := range strings.Split(mask.Valid, ",") {
							if v == flag.Value || v == "*" {
								valid = true
							}
						}
						match = valid

						// check for type
						switch mask.Type {
						case "string":
							values[mask.Name] = flag.Value
						case "int":
							values[mask.Name] = flag.Value
							_, err := strconv.ParseInt(flag.Value, 10, 64)
							if err != nil {
								match = false
							}
						case "bool":
							values[mask.Name] = flag.Value
							_, err := strconv.ParseBool(flag.Value)
							if err != nil {
								match = false
							}
						}

						flagMatch = true

					}
				}
				if mask.Required && !flagMatch {
					match = false
				}
			} else {
				if pos < len(args) {

					// eval remainder
					arg := args[pos]
					if mask.Remainder {
						arg = strings.Join(args[pos:], " ")
						pos = len(args)
					}

					// check for valid values
					valid := false
					for _, v := range strings.Split(mask.Valid, ",") {
						if v == arg || v == "*" {
							valid = true
						}
					}
					match = valid

					// check for type
					switch mask.Type {
					case "text":
						if mask.Name != arg {
							match = false
						} else {
							command += arg + " "
						}
					case "string":
						values[mask.Name] = arg
					case "int":
						values[mask.Name] = arg
						_, err := strconv.ParseInt(arg, 10, 64)
						if err != nil {
							match = false
						}
					case "bool":
						values[mask.Name] = arg
						_, err := strconv.ParseBool(arg)
						if err != nil {
							match = false
						}
					}
				} else {

					// check required
					if mask.Required {
						match = false
					}

				}
			}
			pos++
		}
	}

	// check for arg length
	if pos < len(args) {
		match = false
	}

	// reset values on no match
	if !match {
		command = ""
		values = make(map[string]string)
	}

	// return
	return match, strings.TrimSpace(command), values

}
