package aspartasm

import (
	"encoding/binary"
	"github.com/zephyyrr/soda"
	"io"
)

func Assemble(in io.Reader, out io.Writer) error {
	tokens := lex(in)
	instructions, err := parse(tokens)
	if err != nil {
		return err
	}

	binary.Write(out, binary.BigEndian, soda.MagicBytes)

	for _, instruction := range instructions {
		binary.Write(out, binary.BigEndian, instruction)
	}

	return nil
}

func parse(<-chan token) (ins []soda.Instruction, err error) {
	ins = append(ins, soda.Instruction{0x53, 0, 0, 65})
	ins = append(ins, soda.Instruction{0x82, 0, 0, 0})
	return
}
