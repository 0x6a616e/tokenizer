package internal

import ()

type Token struct {
	Type    string
	Content string
}

type Tokenizer struct {
	Tokens []Token
}

func shiftToken(s string) (token, remainder string) {
	splitIndex := -1
looking:
	for i, ch := range s {
		switch string(ch) {
		case " ", "\n":
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
		tokenizer.Tokens = append(tokenizer.Tokens, Token{"Tipo 1", t})
	}
}

func NewTokenizer() Tokenizer {
	t := Tokenizer{}
	return t
}
