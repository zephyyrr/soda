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
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)
	tokens := make(chan Token, 3)

	go func() {
		defer close(tokens)
		// Set the split function for the scanning operation.
		for scanner.Scan() {
			s := scanner.Text()
			//Check language features:

			if strings.HasSuffix(s, ":") {
				s = strings.TrimRight(s, ":")
				tokens <- Token{label, s[:len(s)-2]}
				continue
			}
			if _, err := MapOperation(s); err == nil {
				tokens <- Token{operation, s}
				continue
			}

			if strings.HasPrefix(s, "r") {
				tokens <- Token{register, s}
				continue
			}

			if _, err := strconv.Atoi(s); err == nil {
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
