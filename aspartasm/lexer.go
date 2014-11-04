package aspartasm

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Token struct {
	Kind  string
	Value string
}

const (
	unknown   = "wft?"
	label     = "LABEL"
	operation = "operation"
	register  = "register"
	number    = "number"
	sträng    = "string"
)

func lex(in io.Reader) <-chan Token {
	bufd := bufio.NewReader(in)
	tokens := make(chan Token, 3)

	go func() {
		for {
			s, err := bufd.ReadString(' ')
			if err != nil {
				close(tokens)
			}

			s = strings.Trim(s, " \n\t")

			//Check language features:

			if strings.HasSuffix(s, ":") {
				s = strings.TrimRight(s, ":")
				tokens <- Token{label, s[:len(s)-2]}
				continue
			}
			if _, err := MapOperation(s); err != nil {
				tokens <- Token{operation, s}
				continue
			}

			if strings.HasPrefix(s, "r") {
				tokens <- Token{register, s}
				continue
			}

			if _, err := strconv.Atoi(s); err != nil {
				tokens <- Token{number, s}
			}

			if strings.HasPrefix(s, "\"") &&
				strings.HasSuffix(s, "\"") {
				tokens <- Token{sträng, s[1 : len(s)-2]}
				continue
			}

			tokens <- Token{unknown, s}
		}
	}()

	return tokens
}
