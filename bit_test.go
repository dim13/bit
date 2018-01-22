package bit

import "testing"

func TestBit(t *testing.T) {
	bf := new(Field)
	for i := 0; i < 64; i++ {
		if bf.IsSet(i) || !bf.IsClr(i) {
			t.Error("got set, want clear")
		}
		bf.Set(i)
		if bf.IsClr(i) || !bf.IsSet(i) {
			t.Error("got clear, want set")
		}
		bf.Clr(i)
		if bf.IsSet(i) || !bf.IsClr(i) {
			t.Error("got set, want clear")
		}
	}
}
