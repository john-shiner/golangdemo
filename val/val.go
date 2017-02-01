package val

import "strconv"

type Val struct {
	a int
}

/******
type Value interface  {
	(c *Val) Plus(b int) int;
	(c *Val) Times(b int) int;
	(c *Val) NewValue(prop string, value string);
}
********/

func NewVal(x int) Val {
	return Val{a: x}
}

func (v *Val) Plus(b int) *Val {
	return &Val{v.a + b}
}

func (v *Val) Times(b int) *Val {
	return &Val{v.a * b}
}

func (v *Val) V() int {
	return v.a
}

func (v Val) String() string {
	return strconv.Itoa(v.a) // convert int to string
}

func (v Val) Int() int {
	return v.a
}
