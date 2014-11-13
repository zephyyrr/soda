package aspartasm

import (
	"strings"
	"testing"
	"time"
)

const testCode = `# Testcode.
main:
	LDI r0 1
	ldi r1 0x45
	ldi r3 0
	ldi r5 -28
	LdI r2 0b01010101

	% This is a comment
	# PRNCI 'S'
	PRNCI ' '
	PRNCI '\''
	PRNCI '\n'
	ADD r0 r0 r2
	HALT 42
`

func TestLexer(t *testing.T) {
	tokens := lex(strings.NewReader(testCode))

	expects := []Token{
		Token{label, "main"},
		Token{operation, "LDI"},
		Token{register, "r0"},
		Token{number, "1"},

		Token{operation, "ldi"},
		Token{register, "r1"},
		Token{number, "0x45"},

		Token{operation, "ldi"},
		Token{register, "r3"},
		Token{number, "0"},

		Token{operation, "ldi"},
		Token{register, "r5"},
		Token{number, "-28"},

		Token{operation, "LdI"},
		Token{register, "r2"},
		Token{number, "0b01010101"},

		Token{operation, "PRNCI"},
		Token{number, "32"},

		Token{operation, "PRNCI"},
		Token{number, "39"},

		Token{operation, "PRNCI"},
		Token{number, "10"},

		Token{operation, "ADD"},
		Token{register, "r0"},
		Token{register, "r0"},
		Token{register, "r2"},

		Token{operation, "HALT"},
		Token{number, "42"},
	}

	for _, expected := range expects {
		select {
		case <-time.After(time.Second):
			t.Fatal("Timeout")
		case tok, ok := <-tokens:
			if ok && tok == expected {
				continue
			}
			if !ok {
				t.Fatalf("Token channel is closed.")
			} else {
				t.Errorf("Expected %v, got %v", expected, tok)
			}
		}
	}

	if tok, ok := <-tokens; ok {
		t.Errorf("Recieved unexpected token %v", tok)
	}
}
