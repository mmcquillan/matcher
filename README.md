# Matcher

A library for parsing and matching based on a mask for use with CLI or Bots.


## Rules

- Masks with text matches text exactly
- Mask text becomes command
- All non-text becomes variables
- Variables are named in the mask
- Variables are matched by order
- Variables with <var> will be required
- Variables with [var] will be optional
- Optionals variables should be at the end of the mask
- Variables with <var...> or [var...] will match the remaining input
- Remainder variables must be at the end of the mask
- Flags can be specified as <--flag> or [--flag]
- Flag order does not matter for the Mask
- Flag order does not matter for the Input
- Flags with no value are considered true/false
- A Mask with the flag [--] will allow any flag to match
- Extra or non-matching flags will fail the match


## Examples

Mask: `run`

Match: `run`


Mask: `run <speed> [distance]`

Match: `run fast far`


Mask: `run <speed> [distance] [--jump]`

Match: `run fast far --jump=high`


Mask: `run <speed> [distance...]`

Match: `run fast very very far`


Mask: `run <speed(string)> [distance(int)] [enthusiasm(list:low,high)] [--jump(bool)]`

Match: `run fast 24 low --jump=false`


## Matcher

```
match, command, values := Matcher("run <speed> [distance]", "run fast far")
if match {
  fmt.Printf("Command:%v", command)
  for k, v := range values {
    fmt.Printf("%v:%v", k, v)
  }
  fmt.Print("\n")
}
```


## Masker

```
tokens := Masker("run <speed> [distance] [--jump]")
for _, t := range tokens {
  fmt.Printf("Name:%v", t.Name)
  fmt.Printf("Type:%v", t.Type)
  fmt.Printf("Valid:%v", t.Valid)
  fmt.Printf("Required:%v", t.Required)
  fmt.Printf("Flag:%v", t.Flag)
  fmt.Printf("Remainder:%v", t.Reemainder)
  fmt.Print("\n")
}
```


## Parser

```
args, flags := Parser("run far away --jump=high")
for i, arg := range args {
  fmt.Printf("Arg $v: %v\n", i, arg)
}
for _, flag := range flags {
  fmt.Printf("$v: %v\n", flag.Name, flag.Value)
}
```


## Tokenize

```
tokens := Tokenize(" run far away ")
for _, t := range tokens {
  fmt.Println(t)
}
```

