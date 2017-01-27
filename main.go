package main

import "fmt"
import "./val"

func main() {

	var vals = make([]val.Val, 10)

	vals[0] = val.NewVal(11)
	vals[1] = val.NewVal(9)

	x := vals[0]
	y := vals[1]
	z := val.NewVal(33)

	a, b := 9, 5

	fmt.Println(x.Plus(0))          // 11 = 11 + 0
	fmt.Println(y.Times(1))         //  9 = 9 * 1
	fmt.Println(vals[0].Times(a))   // 99 = 11 * 9
	fmt.Println(vals[0].Plus(b))    // 16 = 11 + 5
	fmt.Println(vals[1].Times(a))   // 81 = 9 * 9
	fmt.Println(vals[0].Times(4))   // 44 = 11 * 4
	fmt.Println(z.Times(3).Plus(4)) // 103 = (33 * 3) + 4

	zz := val.NewVal(7)
	HypotSquared := val.NewVal(a*a + b*b)
	fmt.Println(HypotSquared)                                   // 106 = 9*9 + 5*5
	fmt.Println(HypotSquared.Plus(zz.Int()).Plus(-4).Times(10)) // 1090 = ((106 + 7 - 4)*10)
}

/***** Output *****
[ `go run main.go` | done: 206.294017ms ]
	11
	9
	99
	16
	81
	44
	103
	106
	1090
*******************/
