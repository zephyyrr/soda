package aspartasm

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
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
				continue
			}

			if strings.HasPrefix(s, "'") &&
				strings.HasSuffix(s, "'") && len(s) > 2 {
				val, _, _, _ := strconv.UnquoteChar(s[1:len(s)-1], '"')
				//Multibyte is dropped because we will drop extra bytes anyway shortly.
				//The tail is dropped because it should not be there
				//the error is dropped because then we just give a zero literal.
				//println("Char literal", s, "=", int(val), "detected:", err.Error())
				//println("Multi-byte:", multi, "; Tail:", tail)
				tokens <- Token{number, strconv.Itoa(int(val))}
				continue
			}

			if strings.HasPrefix(s, "'") &&
				strings.HasSuffix(s, "'") {
				var str string
				fmt.Sscanf(s, "%q", &str)
				tokens <- Token{sträng, str}
				continue
			}

			tokens <- Token{unknown, s}
		}
	}()

	return tokens
}
