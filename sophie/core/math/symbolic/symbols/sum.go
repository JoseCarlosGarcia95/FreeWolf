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
	factors []symbolic.Magnitude
	terms   symbolic.Terms
}

// NewSum generate a new sum
func NewSum() symbolic.Symbol {
	sum := Sum{}
	sum.childs = make([]symbolic.Symbol, 0)
	sum.factors = make([]symbolic.Magnitude, 0)
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

// Term return term
func (sum Sum) Term() symbolic.Terms {
	return sum.terms
}

// AppendNumber append a new symbol to child.
func (sum Sum) AppendNumber(sym symbolic.Symbol, num atoms.Number) symbolic.Symbol {
	return sum.Append(sym, magnitudes.NewSymbolicNumberFromNumber(num))
}

// Append append a new symbol to child.
func (sum Sum) Append(sym symbolic.Symbol, magn symbolic.Magnitude) symbolic.Symbol {
	sum.childs = append(sum.childs, sym)
	sum.factors = append(sum.factors, magn)
	return sum
}

// Add sum a magnitude.
func (sum Sum) Add(magn symbolic.Magnitude) symbolic.Symbol {
	newSum := NewSum().(Sum)
	newSum.terms = magn

	return sum.AppendNumber(newSum, atoms.NewComplexFromInteger(1))
}

// Multiply product a magnitude.
func (sum Sum) Multiply(magn symbolic.Magnitude) symbolic.Symbol {
	newProduct := NewProduct().(Product)
	newProduct.terms = magn

	return sum.AppendNumber(newProduct, atoms.NewComplexFromInteger(1))
}

// Factors return factors of sum.
func (sum Sum) Factors() []symbolic.Magnitude {
	return sum.factors
}

// Type return this type.
func (sum Sum) Type() symbolic.SymbolType {
	if len(sum.childs) > 0 {
		return symbolic.SymbolTypeRoot
	}

	return sum.terms.Type()
}

// Evaluate eval expression.
func (sum Sum) Evaluate() symbolic.Symbol {
	// 1. Evaluate all childs.
	// 2. If parent and childs has the same type, operate.

	childs := len(sum.childs)
	for i := 0; i < childs; i++ {
		child := sum.childs[i].Evaluate()
		factor := sum.factors[i]

		if child.Type() == symbolic.SymbolTypeNumber &&
			child.Type() == sum.Term().Type() {
			num1 := sum.terms.(magnitudes.Number)
			num2 := child.Term().(magnitudes.Number)

			sum.terms = num1.Set(num1.Val().(atoms.Number).Add(num2.Val().(atoms.Number)))

			if factor.Type() == sum.Term().Type() {
			}

			sum.childs[i] = sum.childs[len(sum.childs)-1]
			sum.childs[len(sum.childs)-1] = nil
			sum.childs = sum.childs[:len(sum.childs)-1]

			i--
			childs--
		}
	}

	return sum
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

			switch factor.(type) {
			case magnitudes.Number:
				if factor.Val().(atoms.Number).IsReal() {
					cmp, _ := factor.Val().(atoms.Number).Compare(atoms.ComplexZero)
					if cmp < 0 {
						printOp = false
					}
					cmp, _ = factor.Val().(atoms.Number).Compare(atoms.ComplexOne)
					printFactor = cmp != 0
				}
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
