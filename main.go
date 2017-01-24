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

	fmt.Println(x.Plus(0).V())          // 11 = 11 + 0
	fmt.Println(y.Times(1).V())         //  9 = 9 * 1
	fmt.Println(vals[0].Times(a).V())   // 99 = 11 * 9
	fmt.Println(vals[0].Plus(b).V())    // 16 = 11 + 5
	fmt.Println(vals[1].Times(a).V())   // 81 = 9 * 9
	fmt.Println(vals[0].Times(4).V())   // 44 = 11 * 4
	fmt.Println(z.Times(3).Plus(4).V()) // 99 = 33 * 3

	zz := val.NewVal(7)
	HypotSquared := val.NewVal(a*a + b*b)
	fmt.Println(HypotSquared.V())                                 // 106 = 9*9 + 5*5
	fmt.Println(HypotSquared.Plus(zz.V()).Plus(-4).Times(10).V()) // 1090 = ((106 + 7 - 4)*10)
}
