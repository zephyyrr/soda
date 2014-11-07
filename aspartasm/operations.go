package aspartasm

import (
	"bufio"
	"encoding/binary"
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

var opToArgReader = map[Op]argReader{
	HALT: noRegs,

	NRS:  noRegs,
	PRS:  noRegs,
	MOVS: threeRegs,

	ZERO: oneReg,
	ADD:  threeRegs,
	SUB:  threeRegs,
	MUL:  threeRegs,
	DIV:  threeRegs,
	MOD:  threeRegs,
	POW:  threeRegs,

	ADDU: threeRegs,
	SUBU: threeRegs,
	MULU: threeRegs,
	DIVU: threeRegs,
	MODU: threeRegs,
	POWU: threeRegs,

	AND: threeRegs,
	OR:  threeRegs,
	XOR: threeRegs,
	INV: threeRegs,
	LSH: threeRegs,
	RSH: threeRegs,

	MALC: twoRegs,
	LDW:  threeRegs,
	LDB:  threeRegs,
	LDI:  oneRegImm,
	STW:  threeRegs,
	STB:  threeRegs,
	FREE: oneReg,

	JMP:   oneReg,
	JMPE:  threeRegs,
	JMPN:  threeRegs,
	JMPL:  threeRegs,
	JMPLE: threeRegs,
	BRA:   oneReg,
	BRAE:  threeRegs,
	BRAN:  threeRegs,
	BRAL:  threeRegs,
	BRALE: threeRegs,

	PRNI: oneReg,
	PRNC: oneReg,
}

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

func (op Op) ReadArgs(in *bufio.Reader) ([]Arg, error) {
	reader, ok := opToArgReader[op]

	if !ok {
		return nil, ErrInvalidOperation
	}

	return reader(in)
}

type argReader func(*bufio.Reader) ([]Arg, error)

func noRegs(in *bufio.Reader) ([]Arg, error) {
	for i := 0; i < 3; i++ {
		if _, err := in.ReadByte(); err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func oneReg(in *bufio.Reader) ([]Arg, error) {
	r1, err := in.ReadByte()

	if err != nil {
		return nil, err
	}

	for i := 0; i < 2; i++ {
		if _, err := in.ReadByte(); err != nil {
			return nil, err
		}
	}

	return []Arg{Reg(r1)}, nil
}

func twoRegs(in *bufio.Reader) ([]Arg, error) {
	r1, err := in.ReadByte()

	if err != nil {
		return nil, err
	}

	r2, err := in.ReadByte()

	if err != nil {
		return nil, err
	}

	if _, err := in.ReadByte(); err != nil {
		return nil, err
	}

	return []Arg{Reg(r1), Reg(r2)}, nil
}

func threeRegs(in *bufio.Reader) ([]Arg, error) {
	r1, err := in.ReadByte()

	if err != nil {
		return nil, err
	}

	r2, err := in.ReadByte()

	if err != nil {
		return nil, err
	}

	r3, err := in.ReadByte()

	if err != nil {
		return nil, err
	}

	return []Arg{Reg(r1), Reg(r2), Reg(r3)}, nil
}

func oneRegImm(in *bufio.Reader) ([]Arg, error) {
	r1, err := in.ReadByte()

	if err != nil {
		return nil, err
	}

	var imm uint32
	if err := binary.Read(in, binary.BigEndian, &imm); err != nil {
		return nil, err
	}

	return []Arg{Reg(r1), Imm(imm)}, nil
}
