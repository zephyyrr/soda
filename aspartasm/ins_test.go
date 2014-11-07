package aspartasm

import (
	"testing"
)

var instTable = map[string]Inst{
	"HALT": Inst{Op(0), nil},
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
