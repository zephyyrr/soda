package aspartasm

import (
	"encoding/binary"
	"errors"
	"github.com/zephyyrr/soda"
	"io"
	"strconv"
	"strings"
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
	tree.append(token{operation, "LDI"}).appendAll(
		token{register, "r0"},
		token{number, "257"})

	tree.append(token{operation, "LDI"}).
		appendAll(token{register, "r1"},
		token{number, "1"})

	tree.append(token{operation, "ADD"}).
		appendAll(token{register, "r2"},
		token{register, "r0"},
		token{register, "r1"})

	tree.append(token{operation, "PRNI"}).
		append(token{register, "r2"})
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
		in, err := linearize_rec(ch, soda.Instruction{}, 0)
		if err != nil {
			return ins, err
		}
		ins = append(ins, in)
	}

	return
}

func linearize_rec(tree ast, curr soda.Instruction, i int) (soda.Instruction, error) {
	var err error
	switch tree.token.kind {
	case operation:
		var err error
		curr.Operation, err = MapOperation(tree.token.value)
		if err != nil {
			return curr, err
		}
		for i, ch := range tree.children {
			curr, err = linearize_rec(ch, curr, i)
			if err != nil {
				return curr, err
			}
		}
		return curr, nil

	case register:
		switch i {
		case 0:
			curr.A, err = RegisterLookup(tree.value)
		case 1:
			curr.B, err = RegisterLookup(tree.value)
		case 2:
			curr.C, err = RegisterLookup(tree.value)
		}
		return curr, err
	case number:
		//Place number in slots B and C
		num, _ := strconv.Atoi(tree.token.value)
		curr.B = byte(num >> 8)
		curr.C = byte(num)
	}
	return curr, err
}

func RegisterLookup(r string) (byte, error) {
	if strings.HasPrefix(r, "r") {
		//Proper register
		num := strings.TrimPrefix(r, "r")
		register, err := strconv.Atoi(num)
		return byte(register), err
	}
	return 0, errors.New("Unknown register " + r)
}
