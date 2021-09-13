// Package bit implements variable-size bit-fields
package bit

// Field is a variable-length bit-field
type Field []byte

// Set bit
func (f *Field) Set(k int) {
	n, m := k>>3, byte(1<<uint(k&7))
	if len(*f) <= n {
		*f = append(*f, make([]byte, n-len(*f)+1)...)
	}
	(*f)[n] |= m
}

// IsSet checks if bit is set
func (f *Field) IsSet(k int) bool {
	n, m := k>>3, byte(1<<uint(k&7))
	return len(*f) > n && (*f)[n]&m != 0
}

// Clear bit
func (f *Field) Clear(k int) {
	n, m := k>>3, byte(1<<uint(k&7))
	if len(*f) > n {
		(*f)[n] &^= m
	}
}

// IsClear checks if bit is cleared
func (f *Field) IsClear(k int) bool {
	n, m := k>>3, byte(1<<uint(k&7))
	return len(*f) <= n || (*f)[n]&m == 0
}
