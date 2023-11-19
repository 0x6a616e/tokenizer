package internal

import (
	"strings"
)

type TokenType int

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

type Token struct {
	Type    TokenType
	Content string
}

type Tokenizer struct {
	Tokens []Token
}

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
	for t, r := shiftToken(s); t != ""; t, r = shiftToken(r) {
		tokenizer.Tokens = append(tokenizer.Tokens, Token{Undefined, t})
	}

	var automata Automata
	for i, t := range tokenizer.Tokens {
		tokenizer.Tokens[i].Type = automata.Analyze(t.Content)
	}
}
