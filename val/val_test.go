package val

import "testing"
import "fmt"
import "os"

func TestVal(t *testing.T) {
	// os.Setenv("LOUD", "true")
	verbose := os.Getenv("LOUD")

	var tests = []struct {
		name     string
		a        int
		b        int
		GotPlus  int
		GotTimes int
	}{
		{"first", 5, 4, 9, 20},
		{"second", 6, 9, 15, 54},
	}
	if len(verbose) > 0 {
		fmt.Println("Verbose=" + verbose)
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
			t.Errorf("tst(%+v) = got2(%v)", tst, tst.GotTimes)
		}
		if len(verbose) > 0 {

			fmt.Printf("%s, a=%d, b=%d, Plus=%d, Times=%d\n", tst.name, tst.a, tst.b, got, got2)
		}
	}
}
