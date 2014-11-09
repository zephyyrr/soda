package aspartasm

import (
	"io"
	"bytes"
	"testing"
)

var instTable = map[string]Inst{
	"HALT":          Inst{HALT, nil},
	"ADD r1 r1 r2":  Inst{ADD, []Arg{Reg(1), Reg(1), Reg(2)}},
	"MALC r231 r17": Inst{MALC, []Arg{Reg(231), Reg(17)}},
	"PRNC 10":       Inst{PRNC, []Arg{Imm(10)}},
	"JMP r32":       Inst{JMP, []Arg{Reg(32)}},
	"PRNCI 10":      Inst{PRNCI, []Arg{Imm(10)}},
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

type readTestStruct struct {
	name string
	bin []byte
	err error
	ins []string
}

var binaryTable = []readTestStruct {
	{
		name: "Empty slice",
		bin: []byte{},
		err: io.EOF,
		ins: []string{},
	},
	{
		name: "Just magic bytes",
		bin: []byte{
			0x53, 0x4F, 0x44, 0x41,
		},
		err: nil,
		ins: []string{},
	},
	{
		name: "add instruction",
		bin: []byte{
			0x53, 0x4F, 0x44, 0x41,
			byte(ADD), 1, 13, 37,
		},
		err: nil,
		ins: []string{
			"ADD r1 r13 r37",
		},
	},
	{
		name: "print char immediate",
		bin: []byte{
			0x53, 0x4F, 0x44, 0x41,
			byte(PRNCI), 0, 0, 10,
		},
		err: nil,
		ins: []string{
			"PRNCI 10",
		},
	},
}

func TestReadInstructions(t *testing.T) {
	f := "ReadInstructions failed for %s:\n"

	for _, s := range binaryTable {
		ins, err := ReadInstructions(bytes.NewBuffer(s.bin))

		switch {
		case err != nil:
			if (s.err == nil) {
				t.Errorf(f + "didn't expect error, got %s", s.name, err)
			} else if (s.err != err) {
				t.Errorf(f + "%s != %s", s.name, err, s.err)
			}

		case len(ins) != len(s.ins):
			t.Error(f + "read %d instructions expected %d", s.name, len(ins), len(s.ins))

		default:
			for i, in := range ins {
				if in.String() != s.ins[i] {
					t.Errorf(f + "ins[%d] was read as %s, expected %s", s.name, i, in.String(), s.ins[i])
				}
			}
		}
	}
}
