package aspartasm

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/zephyyrr/soda"
	"io"
)

var (
	ErrNoMagicBytes     = errors.New("No magic bytes")
	ErrInvalidOperation = errors.New("Invalid operation")
)

// A Soda operation argument
type Arg fmt.Stringer

// An Inst is a single instruction.
type Inst struct {
	Op   Op    // Opcode mnemonic
	Args []Arg // Instruction arguments
}

func (i Inst) String() string {
	var buf bytes.Buffer
	buf.WriteString(i.Op.String())

	for _, arg := range i.Args {
		if arg == nil {
			break
		}

		buf.WriteString(" ")
		buf.WriteString(arg.String())
	}

	return buf.String()
}

// A register
type Reg byte

func (r Reg) String() string {
	return fmt.Sprintf("r%d", r)
}

// An immediate value
type Imm int16

func (i Imm) String() string {
	return fmt.Sprintf("%d", i)
}

// ReadInstructions reads a sequence of Soda instructions from an io.Reader.
// The first 4 bytes read must be the Soda magic bytes.
func ReadInstructions(raw io.Reader) ([]Inst, error) {
	in := bufio.NewReader(raw)

	var magic uint32 //soda.word

	if err := binary.Read(in, binary.BigEndian, &magic); err != nil {
		println("Magic")
		return nil, err
	}

	// Ensure magic bytes
	if magic != uint32(soda.MagicBytes) {
		return nil, ErrNoMagicBytes
	}

	instrs := make([]Inst, 0)

	// Loop while not EOF
	for {
		opCode, err := in.ReadByte()

		switch err {
		case nil:
			break
		case io.EOF:
			return instrs, nil
		default:
			return nil, err
		}

		args, err := Op(opCode).ReadArgs(in)

		if err != nil {
			println("args")
			return nil, err
		}

		instrs = append(instrs, Inst{Op(opCode), args})
	}
}
