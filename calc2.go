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

func main() {
	fmt.Println("Go Calculation Engine Main")
	var c Calculator
	var ops = []operation{
		{Add, 9},
		{Mul, 4},
		{Sub, 2},
		{Div, 2},
	}

	for i, v := range ops {
		x := c.Do(v.f, v.v)
		fmt.Printf("%d\t\t%6.1f\t\t%v\n", i, x, x)
	}
}
