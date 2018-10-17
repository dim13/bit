// Package bit implements variable-size bit-fields
package bit

// Field is a variable-length bit-field
type Field []byte

// IsSet checks if bit n is set
func (bf *Field) IsSet(n int) bool {
	bit, off := n>>3, byte(1<<uint(n&7))
	if len(*bf) <= bit {
		return false
	}
	return (*bf)[bit]&off != 0
}

// Set bit n
func (bf *Field) Set(n int) *Field {
	bit, off := n>>3, byte(1<<uint(n&7))
	if len(*bf) <= bit {
		*bf = append(*bf, make([]byte, bit-len(*bf)+1)...)
	}
	(*bf)[bit] |= off
	return bf
}

// IsClear checks if bit n is cleared
func (bf *Field) IsClear(n int) bool {
	bit, off := n>>3, byte(1<<uint(n&7))
	if len(*bf) <= bit {
		return true
	}
	return (*bf)[bit]&off == 0
}

// Clear bit n
func (bf *Field) Clear(n int) *Field {
	bit, off := n>>3, byte(1<<uint(n&7))
	if len(*bf) <= bit {
		return bf
	}
	(*bf)[bit] &^= off
	return bf
}
