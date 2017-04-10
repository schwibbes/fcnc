package query

import (
	"encoding/json"
	. "github.com/schwibbes/fcnc/util"
	"testing"
)

func TestLiterals(t *testing.T) {
	y := Literal{true}
	AssertTrue(y.Eval(), t)
	n := Literal{false}
	AssertFalse(n.Eval(), t)
}

func TestJsonLiteral(t *testing.T) {
	str := ` {"Val": true} `
	obj := new(Literal)
	err := json.Unmarshal([]byte(str), &obj)
	AssertTrue(err == nil, t)
	AssertTrue(obj.Eval(), t)
}

// AND
func TestConjunctionTT(t *testing.T) {

	l1 := Literal{true}
	l2 := Literal{true}

	result := And{[]Condition{l1, l2}}
	AssertTrue(result.Eval(), t)
}
func TestConjunctionTF(t *testing.T) {

	l1 := Literal{true}
	l2 := Literal{false}

	result := And{[]Condition{l1, l2}}
	AssertFalse(result.Eval(), t)
}
func TestConjunctionFFF(t *testing.T) {

	l1 := Literal{false}
	l2 := Literal{false}
	l3 := Literal{false}

	result := And{[]Condition{l1, l2, l3}}
	AssertFalse(result.Eval(), t)
}

// OR
func TestDisjunctionTT(t *testing.T) {

	l1 := Literal{true}
	l2 := Literal{true}

	result := Or{[]Condition{l1, l2}}
	AssertTrue(result.Eval(), t)
}
func TestDisjunctionTF(t *testing.T) {

	l1 := Literal{true}
	l2 := Literal{false}

	result := Or{[]Condition{l1, l2}}
	AssertTrue(result.Eval(), t)
}
func TestDisjunctionFFF(t *testing.T) {

	l1 := Literal{false}
	l2 := Literal{false}
	l3 := Literal{false}

	result := Or{[]Condition{l1, l2, l3}}
	AssertFalse(result.Eval(), t)
}

// COMBINED
func TestConjectionOfDisjunction1(t *testing.T) {

	y := Literal{true}
	n := Literal{false}

	or := Or{[]Condition{y, n}}
	and := And{[]Condition{or, n}}
	// (n && (y || n))
	AssertFalse(and.Eval(), t)
}
func TestConjectionOfDisjunction2(t *testing.T) {

	y := Literal{true}
	n := Literal{false}

	or := Or{[]Condition{y, n}}
	and := And{[]Condition{y, or}}
	// (n && (y || n))
	AssertTrue(and.Eval(), t)
}

func TestNestedDisjunctions(t *testing.T) {

	y := Literal{true}
	n := Literal{false}

	result := Or{[]Condition{
		And{[]Condition{y, y, n}},
		And{[]Condition{y, Or{[]Condition{y, n, n}}}}}}

	// (n && (y || n))
	AssertTrue(result.Eval(), t)
}
