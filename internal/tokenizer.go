package internal

import (
	"fmt"
	"strings"
)

type TokenType int

const (
	Undefined TokenType = iota
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
	fmt.Println(Undefined == TokenType(ShowingWelcome))
	for t, r := shiftToken(s); t != ""; t, r = shiftToken(r) {
		tokenizer.Tokens = append(tokenizer.Tokens, Token{Undefined, t})
	}
}

func NewTokenizer() Tokenizer {
	t := Tokenizer{}
	return t
}
