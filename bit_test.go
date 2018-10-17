package bit

import "testing"

func TestBit(t *testing.T) {
	bf := new(Field)
	for i := 0; i < 64; i++ {
		if bf.IsSet(i) || !bf.IsClear(i) {
			t.Error("got set, want clear")
		}
		bf.Clear(i)
		if bf.IsSet(i) || !bf.IsClear(i) {
			t.Error("got set, want clear")
		}
		bf.Set(i)
		if bf.IsClear(i) || !bf.IsSet(i) {
			t.Error("got clear, want set")
		}
	}
}
