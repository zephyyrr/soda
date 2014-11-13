package aspartasm

import (
	"errors"
	"strings"
)

var ops = map[string]byte{
	"HALT": 0x00,
	"BRKP": 0x01,
	"MOV":  0x02,
	"NRS":  0x0A,
	"PRS":  0x0B,
	"MOVS": 0x0C,

	"ZERO": 0x10,
	"ADD":  0x11,
	"SUB":  0x12,
	"MUL":  0x13,
	"DIV":  0x14,
	"POW":  0x15,

	"ADDU": 0x19,
	"SUBU": 0x1A,
	"MULU": 0x1B,
	"DIVU": 0x1C,
	"POWU": 0x1D,

	"ADDI": 0x21,
	"SUBI": 0x22,
	"MULI": 0x23,
	"DIVI": 0x24,
	"POWI": 0x25,

	"ADDUI": 0x29,
	"SUBUI": 0x2A,
	"MULUI": 0x2B,
	"DIVUI": 0x2C,
	"POWUI": 0x2D,

	"AND": 0x31,
	"OR":  0x32,
	"XOR": 0x33,
	"LSH": 0x35,
	"RSH": 0x36,

	"MALC": 0x50,
	"LDW":  0x51,
	"LDB":  0x52,
	"LDI":  0x53,
	"LDIL": 0x54,
	"LDIH": 0x55,
	"STW":  0x56,
	"STB":  0x57,
	"FREE": 0x5F,

	"JMP":   0x61,
	"JMPE":  0x62,
	"JMPN":  0x63,
	"JMPL":  0x64,
	"JMPLE": 0x65,

	"BRA":   0x68,
	"BRAE":  0x69,
	"BRAN":  0x6A,
	"BRAL":  0x6B,
	"BRALE": 0x6C,

	"PRNI":  0x81,
	"PRNC":  0x82,
	"PRNII": 0x89,
	"PRNCI": 0x8a,
}

func MapOperation(op string) (byte, error) {
	if b, ok := ops[strings.ToUpper(op)]; ok {
		return b, nil
	}
	return 0, UnknownOperation(op)
}

func UnknownOperation(op string) error {
	return errors.New("Unknown operation")
}
