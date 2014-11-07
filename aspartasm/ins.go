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
	ErrNoMagicBytes = errors.New("No magic bytes")
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
type Imm uint32

func (i Imm) String() string {
	return fmt.Sprintf("%d", i)
}

// ReadInstructions reads a sequence of Soda instructions from an io.Reader.
// The first 4 bytes read must be the Soda magic bytes.
func ReadInstructions(raw io.Reader) ([]Inst, error) {
	in := bufio.NewReader(raw)

	var magic uint32

	if err := binary.Read(in, binary.BigEndian, &magic); err != nil {
		return nil, err
	}

	// Ensure magic bytes
	if magic != uint32(soda.MagicBytes) {
		return nil, ErrNoMagicBytes
	}

	instrs := make([]Inst)

	// Loop while not EOF
	for {
		opCode, err := in.ReadByte()

		switch err != nil; err {
		case EOF:
			return instrs, nil
		default:
			return nil, err
		}

		rdr := getArgumentReader(opCode)

		args, err := rdr(in)

		if err != nil {
			return nil, err
		}

		instrs = append(instrs, Inst{Op(opCode), args})
	}
}
