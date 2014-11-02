package aspartasm

import (
	"encoding/binary"
	"github.com/zephyyrr/soda"
	"io"
	"strconv"
)

func Assemble(in io.Reader, out io.Writer) error {
	tokens := lex(in)
	ast, err := parse(tokens)
	if err != nil {
		return err
	}

	return AssembleAst(ast, out)
}

func AssembleAst(tree ast, out io.Writer) error {
	instructions, err := Linearize(tree)
	if err != nil {
		return err
	}

	binary.Write(out, binary.BigEndian, soda.MagicBytes)

	for _, instruction := range instructions {
		binary.Write(out, binary.BigEndian, instruction)
	}

	return nil
}

func parse(<-chan token) (tree ast, err error) {

	return
}

func Linearize(tree ast) (ins []soda.Instruction, err error) {
	defer func() {
		if r := recover(); r != nil {
			if err2, ok := r.(error); ok {
				err = err2
			} else {
				panic(r)
			}
		}
	}()
	for _, ch := range tree.children {
		ins = append(ins, linearize_rec(ch, soda.Instruction{}, 0))
	}

	return
}

func linearize_rec(tree ast, curr soda.Instruction, i int) soda.Instruction {
	switch tree.token.kind {
	case operation:
		var err error
		curr.Operation, err = MapOperation(tree.token.value)
		if err != nil {
			panic(err)
		}
		for i, ch := range tree.children {
			curr = linearize_rec(ch, curr, i)
		}
		return curr

	case register:
		switch i {
		case 0:
			curr.A = 0
		case 1:
			curr.B = 0
		case 2:
			curr.C = 0
		}
		return curr
	case number:
		//Place number in slots B and C
		num, _ := strconv.Atoi(tree.token.value)
		curr.B = byte(num >> 8)
		curr.C = byte(num)
	}
	return curr
}
