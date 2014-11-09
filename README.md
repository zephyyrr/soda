Byte code interpreted language with potential JIT-compilation.
Instruction set is register based with 256 standard registers (Oh yeah!).

[![Build Status](https://travis-ci.org/zephyyrr/soda.svg)](https://travis-ci.org/zephyyrr/soda)

# Assembler
The official soda assembler is of course aspartASMe.

# Instructions
Each instruction is made up on the following format:

ooooooooaaaaaaaabbbbbbbbcccccccc

or:

ooooooooaaaaaaaaiiiiiiiiiiiiiiii

depending on the operation.
Most operations use the first option, but the immediate operations use the latter.

## Legend
	o: operation identifier
	a: first register identifier
	b: second register identifier
	c: third register identifier
	i: immediate argument

# Instruction set
The instruction set is organised into parts by their first nibble.
The 0x0X part for example contains VM directives to control the registers and halt execution,
while 0x1X contains arithmetic operations and so on.

