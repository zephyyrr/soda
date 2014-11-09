package soda

import (
	"testing"
)

func TestMainMemory_Allocate(t *testing.T) {
	//Create a MainMemory
	mm := make(MainMemory)
	//Allocate some memory blocks in various sizes
	for i := word(1); i < 257; i++ {
		addr, err := mm.Allocate(i)
		if err != nil {
			t.Error(err)
		}

		//Make sure blocks are in map
		if val, ok := mm[addr]; !ok || val == nil {
			t.Fail()
		}
	}
}

func TestMainMemory_Free(t *testing.T) {
	//Create MainMemory
	mm := make(MainMemory)

	//Create some memoryblocks and but them in the map.
	var addrs []address = nil
	for i := address(0); i < 257; i++ {
		mm[i+32456] = make(MemoryBlock, 1)
		addrs = append(addrs, i+32456)
	}

	defer func() {
		if r := recover(); r != nil {
			t.Fail()
		}
	}()

	//Call free on all blocks
	for _, addr := range addrs {
		err := mm.Free(addr)
		if err != nil {
			t.Error(err)
		}

		//Make sure blocks are not in map
		if _, ok := mm[addr]; ok {
			t.Fail()
		}

	}
}

func TestMainMemory_LoadWord(t *testing.T) {
	mm := make(MainMemory)
	mm[0] = MemoryBlock{1, 2, 3, 4, 5, 6}
	expected := []word{0x04030201, 0x05040302, 0x06050403}
	for i := 0; i < len(mm[0])-4; i++ {
		t.Log(i)
		w, err := mm.LoadWord(0, word(i))
		if err != nil {
			t.Error(err)
		}
		if w != expected[i] {
			t.Fail()
		}
	}
}
