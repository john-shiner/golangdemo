package val

import "testing"

func TestVal(t *testing.T) {
	var tests = []struct {
		name string
		a    int
		b    int
		GotPlus  int
		GotTimes int
	}{
		{"first", 5, 4, 9, 20},
		{"second", 6, 9, 15, 54},
	}
	for _, tst := range tests {
		v := NewVal(tst.a)
		got := v.Plus(tst.b).Int()
		if got != tst.GotPlus {
			t.Errorf("tst(%+v) = got(%v)", tst, tst.GotPlus)
		}
		v2 := NewVal(tst.a)
		got2 := v2.Times(tst.b).Int()
		if got2 != tst.GotTimes {
			t.Errorf("tst(%+v) = got(%v)", tst, tst.GotTimes)
		}
	}
}

