package soda

import (
	"fmt"
	"math"
)

func sodaIS(ins byte) Operation {
	switch ins {
	case 0x00:
		return halt
	case 0x01:
		return breakpoint
	case 0x02:
		return move
	case 0x0A:
		return nRegSet
	case 0x0B:
		return pRegSet
	case 0x0C:
		return cReg

	case 0x10:
		return zero
	case 0x11:
		return addition
	case 0x12:
		return subtraction
	case 0x13:
		return multiplication
	case 0x14:
		return division
	case 0x15:
		return mod
	case 0x16:
		return power

	case 0x19:
		return uaddition
	case 0x1A:
		return usubtraction
	case 0x1B:
		return umultiplication
	case 0x1C:
		return udivision
	case 0x1D:
		return umod
	case 0x1E:
		return upower

	case 0x21:
		return additionI
	case 0x22:
		return subtractionI
	case 0x23:
		return multiplicationI
	case 0x24:
		return divisionI
	case 0x25:
		return modI
	case 0x26:
		return powerI

	case 0x29:
		return uadditionI
	case 0x2A:
		return usubtractionI
	case 0x2B:
		return umultiplicationI
	case 0x2C:
		return udivisionI
	case 0x2D:
		return umodI
	case 0x2E:
		return upowerI

	case 0x31:
		return and
	case 0x32:
		return or
	case 0x33:
		return xor
	case 0x34:
		return invert
	case 0x35:
		return lshift
	case 0x36:
		return rshift

	case 0x50:
		return malloc
	case 0x51:
		return load
	case 0x52:
		return loadb
	case 0x53:
		return loadi
	case 0x56:
		return store
	case 0x57:
		return storeb
	case 0x5F:
		return free

	case 0x61:
		return jump
	case 0x68:
		return branch
	case 0x69:
		return branchEq
	case 0x6B:
		return branchLess

	case 0x81:
		return printi
	case 0x82:
		return printc
	case 0x89:
		return printii
	case 0x8a:
		return printci
	}

	return undefined
}

func undefined(v *vm, a, b, c byte) error {
	panic(UndefinedBehaviour)
}

func halt(v *vm, a, b, c byte) error {
	v.halting = true
	return nil
}

func nRegSet(v *vm, a, b, c byte) error {
	v.currset++
	v.regs = &v.regsets[v.currset]
	return nil
}

func pRegSet(v *vm, a, b, c byte) error {
	v.currset--
	v.regs = &v.regsets[v.currset]
	return nil
}

func cReg(v *vm, a, b, c byte) error {
	v.regsets[v.regs[a]][v.regs[c]] = v.regsets[v.regs[b]][v.regs[c]]
	return nil
}

func move(v *vm, a, b, c byte) error {
	v.regs[a] = v.regs[b]
	return nil
}

func breakpoint(v *vm, a, b, c byte) error {
	v.Break()
	return nil
}

func zero(v *vm, a, b, c byte) error {
	v.regs[a] = 0
	return nil
}

func addition(v *vm, a, b, c byte) error {
	v.regs[a] = register(int32(v.regs[b]) + int32(v.regs[c]))
	return nil
}

func subtraction(v *vm, a, b, c byte) error {
	v.regs[a] = v.regs[b] - v.regs[c]
	return nil
}

func multiplication(v *vm, a, b, c byte) error {
	v.regs[a] = v.regs[b] * v.regs[c]
	return nil
}

func division(v *vm, a, b, c byte) error {
	v.regs[a] = v.regs[b] / v.regs[c]
	return nil
}

func mod(v *vm, a, b, c byte) error {
	v.regs[a] = v.regs[b] % v.regs[c]
	return nil
}

func power(v *vm, a, b, c byte) error {
	v.regs[a] = register(math.Pow(float64(v.regs[b]), float64(v.regs[c])))
	return nil
}

func uaddition(v *vm, a, b, c byte) error {
	v.regs[a] = v.regs[b] + v.regs[c]
	return nil
}

func usubtraction(v *vm, a, b, c byte) error {
	v.regs[a] = v.regs[b] - v.regs[c]
	return nil
}

func umultiplication(v *vm, a, b, c byte) error {
	v.regs[a] = v.regs[b] * v.regs[c]
	return nil
}

func udivision(v *vm, a, b, c byte) error {
	v.regs[a] = v.regs[b] / v.regs[c]
	return nil
}

func umod(v *vm, a, b, c byte) error {
	v.regs[a] = v.regs[b] % v.regs[c]
	return nil
}

func upower(v *vm, a, b, c byte) error {
	v.regs[a] = register(math.Pow(float64(v.regs[b]), float64(v.regs[c])))
	return nil
}

/*
	Immediate artithmetic operations
*/

