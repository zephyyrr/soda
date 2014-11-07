package aspartasm

import (
	"bytes"
	"fmt"
)

var opToString = make(map[Op]string)

func init() {
	for k, v := range ops {
		opToString[Op(v)] = k
	}
}

// A Soda opcode
type Op byte

func (op Op) String() string {
	return opToString[op]
}

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

type Reg byte

func (r Reg) String() string {
	return fmt.Sprintf("r%d", r)
}

type Imm uint32

func (i Imm) String() string {
	return fmt.Sprintf("%d", i)
}
