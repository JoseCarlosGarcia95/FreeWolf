package symbols

import (
	"strings"

	"github.com/JoseCarlosGarcia95/FreeWolf/sophie/core/math/symbolic"
	"github.com/JoseCarlosGarcia95/FreeWolf/sophie/core/math/symbolic/atoms"
	"github.com/JoseCarlosGarcia95/FreeWolf/sophie/core/math/symbolic/magnitudes"
)

// Product represent symbolic product
type Product struct {
	childs  []symbolic.Symbol
	factors []symbolic.Magnitude
	terms   symbolic.Terms
}

// NewProduct generate a new product
func NewProduct() symbolic.Symbol {
	product := Product{}
	product.childs = make([]symbolic.Symbol, 0)
	product.factors = make([]symbolic.Magnitude, 0)
	product.terms = magnitudes.NewSymbolicNumber()
	return product
}

// NewProductFromNumber create a new product object from number.
func NewProductFromNumber(num atoms.Number) symbolic.Symbol {
	product := NewProduct().(Product)
	product.terms = magnitudes.NewSymbolicNumberFromNumber(num)
	return product
}

// Childs return childs of product.
func (product Product) Childs() []symbolic.Symbol {
	return product.childs
}

// Add sum a magnitude.
func (product Product) Add(magn symbolic.Magnitude) symbolic.Symbol {
	newSum := NewSum().(Sum)
	newSum.terms = magn

	return product.AppendNumber(newSum, atoms.NewComplexFromInteger(1))
}

// Multiply product a magnitude.
func (product Product) Multiply(magn symbolic.Magnitude) symbolic.Symbol {
	newProduct := NewProduct().(Product)
	newProduct.terms = magn

	return product.AppendNumber(newProduct, atoms.NewComplexFromInteger(1))
}

// AppendNumber append a new symbol to child.
func (product Product) AppendNumber(sym symbolic.Symbol, num atoms.Number) symbolic.Symbol {
	return product.Append(sym, magnitudes.NewSymbolicNumberFromNumber(num))
}

// Append append a new symbol to child.
func (product Product) Append(sym symbolic.Symbol, magn symbolic.Magnitude) symbolic.Symbol {
	product.childs = append(product.childs, sym)
	product.factors = append(product.factors, magn)
	return product
}

// Factors return factors of product.
func (product Product) Factors() []symbolic.Magnitude {
	return product.factors
}

// Type return this type.
func (product Product) Type() symbolic.SymbolType {
	if len(product.childs) > 0 {
		return symbolic.SymbolTypeProduct
	}

	return product.terms.Type()
}

// Term return term
func (product Product) Term() symbolic.Terms {
	return product.terms
}

// Evaluate eval expression.
func (product Product) Evaluate() symbolic.Symbol {
	return product
}

// String return a string representation
func (product Product) String() string {
	var sb strings.Builder
	childs := len(product.childs)

	if childs == 0 {
		sb.WriteString(product.terms.String())
	} else {
		sb.WriteString("(")
		sb.WriteString(product.terms.String())

		for i := 0; i < childs; i++ {
			child := product.childs[i]
			factor := product.factors[i]
			printFactor := true

			switch factor.(type) {
			case magnitudes.Number:
				if factor.Val().(atoms.Number).IsReal() {
					cmp, _ := factor.Val().(atoms.Number).Compare(atoms.ComplexOne)
					printFactor = cmp != 0
				}
			}

			sb.WriteString("*")
			sb.WriteString(child.String())

			if printFactor {
				sb.WriteString("^(")
				sb.WriteString(factor.String())
				sb.WriteString(")")
			}

		}

		sb.WriteString(")")
	}

	return sb.String()
}
