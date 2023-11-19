# Todo

- Doc whole project

# Tokenizer

This is my final project for my automata theory class.

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

1. Compile program

~~~bash
make
~~~

1. Execute binary (they are saved in the *bin* folder)

### Without compiling

1. Run the program from the main directory

~~~bash
go run cmd/tokenizer/main.go
~~~

---

jan :3
