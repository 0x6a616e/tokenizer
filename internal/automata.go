package internal

type State int

var alphabet = map[string]bool{
	"a": true,
	"b": true,
	"c": true,
	"d": true,
	"e": true,
	"f": true,
	"g": true,
	"h": true,
	"i": true,
	"j": true,
	"k": true,
	"l": true,
	"m": true,
	"n": true,
	"o": true,
	"p": true,
	"q": true,
	"r": true,
	"s": true,
	"t": true,
	"u": true,
	"v": true,
	"w": true,
	"x": true,
	"y": true,
	"z": true,
	"A": true,
	"B": true,
	"C": true,
	"D": true,
	"E": true,
	"F": true,
	"G": true,
	"H": true,
	"I": true,
	"J": true,
	"K": true,
	"L": true,
	"M": true,
	"N": true,
	"O": true,
	"P": true,
	"Q": true,
	"R": true,
	"S": true,
	"T": true,
	"U": true,
	"V": true,
	"W": true,
	"X": true,
	"Y": true,
	"Z": true,
}

var numbers = map[string]bool{
	"0": true,
	"1": true,
	"2": true,
	"3": true,
	"4": true,
	"5": true,
	"6": true,
	"7": true,
	"8": true,
	"9": true,
}

var operators = map[string]bool{
	"+": true,
	"-": true,
	"/": true,
	"*": true,
	"&": true,
	"|": true,
	"^": true,
	"!": true,
	"<": true,
	">": true,
	"%": true,
	"=": true,
}

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

const (
	q0 State = iota
	qIdentifier
	qOperator
	qIncompleteString
	qStringLiteral
	qIntegerLiteral
	qIntToDec
	qDecimalLiteral
	qInvalid
)

type Automata struct {
	state State
}

func (automata *Automata) Transform(character string) {
	switch automata.state {
	case q0:
		switch {
		case alphabet[character] || character == "_":
			automata.state = qIdentifier
		case operators[character]:
			automata.state = qOperator
		case character == "\"":
			automata.state = qIncompleteString
		case numbers[character]:
			automata.state = qIntegerLiteral
		default:
			automata.state = qInvalid
		}
	case qIdentifier:
		switch {
		case alphabet[character] || character == "_":
			automata.state = qIdentifier
		case numbers[character]:
			automata.state = qIdentifier
		default:
			automata.state = qInvalid
		}
	case qOperator:
		switch {
		default:
			automata.state = qInvalid
		}
	case qIncompleteString:
		switch {
		case character == "\"":
			automata.state = qStringLiteral
		default:
			automata.state = qIncompleteString
		}
	case qStringLiteral:
		switch {
		default:
			automata.state = qInvalid
		}
	case qIntegerLiteral:
		switch {
		case numbers[character]:
			automata.state = qIntegerLiteral
		case character == ".":
			automata.state = qIntToDec
		default:
			automata.state = qInvalid
		}
	case qIntToDec:
		switch {
		case numbers[character]:
			automata.state = qDecimalLiteral
		default:
			automata.state = qInvalid
		}
	case qDecimalLiteral:
		switch {
		case numbers[character]:
			automata.state = qDecimalLiteral
		default:
			automata.state = qInvalid
		}
	case qInvalid:
		switch {
		default:
			automata.state = qInvalid
		}
	}
}

func (automata *Automata) Analyze(s string) TokenType {
	automata.state = q0
	for _, ch := range s {
		automata.Transform(string(ch))
	}

	switch automata.state {
	case qIdentifier:
		if keywords[s] {
			return Keyword
		}
		return Identifier
	case qOperator:
		return Operator
	case qStringLiteral:
		return StringLiteral
	case qIntegerLiteral:
		return IntegerLiteral
	case qDecimalLiteral:
		return DecimalLiteral
	default:
		return Invalid
	}
}
