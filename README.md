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
match, command, values := matcher.Matcher("run <speed> [distance]", "run fast far")
if match {
	fmt.Printf("Command:%v\n", command)
	for k, v := range values {
		fmt.Printf("%v:%v\n", k, v)
	}
	fmt.Print("\n")
}

- - -
Command:run
speed:fast
distance:far
```


## Masker

```
tokens := matcher.Masker("run <speed> [distance] [--jump]")
for _, t := range tokens {
	fmt.Printf("Name:%v ", t.Name)
	fmt.Printf("Type:%v ", t.Type)
	fmt.Printf("Valid:%v ", t.Valid)
	fmt.Printf("Required:%v ", t.Required)
	fmt.Printf("Flag:%v ", t.Flag)
	fmt.Printf("Remainder:%v\n", t.Remainder)
}

- - -
Name:run Type:text Valid:run Required:true Flag:false Remainder:false
Name:speed Type:string Valid:* Required:true Flag:false Remainder:false
Name:distance Type:string Valid:* Required:false Flag:false Remainder:false
Name:jump Type:string Valid:* Required:false Flag:true Remainder:false
```


## Parser

```
args, flags := matcher.Parser("run far away --jump=high")
for i, arg := range args {
	fmt.Printf("Arg %v: %v\n", i, arg)
}
for _, flag := range flags {
	fmt.Printf("%v: %v\n", flag.Name, flag.Value)
}

- - -
Arg 0: run
Arg 1: far
Arg 2: away
jump: high
```


## Tokenize

```
tokens := matcher.Tokenize(" run far away ")
for _, t := range tokens {
	fmt.Println(t)
}

- - -
run
far
away
```

