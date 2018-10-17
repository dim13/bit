// Package bit implements variable-size bit-fields
package bit

// Field is a variable bit-field
type Field []byte

// IsSet checks if bit n is set
func (bf *Field) IsSet(n int) bool {
	bit, off := n>>3, 1<<uint(n&7)
	if len(*bf) <= bit {
		return false
	}
	return (*bf)[bit]&byte(off) != 0
}

// Set bit n
func (bf *Field) Set(n int) *Field {
	bit, off := n>>3, 1<<uint(n&7)
	if len(*bf) <= bit {
		*bf = append(*bf, make([]byte, bit-len(*bf)+1)...)
	}
	(*bf)[bit] |= byte(off)
	return bf
}

// IsClr checks if bit n is set
func (bf *Field) IsClr(n int) bool {
	bit, off := n>>3, 1<<uint(n&7)
	if len(*bf) <= bit {
		return true
	}
	return (*bf)[bit]&byte(off) == 0
}

// Clr bit n
func (bf *Field) Clr(n int) *Field {
	bit, off := n>>3, 1<<uint(n&7)
	if len(*bf) <= bit {
		return bf
	}
	(*bf)[bit] &^= byte(off)
	return bf
}
