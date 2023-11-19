package internal

// Making a set of all the letters
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

// A set of numbers
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

// A set of accepted operators
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

// A set of valid keywords
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

type State int

// Enum for the states of the automata
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

// Struct for the automata, it only has its current state
type Automata struct {
	state State
}

func (automata *Automata) Transform(character string) {
	// Based on the current state it should act differently
	switch automata.state {
	case q0:
		// This switch acts as an if else chain for the received character
		switch {
		// If it's letter or an underscore
		case alphabet[character] || character == "_":
			automata.state = qIdentifier
		// If it's an operator
		case operators[character]:
			automata.state = qOperator
		// If it's a quotation mark
		case character == "\"":
			automata.state = qIncompleteString
		// If it's a number
		case numbers[character]:
			automata.state = qIntegerLiteral
		// Anything else
		default:
			automata.state = qInvalid
		}
	case qIdentifier:
		// This switch acts as an if else chain for the received character
		switch {
		// If it's letter or an underscore
		case alphabet[character] || character == "_":
			automata.state = qIdentifier
		// If it's a number
		case numbers[character]:
			automata.state = qIdentifier
		// Anything else
		default:
			automata.state = qInvalid
		}
	case qOperator:
		// This switch acts as an if else chain for the received character
		switch {
		// Anything else
		default:
			automata.state = qInvalid
		}
	case qIncompleteString:
		// This switch acts as an if else chain for the received character
		switch {
		// If it's a quotation mark
		case character == "\"":
			automata.state = qStringLiteral
		// Anything else
		default:
			automata.state = qIncompleteString
		}
	case qStringLiteral:
		// This switch acts as an if else chain for the received character
		switch {
		// Anything else
		default:
			automata.state = qInvalid
		}
	case qIntegerLiteral:
		// This switch acts as an if else chain for the received character
		switch {
		// If it's a number
		case numbers[character]:
			automata.state = qIntegerLiteral
		// If it's a dot
		case character == ".":
			automata.state = qIntToDec
		// Anything else
		default:
			automata.state = qInvalid
		}
	case qIntToDec:
		// This switch acts as an if else chain for the received character
		switch {
		// If it's a number
		case numbers[character]:
			automata.state = qDecimalLiteral
		// Anything else
		default:
			automata.state = qInvalid
		}
	case qDecimalLiteral:
		// This switch acts as an if else chain for the received character
		switch {
		// If it's a number
		case numbers[character]:
			automata.state = qDecimalLiteral
		// Anything else
		default:
			automata.state = qInvalid
		}
	case qInvalid:
		// This switch acts as an if else chain for the received character
		switch {
		// Anything else
		default:
			automata.state = qInvalid
		}
	}
}

// This classifies a token based on its content
func (automata *Automata) Analyze(s string) TokenType {
	// Initial state
	automata.state = q0
	// For each character the automata checks for transformations
	for _, ch := range s {
		automata.Transform(string(ch))
	}

	// Based on the final state it returns a token type
	switch automata.state {
	case qIdentifier:
		// If is an identifier and is in the list of keywords
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
