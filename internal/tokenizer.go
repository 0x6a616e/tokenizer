package internal

import (
	"strings"
)

// Custom type, needed for enums
type TokenType int

// Enum definition of the tokens it can recognize
const (
	Undefined TokenType = iota
	Identifier
	Keyword
	Operator
	StringLiteral
	IntegerLiteral
	DecimalLiteral
	Invalid
)

// Defines what a Token is, very simple with only the type and content
type Token struct {
	Type    TokenType
	Content string
}

// Defines a Tokenizer and its array of Tokens
type Tokenizer struct {
	Tokens []Token
}

/*
Splits a string into the first token and the remainder of the string, it does
this by finding the first space, new line or ; and splits it based on that
position, for example:
shiftToken("int a = 10") will return "int", "a = 10"
It can be seen as extracting the first token of a string
*/
func shiftToken(s string) (token, remainder string) {
	s = strings.TrimSpace(s)
	splitIndex := -1

looking:
	for i, ch := range s {
		switch string(ch) {
		case " ", "\n", ";":
			splitIndex = i
			break looking
		}
	}

	if splitIndex < 0 {
		return s, ""
	}

	return s[:splitIndex], s[splitIndex+1:]
}

func (tokenizer *Tokenizer) Tokenize(s string) {
	/* This for will extract the first token of the string each iteration
		   Example: s = "int a = 10"
	       First iteration: t = "int", r = "a = 10"
	       Second iteration: t = "a", r = "= 10"
	       Third iteration: t = "=", r = "10"
	       Fourth iteration: t = "10", r = ""
	*/
	for t, r := shiftToken(s); t != ""; t, r = shiftToken(r) {
		// Each token is appended with an undefined type
		tokenizer.Tokens = append(tokenizer.Tokens, Token{Undefined, t})
	}

	var automata Automata
	// Each token is send to the automata for classification
	for i, t := range tokenizer.Tokens {
		tokenizer.Tokens[i].Type = automata.Analyze(t.Content)
	}
}
