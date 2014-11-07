package aspartasm

import (
	"fmt"
)

const (
	HALT Op = iota
	_
	_
	_
	_
	_
	_
	_
	_
	_
	NRS
	PRS
	MOVS
	_
	_
	_

	ZERO
	ADD
	SUB
	MUL
	DIV
	MOD
	POW
	_
	_
	ADDU
	SUBU
	MULU
	DIVU
	MODU
	POWU
	_

	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_

	_
	AND
	OR
	XOR
	INV
	LSH
	RSH
	_
	_
	_
	_
	_
	_
	_
	_
	_

	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_

	MALC
	LDW
	LDB
	LDI
	_
	_
	STW
	STB
	_
	_
	_
	_
	_
	_
	_
	FREE

	_
	JMP
	JMPE
	JMPN
	JMPL
	JMPLE
	_
	_
	BRA
	BRAE
	BRAN
	BRAL
	BRALE
	_
	_
	_

	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_

	_
	PRNI
	PRNC
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
	if v, ok := opToString[op]; ok {
		return v
	}

	return fmt.Sprintf("Op(%x)", byte(op))
}
