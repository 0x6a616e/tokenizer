package internal

import "strings"

type Token struct {
	Type    string
	Content string
}

type Tokenizer struct {
	Tokens []Token
}

func (tokenizer *Tokenizer) Tokenize(s string) {
	for _, t := range strings.Split(s, " ") {
		tokenizer.Tokens = append(tokenizer.Tokens, Token{"Tipo", t})
	}
}

func NewTokenizer() Tokenizer {
	t := Tokenizer{}
	return t
}
