package aspartasm

import (
	"bytes"
	"testing"
)

var instTable = map[string]Inst{
	"HALT":          Inst{HALT, nil},
	"ADD r1 r1 r2":  Inst{ADD, []Arg{Reg(1), Reg(1), Reg(2)}},
	"MALC r231 r17": Inst{MALC, []Arg{Reg(231), Reg(17)}},
	"PRNC 10":       Inst{PRNC, []Arg{Imm(10)}},
	"JMP r32":       Inst{JMP, []Arg{Reg(32)}},
}

func TestInstString(t *testing.T) {
	if opToString[0] != "HALT" {
		t.Error("opToString doesn't work")
	}

	for k, v := range instTable {
		if v.String() != k {
			t.Errorf("Expected %s to equal %s\n", v, k)
		}
	}
}

var (
	emptySlice = []byte{}
	justMagic  = []byte{0x53, 0x4F, 0x44, 0x41}
	singleAdd  = []byte{0x53, 0x4F, 0x44, 0x41, byte(ADD), 0x01, 0x13, 0x37}
)

func TestReadInstructions(t *testing.T) {
	_, err := ReadInstructions(bytes.NewBuffer(emptySlice))

	if err == nil {
		t.Error("ReadInstructions shouldn't succeed on empty buffer")
	}

	ins, err := ReadInstructions(bytes.NewBuffer(justMagic))

	if err != nil {
		t.Error("ReadInstructions shouldn't fail on just magic bytes")
	}

	if len(ins) != 0 {
		t.Error("ReadInstructions shouldn't read instructions from just magic bytes")
	}

	ins, err = ReadInstructions(bytes.NewBuffer(singleAdd))
}
