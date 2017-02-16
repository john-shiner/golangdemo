package main

import "fmt"
import _ "math"

type Calc struct {
	acc float64
}

const (
	OP_ADD = 1 << iota
	OP_SUB
	OP_MUL
	OP_DIV
)

func main() {
	fmt.Println("Go Calculation Engine Main")
	var c Calc

	c.Do(OP_ADD, 9)
	c.Do(OP_MUL, 3)

	c.Do(OP_ADD, 4)
	c.Do(OP_MUL, 5)
}

func (c *Calc) Do(op int, v float64) float64 {
	switch op {
	case OP_ADD:
		c.acc += v
	case OP_MUL:
		c.acc *= v
	}
	fmt.Println(c.acc)
	return c.acc
}
