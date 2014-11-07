package aspartasm

import (
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
		t.Errorf("opToString doesn't work")
	}

	for k, v := range instTable {
		if v.String() != k {
			t.Errorf("Expected %s to equal %s\n", v, k)
		}
	}
}
