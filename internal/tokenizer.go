package internal

import ()

type Token struct {
	Type    string
	Content string
}

type Tokenizer struct {
	Tokens []Token
}

func (tokenizer *Tokenizer) Mockup() {
	tokenizer.Tokens = []Token{
		{
			Type:    "Tipo 1",
			Content: "Contenido 1",
		},
		{
			Type:    "Tipo 2",
			Content: "Contenido 2",
		},
		{
			Type:    "Tipo 3",
			Content: "Contenido 3",
		},
	}
}

func NewTokenizer() Tokenizer {
	t := Tokenizer{}
	return t
}
