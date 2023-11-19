# Tokenizer

This is my final project for my automata theory class.

# File distribution

cmd/tokenizer/main.go
: It's the file where the execution begins, it is almost empty as it only starts
the program defined in other files

internal/bubbletea.go
: It includes all the logic for the BubbleTea framework that creates the
CLI

internal/tokenizer.go
: It includes the logic of the tokenizer, its job is to split the tokens and
save them, the classification parts does not go in here

internal/automata.go
: It includes the automata implementation, its job is to receive a token and
classify it using a deterministic finite automaton

# Tokens it recognizes

Identifier
: Any string that starts with a letter or an underscore and contains only
letters, numbers or underscores

Keyword
: An identifier that appears in a defined list of keywords

Operator
: Any of the following operators (+-/*&|^!<>%=)

Literal
: For a string literal it is any string that starts and ends with quotation
marks ("). For number literals it is any amount of numbers (0-9), then
optionally a dot (.) with more numbers

# Usage

## Pregenerated binaries

There are already three pregenerated binaries in the *bin* folder for the three
main operating systems.

## From source

Requirements:
- Go (version >= 1.20)
- Make (just for the compilation case)

### Compiling

1. Clean bin directory

~~~bash
make clean
~~~

2. Compile program

~~~bash
make
~~~

3. Execute binary (they are saved in the *bin* folder)

### Without compiling

1. Run the program from the main directory

~~~bash
go run cmd/tokenizer/main.go
~~~

---

jan :3
