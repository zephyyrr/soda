package aspartasm

import (
	"fmt"
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
		if err != nil {
			return
		}
	}
	if !done {
		//Did not stop in end-state
		err = io.EOF
	}
	return
}

func parseStart(tree *AST, token Token) (parseFunc, bool, error) {
	switch token.Kind {
	case operation:
		tree.append(token)
		return parseOperation, false, nil
	}
	println("start unknown")
	return nil, true, UnexpectedToken(token)
}

func parseOperation(tree *AST, token Token) (parseFunc, bool, error) {
	switch token.Kind {
	case register:
		fallthrough
	case number:
		tree.Children[len(tree.Children)-1].append(token)
		return parseOperation, true, nil
	default:
		return parseStart(tree, token)
	}
	println("operation unknown")
	return nil, true, UnexpectedToken(token)
}

type UnexpectedToken Token

func (ut UnexpectedToken) Error() string {
	return fmt.Sprintf("Unexpected token %q (%s) found during parse.", ut.Value, ut.Kind)
}
