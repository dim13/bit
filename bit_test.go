package bit

import "testing"

func TestBit(t *testing.T) {
	bf := new(Field)
	for k := 0; k < 64; k++ {
		if bf.IsSet(k) || !bf.IsClear(k) {
			t.Error("got set, want clear", k)
		}
		bf.Clear(k)
		if bf.IsSet(k) || !bf.IsClear(k) {
			t.Error("got set, want clear", k)
		}
		bf.Set(k)
		if bf.IsClear(k) || !bf.IsSet(k) {
			t.Error("got clear, want set", k)
		}
	}
}

func TestShrink(t *testing.T) {
	bf := new(Field)
	bf.Set(0)
	bf.Set(32)
	bf.Set(64)
	t.Log(bf)
	bf.Clear(64)
	bf.Shrink()
	bf.Set(128)
	t.Log(bf)
}
