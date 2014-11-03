package aspartasm

import (
	"github.com/zephyyrr/soda"
	"testing"
)

func TestLinearize(t *testing.T) {
	tree := newAst()
	tree.append(token{operation, "LDI"}).appendAll(
		token{register, "r0"},
		token{number, "257"})

	tree.append(token{operation, "LDI"}).
		appendAll(token{register, "r1"},
		token{number, "1"})

	tree.append(token{operation, "ADD"}).
		appendAll(token{register, "r2"},
		token{register, "r0"},
		token{register, "r1"})

	tree.append(token{operation, "PRNI"}).
		append(token{register, "r2"})

	ins, err := Linearize(*tree)
	if err != nil {
		t.Error(err)
	}

	exps := []soda.Instruction{
		soda.Instruction{0x53, 0, 1, 0x01},
		soda.Instruction{0x53, 1, 0, 0x01},
		soda.Instruction{0x11, 2, 0, 1},
		soda.Instruction{0x81, 2, 0, 0},
	}

	if len(ins) != 4 {
		t.Errorf("Wrong size of resulting instruction list. %d != %d", len(ins), 4)
	}

	for i, exp := range exps {
		if ins[i] != exp {
			t.Errorf("Wrong instruction #%d: %d != %d", i, ins[i], exp)
		}
	}
}
