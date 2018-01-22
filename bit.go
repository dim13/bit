package bit

type Field []byte

const width = 8

func (bf *Field) IsSet(i int) bool {
	bit, off := i/width, 1<<(uint(i)%width)
	if len(*bf) <= bit {
		return false
	}
	return (*bf)[bit]&byte(off) != 0
}

func (bf *Field) Set(i int) *Field {
	bit, off := i/width, 1<<(uint(i)%width)
	if len(*bf) <= bit {
		*bf = append(*bf, make([]byte, bit-len(*bf)+1)...)
	}
	(*bf)[bit] |= byte(off)
	return bf
}

func (bf *Field) IsClr(i int) bool {
	bit, off := i/width, 1<<(uint(i)%width)
	if len(*bf) <= bit {
		return true
	}
	return (*bf)[bit]&byte(off) == 0
}

func (bf *Field) Clr(i int) *Field {
	bit, off := i/width, 1<<(uint(i)%width)
	if len(*bf) <= bit {
		return bf
	}
	(*bf)[bit] &^= byte(off)
	return bf
}
