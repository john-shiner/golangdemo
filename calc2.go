package main

import "fmt"
import _ "math"

type Calculator struct {
	acc float64
}

type opfunc func(float64, float64) float64

func (c *Calculator) Do(op opfunc, v float64) float64 {
	c.acc = op(c.acc, v)
	return c.acc
}

type operation struct {
	f opfunc
	v float64
}

func Add(a, b float64) float64 { return a + b }
func Sub(a, b float64) float64 { return a - b }
func Mul(a, b float64) float64 { return a * b }
func Div(a, b float64) float64 { return a / b }
func Max(a, b float64) float64 {
	if a > b {
		return b
	}
	return a
}

// Consdider storing separate calculators and exchange model semantics
// Max function is generic, but there are rules that should disable subsequent
// calculations once some threshold has been reached.  This needs to be part of the
// operation semantics (hence, instead of "Max()", it might be a "Limit()"
// function).

func main() {
	fmt.Println("Go Calculation Engine Main")
	var c Calculator
	var ops = []operation{
		{Add, 9},
		{Mul, 4},
		{Sub, 2},
		{Div, 2},
		{Mul, 3},
	}
	maxEarnings2 := 48.0
	x := 0.0
	for i, v := range ops {
		x = c.Do(v.f, v.v)
		fmt.Printf("%d\t\t%6.1f\t\t%v\n", i, x, x)
	}
	r := x
	fmt.Printf("original earnings\t\t%6.1f\n", r)
	r1 := Add(c.acc, 5.0)
	r2 := c.Do(Max, maxEarnings2)
	r3 := c.Do(Add, 5.0)
	r4 := Max(c.acc, 52.0)
	fmt.Printf("newearnings\t\t%6.1f, \t\t%6.1f, \t\t%6.1f, \t\t%6.1f, \t\t%6.1f, %6.1f\n", c.acc, r, r1, r2, r3, r4)
}
