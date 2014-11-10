package aspartasm

import (
	"io"
	"strconv"
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
	case scanner.EOF:
		return nil, nil
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
		return lexRadix, append(part, t)
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
		panic(t)
	case ' ':
		fallthrough
	case '\t':
		tokens <- Token{register, string(part)}
		return lexParam, nil
	default:
		return lexRegister, append(part, t)
	}
}

func lexCharLit(tokens chan<- Token, s scanner.Scanner, part []rune) (lexfunc, []rune) {
	t := s.Scan()
	switch t {
	case scanner.EOF:
		panic(t)
	case '\'':
		char, _, _, err := strconv.UnquoteChar(string(part), '\'')
		if err != nil {
			tokens <- Token{unknown, string(append(part, t))}
		}
		tokens <- Token{register, strconv.Itoa(int(char))}
		return lexParam, nil
	default:
		tokens <- Token{unknown, string(append(part, t))}
		return lexCharLit, append(part, t)
	}
}

func lexNumberLit(tokens chan<- Token, s scanner.Scanner, part []rune) (lexfunc, []rune) {
	t := s.Scan()
	switch t {
	case scanner.EOF:
		panic(t)
	case ' ':
		fallthrough
	case '\t':
		tokens <- Token{number, string(part)}
		return lexParam, nil

	case '0':
		return lexRadix, append(part, t)
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
		fallthrough
	case 'A':
		fallthrough
	case 'a':
		fallthrough
	case 'B':
		fallthrough
	case 'b':
		fallthrough
	case 'C':
		fallthrough
	case 'c':
		fallthrough
	case 'D':
		fallthrough
	case 'd':
		fallthrough
	case 'E':
		fallthrough
	case 'e':
		fallthrough
	case 'F':
		fallthrough
	case 'f':
		return lexNumberLit, append(part, t)
	default:
		tokens <- Token{unknown, string(append(part, t))}
		return lexParam, nil
	}
}

func lexRadix(tokens chan<- Token, s scanner.Scanner, part []rune) (lexfunc, []rune) {
	t := s.Scan()
	switch t {
	case scanner.EOF:
		panic(t)
	case 'x':
	case 'o':
	case 'b':
		return lexNumberLit, append(part, t)
	default:
		tokens <- Token{unknown, string(append(part, t))}
		return lexNumberLit, nil
	}
	return lexNumberLit, nil
}
