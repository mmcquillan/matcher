package matcher

import (
	"strings"
)

// Mask struct
type Mask struct {
	Name      string
	Type      string
	Valid     string
	Required  bool
	Flag      bool
	Remainder bool
}

// Masker function
func Masker(input string) (tokens []Mask) {
	ts := Tokenize(input)
	tokens = make([]Mask, len(ts))
	for i, tn := range ts {
		l := len(tn)
		if strings.HasPrefix(tn, "<--") && strings.HasSuffix(tn, ">") {
			n, t, v := typer(tn[3 : l-1])
			tokens[i] = Mask{
				Name:      n,
				Required:  true,
				Flag:      true,
				Remainder: false,
				Type:      t,
				Valid:     v,
			}
		} else if strings.HasPrefix(tn, "[--") && strings.HasSuffix(tn, "]") {
			n, t, v := typer(tn[3 : l-1])
			tokens[i] = Mask{
				Name:      n,
				Required:  false,
				Flag:      true,
				Remainder: false,
				Type:      t,
				Valid:     v,
			}
		} else if strings.HasPrefix(tn, "[") && strings.HasSuffix(tn, "...]") {
			n, t, v := typer(tn[1 : l-4])
			tokens[i] = Mask{
				Name:      n,
				Required:  false,
				Flag:      false,
				Remainder: true,
				Type:      t,
				Valid:     v,
			}
		} else if strings.HasPrefix(tn, "<") && strings.HasSuffix(tn, "...>") {
			n, t, v := typer(tn[1 : l-4])
			tokens[i] = Mask{
				Name:      n,
				Required:  true,
				Flag:      false,
				Remainder: true,
				Type:      t,
				Valid:     v,
			}
		} else if strings.HasPrefix(tn, "[") && strings.HasSuffix(tn, "]") {
			n, t, v := typer(tn[1 : l-1])
			tokens[i] = Mask{
				Name:      n,
				Required:  false,
				Flag:      false,
				Remainder: false,
				Type:      t,
				Valid:     v,
			}
		} else if strings.HasPrefix(tn, "<") && strings.HasSuffix(tn, ">") {
			n, t, v := typer(tn[1 : l-1])
			tokens[i] = Mask{
				Name:      n,
				Required:  true,
				Flag:      false,
				Remainder: false,
				Type:      t,
				Valid:     v,
			}
		} else {
			tokens[i] = Mask{
				Name:      tn,
				Required:  true,
				Flag:      false,
				Remainder: false,
				Type:      "text",
				Valid:     tn,
			}
		}
	}
	return tokens
}

// typer - name, type, value
func typer(mask string) (n string, t string, v string) {
	n = mask
	t = "string"
	v = "*"
	if strings.Contains(mask, "(") && strings.Contains(mask, ")") {
		n = mask[:strings.Index(mask, "(")]
		pt := mask[strings.Index(mask, "(")+1 : strings.Index(mask, ")")]
		if strings.Contains(pt, ":") {
			t = strings.Split(pt, ":")[0]
			v = strings.Split(pt, ":")[1]
		} else {
			t = pt
			v = "*"
		}
	}
	return n, t, v
}
