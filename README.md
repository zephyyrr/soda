Byte code interpreted language with potential JIT-compilation.
Instruction set is register based with 256 standard registers (Oh yeah!).

# Assembler
The official soda assembler is of course aspartASMe.

# Instructions
Each instruction is made up on the following format:

iiiiiiiiaaaaaaaabbbbbbbbcccccccc

## Legend
	i: instruction identifier
	a: first register identifier
	b: second register identifier
	c: third register identifier

# Instruction set
## 0X - Special reserved
	00 - Halt execution
	01 -
	0A - Next Register set
	0B - Previous register set
	0C - Clone register (from one set to another)

## 1X - Arithmetic Operations
	10 - Times zero
	11 - Addition
	12 - Subtraction
	13 - Multiplication
	14 - Division
	15 - Power

	19 - Unsigned Addition
	1A - Unsigned Subtraction
	1B - Unsigned Multiplication
	1C - Unsigned Division
	1D - Unsigned Power

## 3X - Bitwise operations
	31 - And
	32 - Or
	33 - Xor
	34 - Invert
	35 - Shift left
	36 - Shift right

## 5X - Memory Operations
	51 - Load
	52 - Load byte
	53 - Load Immediate
	56 - Store
	57 - Store byte

## 6X - Jumps
	61 - uncoditional jump
	62 - jump if equals
	63 - jump if not equals
	64 - jump if less than
	65 - jump if less than or equals

	68 - branch
	69 - branch if equals
	6A - branch if not equals
	6B - branch if less than
	6C - branch if less than or equals


## 8X - IO Operations
	81 - Print register a as integer
	82 - Print register a as char
##
