package main

import "./val"
import "log"
import "fmt"

func main() {

	var vals = make([]val.Val, 10)

	vals[0] = val.NewVal(11)
	vals[1] = val.NewVal(9)

	x := vals[0]
	y := vals[1]
	z := val.NewVal(33)

	a, b := 9, 5

	fmt.Println("var vals = make([]val.Val, 10)")

	fmt.Println("vals[0] = val.NewVal(11)")
	fmt.Println("vals[1] = val.NewVal(9)")

	fmt.Println("x := vals[0]")
	fmt.Println("y := vals[1]")
	fmt.Println("z := val.NewVal(33)")

	fmt.Println("a, b := 9, 5")

	log.Println("x.Plus(0) = ", x.Plus(0))                   // 11 = 11 + 0 ; change to Log functions
	log.Println("y.Times(1) = ", y.Times(1))                 //  9 = 9 * 1
	log.Println("vals[0].Times(a) = ", vals[0].Times(a))     // 99 = 11 * 9
	log.Println("vals[0].Plus(b) = ", vals[0].Plus(b))       // 16 = 11 + 5
	fmt.Println("vals[1].Times(a) = ", vals[1].Times(a))     // 81 = 9 * 9
	fmt.Println("vals[0].Times(4) = ", vals[0].Times(4))     // 44 = 11 * 4
	fmt.Println("z.Times(3).Plus(4) = ", z.Times(3).Plus(4)) // 103 = (33 * 3) + 4 ; test chained operations
	fmt.Println("vals[2] = ", vals[2])                       // 0 - test default initialized value
	fmt.Println("val.NewVal(2) = ", val.NewVal(2))           // 2 - test constructor

	zz := val.NewVal(7)
	HypotSquared := val.NewVal(a*a + b*b)
	fmt.Println("HypotSquared = ", HypotSquared)                                                                     // 106 = 9*9 + 5*5
	fmt.Println("HypotSquared.Plus(zz.Int()).Plus(-4).Times(10) = ", HypotSquared.Plus(zz.Int()).Plus(-4).Times(10)) // 1090 = ((106 + 7 - 4)*10)
	fmt.Println("Apparently, either log.Println and fmt.Println uses goroutines (order when mixed is not deterministic")
}
