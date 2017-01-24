package calc

type Calc struct {
	A int
}

/***********
type interface Operator {
	(c *Calc) Plus(b int) int;
	(c *Calc) Times(b int) int
}
************/

func (c *Calc) Plus(b int) *Calc {
	return &Calc{c.A + b}
}

func (c *Calc) Times(b int) *Calc {
	return &Calc{c.A * b}
}
