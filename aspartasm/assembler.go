package aspartasm

import (
	"encoding/binary"
	"github.com/zephyyrr/soda"
	"io"
)

func Assemble(in io.Reader, out io.Writer) error {
	tokens := lex(in)
	ast, err := parse(tokens)
	if err != nil {
		return err
	}

	instructions, err := linearize(ast)

	binary.Write(out, binary.BigEndian, soda.MagicBytes)

	for _, instruction := range instructions {
		binary.Write(out, binary.BigEndian, instruction)
	}

	return nil
}

func parse(<-chan token) (tree ast, err error) {

	return
}

func linearize(tree ast) (ins []soda.Instruction, err error) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				err = r
			} else {
				panic(r)
			}
		}
	}()
	for _, ch := range tree.children {
		ins = append(ins, linearize_rec(tree, &soda.Instruction{}, 0))
	}
}

func linearize_rec(tree ast, curr *soda.Instruction, i int) soda.Instruction {
	switch tree.token.kind {
	case operation:
		op, err := MapOperation(tree.token.value)
		if err != nil {
			panic(err)
		}
		curr = &soda.Instruction{op, 0, 0, 0}
		for i, ch := range tree.children {
			linearize_rec(ch, curr, i)
		}

	}
}
