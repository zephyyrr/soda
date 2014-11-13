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
	ast, err := Parse(in)
	if err != nil {
		return err
	}
	return AssembleAst(ast, out)
}

func AssembleAst(tree AST, out io.Writer) error {
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

func Linearize(tree AST) (ins []soda.Instruction, err error) {
	defer func() {
		if r := recover(); r != nil {
			if err2, ok := r.(error); ok {
				err = err2
			} else {
				panic(r)
			}
		}
	}()
	for _, ch := range tree.Children {
		in, err := linearize_rec(ch, soda.Instruction{}, 0)
		if err != nil {
			return ins, err
		}
		ins = append(ins, in)
	}

	return
}

func linearize_rec(tree AST, curr soda.Instruction, i int) (soda.Instruction, error) {
	var err error
	switch tree.Token.Kind {
	case operation:
		var err error
		curr.Operation, err = MapOperation(tree.Token.Value)
		if err != nil {
			return curr, err
		}
		for i, ch := range tree.Children {
			curr, err = linearize_rec(ch, curr, i)
			if err != nil {
				return curr, err
			}
		}
		return curr, nil

	case register:
		switch i {
		case 0:
			curr.A, err = RegisterLookup(tree.Value)
		case 1:
			curr.B, err = RegisterLookup(tree.Value)
		case 2:
			curr.C, err = RegisterLookup(tree.Value)
		}
		return curr, err
	case number:
		if curr.Operation == byte(LDIH) {
			//Loading the high bits.
			num, _ := strconv.Atoi(tree.Token.Value)
			curr.B = byte(num >> 24)
			curr.C = byte(num >> 16)
		} else {
			//Place number in slots B and C
			num, _ := strconv.Atoi(tree.Token.Value)
			curr.B = byte(num >> 8)
			curr.C = byte(num)
		}
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
