package aspartasm

import (
	"io"
)

func Parse(in io.Reader) (tree AST, err error) {
	tokens := lex(in)
	tree, err = parse(tokens)
	return
}

type parseFunc func(*AST, Token) (parseFunc, bool, error)

func parse(tokens <-chan Token) (tree AST, err error) {
	var rule parseFunc = parseStart
	done := false

	for token := range tokens {
		rule, done, err = rule(&tree, token)
	}
	if !done {
		//Did not stop in end-state
		err = io.EOF
	}
	return
}

func parseStart(tree *AST, token Token) (parseFunc, bool, error) {
	return nil, true, nil
}
