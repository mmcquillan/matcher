# Matcher

A library for parsing and matching based on a mask for use with CLI or Bots.


## Mask Rules

- `xyz` text xyz
- `<xyz>` required var named xyz
- `[xyz]` optional var named xyz
- `[xyz...]` optional var named xyz and captures remaining input
- `[xyz(string:foo,bar)]` optional var named xyz that can only be foo or bar
- `<--xyz>` required flag named xyz
- `[--xyz]` optional flag named xyz
- `[--]` capture any flag


## Examples

| Mask                              | Input                       | Match |
| --------------------------------- | --------------------------- | ----- |
| `run`                             | `run`                       | true  |
| `run`                             | `walk`                      | false |
| `run <speed> [distance]`          | `run fast`                  | true  |
| `run <speed> [distance]`          | `run fast far`              | true  |
| `run <speed> [distance]`          | `run fast far forever`      | false |
| `run <speed> [distance]`          | `run fast "far forever"`    | true  |
| `run <speed> [distance] <--jump>` | `run fast far --jump=high`  | true  |
| `run <speed> [distance] <--jump>` | `run fast far`              | false |
| `run <speed> [distance] [--jump]` | `run fast far --jump=high`  | true  |
| `run <speed> [distance] [--jump]` | `run fast far`              | true  |
| `run <speed> [distance] [--jump]` | `run fast far --skip`       | false |
| `run <speed> [distance] [--]`     | `run fast --jump --skip`    | true  |
| `run [song...]`                   | `run Welcome to the Jungle` | true  |
| `run <speed(string)>`             | `run fast`                  | true  |
| `run <speed(string:fast,slow)>`   | `run fast`                  | true  |
| `run <speed(string:fast,slow)>`   | `run quickly`               | false |
| `run <speed(int)>`                | `run 6`                     | true  |
| `run <speed(int:0,2,4,6)>`        | `run 6`                     | true  |
| `run <speed(int:0,2,4,6)>`        | `run 3`                     | false |
| `run <--jump(bool)>`              | `run --jump`                | true  |
| `run <--jump(bool)>`              | `run --jump=false`          | true  |


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