func additionI(v *vm, a, b, c byte) error {
	v.regs[a] += register(b)<<16 | register(c)
	return nil
}

func subtractionI(v *vm, a, b, c byte) error {
	v.regs[a] -= register(b)<<16 | register(c)
	return nil
}

func multiplicationI(v *vm, a, b, c byte) error {
	v.regs[a] *= register(b)<<16 | register(c)
	return nil
}

func divisionI(v *vm, a, b, c byte) error {
	v.regs[a] /= register(b)<<16 | register(c)
	return nil
}

func modI(v *vm, a, b, c byte) error {
	v.regs[a] %= register(b)<<16 | register(c)
	return nil
}

func powerI(v *vm, a, b, c byte) error {
	v.regs[a] = register(math.Pow(float64(v.regs[a]), float64(register(b)<<16|register(c))))
	return nil
}

func uadditionI(v *vm, a, b, c byte) error {
	v.regs[a] += register(b)<<16 | register(c)
	return nil
}

func usubtractionI(v *vm, a, b, c byte) error {
	v.regs[a] -= register(b)<<16 | register(c)
	return nil
}

func umultiplicationI(v *vm, a, b, c byte) error {
	v.regs[a] *= register(b)<<16 | register(c)
	return nil
}

func udivisionI(v *vm, a, b, c byte) error {
	v.regs[a] /= register(b)<<16 | register(c)
	return nil
}

func umodI(v *vm, a, b, c byte) error {
	v.regs[a] %= register(b)<<16 | register(c)
	return nil
}

func upowerI(v *vm, a, b, c byte) error {
	v.regs[a] = register(math.Pow(float64(v.regs[a]), float64(register(b)<<16|register(c))))
	return nil
}

/*
	Logical operations
*/

func and(v *vm, a, b, c byte) error {
	v.regs[a] = v.regs[b] & v.regs[c]
	return nil
}

func or(v *vm, a, b, c byte) error {
	v.regs[a] = v.regs[b] | v.regs[c]
	return nil
}

func xor(v *vm, a, b, c byte) error {
	v.regs[a] = v.regs[b] ^ v.regs[c]
	return nil
}

func invert(v *vm, a, b, c byte) error {
	v.regs[a] = ^v.regs[b]
	return nil
}

func lshift(v *vm, a, b, c byte) error {
	v.regs[a] = v.regs[b] << v.regs[c]
	return nil
}

func rshift(v *vm, a, b, c byte) error {
	v.regs[a] = v.regs[b] >> v.regs[c]
	return nil
}

func jump(v *vm, a, b, c byte) error {
	v.code.Seek(int64(v.regs[a]), 0)
	return nil
}

func branch(v *vm, a, b, c byte) error {
	v.code.Seek(int64(v.regs[a]), 0)
	return nil
}

func branchEq(v *vm, a, b, c byte) error {
	if v.regs[b] == v.regs[c] {
		v.code.Seek(int64(v.regs[a]), 0)
	}
	return nil
}

func branchLess(v *vm, a, b, c byte) error {
	if v.regs[b] < v.regs[c] {
		v.code.Seek(int64(v.regs[a]), 0)
	}
	return nil
}

func malloc(v *vm, a, b, c byte) error {
	addr, err := v.Allocate(word(v.regs[b]))
	v.regs[a] = register(addr)
	return err
}

func free(v *vm, a, b, c byte) error {
	return nil
}

func load(v *vm, a, b, c byte) (err error) {
	var tmp word
	tmp, err = v.LoadWord(address(v.regs[b]), word(v.regs[c]))
	v.regs[a] = register(tmp)
	return
}

func loadb(v *vm, a, b, c byte) (err error) {
	var tmp byte
	tmp, err = v.LoadByte(address(v.regs[b]), word(v.regs[c]))
	v.regs[a] = register(tmp)
	return
}

func loadi(v *vm, a, b, c byte) error {
	v.regs[a] = register(b)<<8 | register(c)
	return nil
}

func store(v *vm, a, b, c byte) error {
	return v.StoreWord(address(v.regs[b]), word(v.regs[c]), word(v.regs[a]))
}

func storeb(v *vm, a, b, c byte) error {
	return v.StoreByte(address(v.regs[b]), word(v.regs[c]), byte(v.regs[a]))
}

func printi(v *vm, a, b, c byte) error {
	fmt.Printf("%d", v.regs[a])
	return nil
}

func printc(v *vm, a, b, c byte) error {
	fmt.Printf("%c", v.regs[a])
	return nil
}

func printii(v *vm, a, b, c byte) error {
	fmt.Printf("%d", word(b)<<8|word(c))
	return nil
}

func printci(v *vm, a, b, c byte) error {
	fmt.Printf("%c", word(b)<<8|word(c))
	return nil
}
