package aspartasm

import (
	"io"
	"text/scanner"
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

type lexfunc func(chan<- Token, scanner.Scanner, []rune) (lexfunc, []rune)

func lex(in io.Reader) <-chan Token {
	var s scanner.Scanner
	s.Init(in)
	tokens := make(chan Token, 3)

	go func() {
		defer close(tokens)
		defer func() {
			if r := recover(); r != nil {
				//Probably EOF at bad time
			}
		}()

		// Set the split function for the scanning operation.
		f := lexfunc(lexStart)
		for f != nil {

		}

	}()

	return tokens
}

func lexStart(tokens chan<- Token, s scanner.Scanner, part []rune) (lexfunc, []rune) {
	t := s.Scan()
	switch t {
	case scanner.EOF:
		return nil, nil
	case ' ':
		fallthrough
	case '\n':
		fallthrough
	case '\t':
		return lexStart, nil
	case '%':
		fallthrough
	case '#':
		return lexLineComment, nil
	default:
		return lexOp, append(part, t)
	}
}

func lexLineComment(tokens chan<- Token, s scanner.Scanner, part []rune) (lexfunc, []rune) {
	t := s.Scan()
	switch t {
	case '\n':
		//End of comment
		//^ Was that meta or what?
		return lexStart, nil
	default:
		//Drop all input until comment end
		return lexLineComment, nil
	}
}

func lexOp(tokens chan<- Token, s scanner.Scanner, part []rune) (lexfunc, []rune) {
	t := s.Scan()
	switch t {
	case scanner.EOF:
		tokens <- Token{operation, string(part)}
		return nil, nil
	case ' ':
		fallthrough
	case '\t':
		tokens <- Token{operation, string(part)}
		return lexParam, nil

	case '\n':
		return lexStart, nil

	case ':':
		//Was really a label all along
		tokens <- Token{label, string(part)}
		return lexStart, nil
	default:
		return lexOp, append(part, t)
	}
}

func lexParam(tokens chan<- Token, s scanner.Scanner, part []rune) (lexfunc, []rune) {
	t := s.Scan()
	switch t {
	case scanner.EOF:
		return nil, nil
	case '#':
		fallthrough
	case '%':
		return lexLineComment, nil
	case ' ':
		fallthrough
	case '\t':
		return lexParam, nil

	case '\n':
		return lexStart, nil
	case 'r':
		return lexRegister, nil
	case '\'':
		return lexCharLit, nil

	case '0':
		return lexRadix, nil
	case '1':
		fallthrough
	case '2':
		fallthrough
	case '3':
		fallthrough
	case '4':
		fallthrough
	case '5':
		fallthrough
	case '6':
		fallthrough
	case '7':
		fallthrough
	case '8':
		fallthrough
	case '9':
		return lexNumberLit, append(part, t)
	default:
		tokens <- Token{unknown, string(append(part, t))}
		return lexParam, nil
	}
}

func lexRegister(tokens chan<- Token, s scanner.Scanner, part []rune) (lexfunc, []rune) {
	t := s.Scan()
	switch t {
	case scanner.EOF:
		tokens <- Token{operation, string(part)}
		return nil, nil
	case ' ':
		fallthrough
	case '\t':
		tokens <- Token{register, string(part)}
		return lexParam, nil
	default:
		tokens <- Token{unknown, string(append(part, t))}
		return lexRegister, append(part, t)
	}
}
