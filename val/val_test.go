package val

import "testing"

func TestVal(t *testing.T) {
	var tests = []struct {
		name string
		a    int
		b    int
		got  int
	}{
		{"plus", 5, 4, 9},
		{"times", 5, 4, 20},
	}
	for _, tst := range tests {
		if tst.name == "plus" {
			v := NewVal(tst.a)
			got := v.Plus(tst.b).Int()
			if got != tst.got {
				t.Errorf("tst(%+v) = got(%v)", tst, got)
			}
			println("Plus %d, %d, %d, ok", tst.a, tst.b, tst.got)
		}
		if tst.name == "times" {
			v := NewVal(tst.a)
			got := v.Times(tst.b).Int()
			if got != tst.got {
				t.Errorf("tst(%+v) = got(%v)", tst, got)
			}
			println("Times %d, %d, %d, ok", tst.a, tst.b, tst.got)
		}
	}
}
