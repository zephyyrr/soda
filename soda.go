package soda

import (
	"encoding/binary"
	"io"
)

const wordsize = 32

type word uint32

type register word

type InstructionSet func(byte) Operation

type Operation func(v *vm, a, b, c byte) error

type Instruction struct {
	Operation, A, B, C byte
}

type tape interface {
	io.Reader
	io.Seeker
	io.Closer
}

type registerSet [256]register

type vm struct {
	regsets [256]registerSet
	currset byte
	regs    *registerSet
	is      InstructionSet
	code    tape
	halting bool
}

func New(code tape) *vm {
	v := &vm{
		code:    code,
		is:      sodaIS,
		halting: false,
	}
	v.regs = &v.regsets[v.currset]
	return v
}

func (v *vm) Execute() error {
	if !v.verify() {
		return InvalidCode
	}

	var ins Instruction
	for !v.halting {
		err := binary.Read(v.code, binary.BigEndian, &ins)
		if err != nil {
			return err
		}
		// log.Printf("%x", ins)
		operation := v.is(ins.Operation)

		if err := operation(v, ins.A, ins.B, ins.C); err != nil {
			return err
		}
	}

	return nil
}

const MagicBytes word = 0x534F4441 // "SODA"

func (v *vm) verify() bool {
	var w word

	err := binary.Read(v.code, binary.BigEndian, &w)
	if err != nil {
		return false
	}

	return w == MagicBytes
}
