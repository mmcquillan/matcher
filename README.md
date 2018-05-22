# Matcher

A library for parsing and matching based on a mask for use with CLI or Bots.

Examples:

Mask: `run`
Match: `run`

Mask: `run <speed> [distance]`
Match: `run fast far`

Mask: `run <speed> [distance] [--jump]`
Match: `run fast far --jump=high`


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
  fmt.Printf("Value:%v", t.Value)
  fmt.Printf("Required:%v", t.Required)
  fmt.Printf("Text:%v", t.Text)
  fmt.Printf("Flag:%v", t.Flag)
  fmt.Printf("Remainder:%v", t.Reemainder)
  fmt.Print("\n")
}
```


## Parser

```
tokens := Parser("run far away --jump=high")
for _, t := range tokens {
  fmt.Printf("Value:%v", t.Value)
  fmt.Printf("Flag:%v", t.Flag)
  fmt.Print("\n")
}
```


## Tokenize

```
tokens := Tokenize(" run far away ")
for _, t := range tokens {
  fmt.Println(t)
}
```

