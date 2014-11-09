package soda

import (
	"fmt"
	"unsafe"
)

type MemoryBlock []byte
type address word

type MainMemory map[address]MemoryBlock

func (m MainMemory) Allocate(size word) (addr address, err error) {
	area := make(MemoryBlock, size, size)
	addr = address(uintptr(unsafe.Pointer(&area)))
	m[addr] = area
	return
}

func (m MainMemory) Free(addr address) error {
	delete(m, addr)
	return nil
}

func (m MainMemory) LoadWord(addr address, offset word) (word, error) {
	if uint32(len(m[addr])) < uint32(offset+4) {
		return 0, IllegalMemoryAccess(addr + address(offset))
	}
	chunk := m[addr]
	val := word(chunk[offset]) | word(chunk[offset+1])<<8
	val |= word(chunk[offset+2])<<16 | word(chunk[offset+3])<<24
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
	m[addr][offset] = byte(value)
	m[addr][offset+1] = byte(value >> 8)
	m[addr][offset+2] = byte(value >> 16)
	m[addr][offset+3] = byte(value >> 24)
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
