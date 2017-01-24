package main

import "fmt"
import "./calc"

func main() {

	var calcs = make([]calc.Calc, 10)

	calcs[0].A = 11
	calcs[1].A = 9

	x := calcs[0]
	y := calcs[1]
	z := calc.Calc{33}

	a, b := 9, 5

	fmt.Println(x.Plus(0).A)          // 11 = 11 + 0
	fmt.Println(y.Times(1).A)         //  9 = 9 * 1
	fmt.Println(calcs[0].Times(a).A)  // 99 = 11 * 9
	fmt.Println(calcs[0].Plus(b).A)   // 16 = 11 + 5
	fmt.Println(calcs[1].Times(a).A)  // 81 = 9 * 9
	fmt.Println(calcs[0].Times(4).A)  // 44 = 11 * 4
	fmt.Println(z.Times(3).Plus(4).A) // 99 = 33 * 3

	zz := calc.Calc{7}
	HypotSquared := calc.Calc{a*a + b*b}
	fmt.Println(HypotSquared.A)                               // 106 = 9*9 + 5*5
	fmt.Println(HypotSquared.Plus(zz.A).Plus(-4).Times(10).A) // 1090 = ((106 + 7 - 4)*10)
}
