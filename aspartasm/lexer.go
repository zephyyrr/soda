package aspartasm

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type token struct {
	kind  string
	value string
}

const (
	unknown   = "wft?"
	label     = "LABEL"
	operation = "operation"
	register  = "register"
	number    = "number"
	sträng    = "string"
)

func lex(in io.Reader) <-chan token {
	bufd := bufio.NewReader(in)
	tokens := make(chan token, 3)

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
				tokens <- token{label, s[:len(s)-2]}
				continue
			}
			if _, err := MapOperation(s); err != nil {
				tokens <- token{operation, s}
				continue
			}

			if strings.HasPrefix(s, "r") {
				tokens <- token{register, s}
				continue
			}

			if _, err := strconv.Atoi(s); err != nil {
				tokens <- token{number, s}
			}

			if strings.HasPrefix(s, "\"") &&
				strings.HasSuffix(s, "\"") {
				tokens <- token{sträng, s[1 : len(s)-2]}
				continue
			}

			tokens <- token{unknown, s}
		}
	}()

	return tokens
}
