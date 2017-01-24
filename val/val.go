package val

type Val struct {
	a int
}

/***********
type interface Operator {
	(c *Calc) Plus(b int) int;
	(c *Calc) Times(b int) int
}
************/

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
