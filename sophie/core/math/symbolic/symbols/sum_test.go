package symbols

import (
	"fmt"
	"testing"

	"github.com/JoseCarlosGarcia95/FreeWolf/sophie/core/math/symbolic/magnitudes"
)

func TestSum(t *testing.T) {
	sym := NewSum()
	sym = sym.Add(magnitudes.NewSymbolicNumberFromInteger(30))
	sym = sym.Add(magnitudes.NewSymbolicNumberFromInteger(30))
	sym = sym.Append(sym, magnitudes.NewSymbolicNumberFromInteger(1))

	fmt.Println(sym)
	sym = sym.Evaluate()
	fmt.Println(sym)

	if sym.String() != "(0+30+30+30*(0+30+30))" {
		t.Errorf("Sum not generated correctly")
	}
}
