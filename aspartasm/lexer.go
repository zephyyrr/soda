package aspartasm

import (
	"bufio"
	"io"
	"strconv"
)

type Token struct {
	Kind  string
	Value string
}

const (
	unknown   = "unknown"
	label     = "label"
	operation = "operation"
	register  = "register"
	number    = "number"
	sträng    = "string"
)

type lexfunc func(chan<- Token, rune, []rune) (lexfunc, []rune)

func lex(in io.Reader) <-chan Token {
	tokens := make(chan Token, 3)
	bufin := bufio.NewReader(in)

	go func() {
		defer close(tokens)
		defer func() {
			if r := recover(); r != nil {
				//Probably EOF at bad time
			}
		}()

		// Set the split function for the scanning operation.
		f := lexfunc(lexStart)
		var part []rune
		t, _, err := bufin.ReadRune()
		for !(f == nil || err != nil) {
			f, part = f(tokens, t, part)
			t, _, err = bufin.ReadRune()
		}
	}()

	return tokens
}

func lexStart(tokens chan<- Token, t rune, part []rune) (lexfunc, []rune) {
	switch t {
	case ' ', '\n', '\t':
		return lexStart, nil
	case '%', '#':
		return lexLineComment, nil
	default:
		return lexOp, []rune{rune(byte(t))}
	}
}

func lexLineComment(tokens chan<- Token, t rune, part []rune) (lexfunc, []rune) {
	switch t {
	case '\r', '\n':
		//End of comment
		//^ Was that meta or what?
		return lexStart, nil
	default:
		//Drop all input until comment end
		return lexLineComment, nil
	}
}

func lexOp(tokens chan<- Token, t rune, part []rune) (lexfunc, []rune) {
	switch t {
	case ' ':
		fallthrough
	case '\t':
		tokens <- Token{operation, string(part)}
		return lexParam, nil

	case '\n':
		tokens <- Token{operation, string(part)}
		return lexStart, nil

	case ':':
		//Was really a label all along
		tokens <- Token{label, string(part)}
		return lexStart, nil
	default:
		return lexOp, append(part, t)
	}
}

func lexParam(tokens chan<- Token, t rune, part []rune) (lexfunc, []rune) {
	switch t {
	case '#', '%':
		return lexLineComment, nil
	case ' ', '\t':
		return lexParam, nil

	case '\n':
		return lexStart, nil
	case 'r':
		return lexRegister, append(part, t)
	case '\'':
		return lexCharLit, nil

	case '-':
		return lexNegNumberLit, append(part, t)
	case '0':
		return lexRadix, append(part, t)
	case '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return lexNumberLit, append(part, t)
	default:
		tokens <- Token{unknown, string(append(part, t))}
		return lexParam, nil
	}
}

func lexRegister(tokens chan<- Token, t rune, part []rune) (lexfunc, []rune) {
	switch t {
	case ' ', '\t':
		tokens <- Token{register, string(part)}
		return lexParam, nil
	case '\n':
		tokens <- Token{register, string(part)}
		return lexStart, nil
	default:
		return lexRegister, append(part, t)
	}
}

var escaped = false

func lexCharLit(tokens chan<- Token, t rune, part []rune) (lexfunc, []rune) {
	switch t {
	case '\\':
		escaped = !escaped
		return lexCharLit, append(part, t)
	case '\'':
		if !escaped {
			char, _, _, err := strconv.UnquoteChar(string(part), '\'')
			if err != nil {
				tokens <- Token{unknown, err.Error()}
				return lexStart, nil
			}
			tokens <- Token{number, strconv.Itoa(int(char))}
			return lexParam, nil
		}
		fallthrough
	default:
		escaped = false
		return lexCharLit, append(part, t)
	}
}

func lexNumberLit(tokens chan<- Token, t rune, part []rune) (lexfunc, []rune) {
	switch t {
	case ' ', '\t':
		tokens <- Token{number, string(part)}
		return lexParam, nil
	case '\n':
		tokens <- Token{number, string(part)}
		return lexStart, nil

	case '0', '1', '2', '3', '4', '5', '6', '7',
		'8', '9', 'A', 'a', 'B', 'b', 'C', 'c',
		'D', 'd', 'E', 'e', 'F', 'f':
		return lexNumberLit, append(part, t)
	default:
		tokens <- Token{unknown, string(append(part, t))}
		return lexParam, nil
	}
}

func lexRadix(tokens chan<- Token, t rune, part []rune) (lexfunc, []rune) {
	switch t {
	case ' ', '\t':
		tokens <- Token{number, string(part)}
		return lexParam, nil
	case '\n':
		tokens <- Token{number, string(part)}
		return lexStart, nil
	case 'x', 'o', 'b':
		return lexNumberLit, append(part, t)
	default:
		tokens <- Token{unknown, string(append(part, t))}
		return lexNumberLit, nil
	}
	return lexNumberLit, nil
}

func lexNegNumberLit(tokens chan<- Token, t rune, part []rune) (lexfunc, []rune) {
	switch t {
	case '0':
		return lexRadix, append(part, t)
	case '1', '2', '3', '4', '5', '6', '7',
		'8', '9':
		return lexNumberLit, append(part, t)
	}
	tokens <- Token{unknown, string(append(part, t))}
	return lexStart, nil
}
