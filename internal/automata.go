package internal

type State int

const (
	q0 State = iota
	qIdentifier
	qKeyword
	qOperator
	qLiteral
	qInvalid
)

var keywords = map[string]bool{
	"var":        true,
	"const":      true,
	"bool":       true,
	"true":       true,
	"false":      true,
	"string":     true,
	"int":        true,
	"int8":       true,
	"int16":      true,
	"int32":      true,
	"int64":      true,
	"uint":       true,
	"uint8":      true,
	"uint16":     true,
	"uint32":     true,
	"uint64":     true,
	"uintptr":    true,
	"byte":       true,
	"rune":       true,
	"float32":    true,
	"float64":    true,
	"complex64":  true,
	"complex128": true,
}

type Automata struct {
	state State
}

func (automata Automata) Analyze(s string) TokenType {
	var tokenType TokenType

	if tokenType == Identifier && keywords[s] {
		tokenType = Keyword
	}

	return tokenType
}
