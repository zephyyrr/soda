package aspartasm

import (
	"github.com/zephyyrr/soda"
	"testing"
)

func TestLinearize(t *testing.T) {
	tree := newAst()
	tree.append(Token{operation, "LDI"}).appendAll(
		Token{register, "r0"},
		Token{number, "257"})

	tree.append(Token{operation, "LDI"}).
		appendAll(Token{register, "r1"},
		Token{number, "1"})

	tree.append(Token{operation, "ADD"}).
		appendAll(Token{register, "r2"},
		Token{register, "r0"},
		Token{register, "r1"})

	tree.append(Token{operation, "PRNI"}).
		append(Token{register, "r2"})

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
