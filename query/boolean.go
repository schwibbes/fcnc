package query

type Condition interface {
	Eval() bool
}

type Literal struct {
	Val bool
}

func (l Literal) Eval() bool {
	return l.Val
}

type And struct {
	vals []Condition
}

func (op And) Eval() bool {
	result := true
	for _, v := range op.vals {
		result = result && v.Eval()
	}
	return result
}

type Or struct {
	vals []Condition
}

func (op Or) Eval() bool {
	result := false
	for _, v := range op.vals {
		result = result || v.Eval()
	}
	return result
}
