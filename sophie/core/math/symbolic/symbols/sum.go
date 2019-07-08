package symbols

import (
	"strings"

	"github.com/JoseCarlosGarcia95/FreeWolf/sophie/core/math/symbolic"
	"github.com/JoseCarlosGarcia95/FreeWolf/sophie/core/math/symbolic/atoms"
	"github.com/JoseCarlosGarcia95/FreeWolf/sophie/core/math/symbolic/magnitudes"
)

// Sum represent symbolic sum
type Sum struct {
	childs  []symbolic.Symbol
	factors []atoms.Number
	terms   symbolic.Terms
}

// NewSum generate a new sum
func NewSum() symbolic.Symbol {
	sum := Sum{}
	sum.childs = make([]symbolic.Symbol, 0)
	sum.factors = make([]atoms.Number, 0)
	sum.terms = magnitudes.NewSymbolicNumber()
	return sum
}

// NewSumFromNumber create a new sum object from number.
func NewSumFromNumber(num atoms.Number) symbolic.Symbol {
	sum := NewSum().(Sum)
	sum.terms = magnitudes.NewSymbolicNumberFromNumber(num)
	return sum
}

// Childs return childs of sum.
func (sum Sum) Childs() []symbolic.Symbol {
	return sum.childs
}

// Append append a new symbol to child.
func (sum Sum) Append(sym symbolic.Symbol, num atoms.Number) symbolic.Symbol {
	sum.childs = append(sum.childs, sym)
	sum.factors = append(sum.factors, num)
	return sum
}

// Factors return factors of sum.
func (sum Sum) Factors() []atoms.Number {
	return sum.factors
}

// Type return this type.
func (sum Sum) Type() symbolic.SymbolType {
	if len(sum.childs) > 0 {
		return symbolic.SymbolTypeRoot
	}

	return sum.terms.Type()
}

// String generate a string representation.
func (sum Sum) String() string {
	var sb strings.Builder
	childs := len(sum.childs)

	if childs == 0 {
		sb.WriteString(sum.terms.String())
	} else {
		sb.WriteString("(")
		sb.WriteString(sum.terms.String())

		for i := 0; i < childs; i++ {
			child := sum.childs[i]
			factor := sum.factors[i]
			printOp := true
			printFactor := true

			if factor.IsReal() {
				cmp, _ := factor.Compare(atoms.ComplexZero)
				if cmp < 0 {
					printOp = false
				}
				cmp, _ = factor.Compare(atoms.ComplexOne)
				printFactor = cmp != 0
			}

			if printOp {
				sb.WriteString("+")
			}

			if printFactor {
				sb.WriteString(factor.String())
				sb.WriteString("*")
			}

			sb.WriteString(child.String())

		}
		sb.WriteString(")")
	}

	return sb.String()
}
