# Todo

- Add requirements for running
- Specify better the running process
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

1. Compile

~~~bash
make
~~~

2. Execute binary (according to OS)

~~~bash
./tokenizer-linux
~~~

It is also possible to run without compiling by running

~~~bash
go run .
~~~

---

jan :3
