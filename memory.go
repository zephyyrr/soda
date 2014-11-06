package soda

import (
	"encoding/binary"
	"fmt"
)

type MemoryBlock []byte
type address word

type MainMemory map[address]MemoryBlock

func (m MainMemory) LoadWord(addr address, offset word) (word, error) {
	if uint32(len(m[addr])) <= uint32(offset+4) {
		return 0, IllegalMemoryAccess(addr + address(offset))
	}
	val, _ := binary.Uvarint(m[addr][offset : offset+4])
	return word(val), nil
}

func (m MainMemory) LoadByte(addr address, offset word) (byte, error) {
	if uint32(len(m[addr])) <= uint32(offset) {
		return 0, IllegalMemoryAccess(addr + address(offset))
	}
	return m[addr][offset], nil
}

func (m MainMemory) StoreWord(addr address, offset word, value word) error {
	if uint32(len(m[addr])) <= uint32(offset+4) {
		return IllegalMemoryAccess(addr + address(offset))
	}
	binary.PutUvarint(m[addr][offset:offset+4], uint64(value))
	return nil
}

func (m MainMemory) StoreByte(addr address, offset word, value byte) error {
	if uint32(len(m[addr])) <= uint32(offset) {
		return IllegalMemoryAccess(addr + address(offset))
	}
	m[addr][offset] = value
	return nil
}

type IllegalMemoryAccess address

func (ima IllegalMemoryAccess) Error() string {
	return fmt.Sprintf("Illegal Memory Access @ %h", word(ima))
}
