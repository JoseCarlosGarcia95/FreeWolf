package symbolic

// SymbolType represent different symbol types.
type SymbolType uint8

const (
	// SymbolTypeRoot return root of tree or subtree
	SymbolTypeRoot SymbolType = 'R'
	// SymbolTypeVariable if variable
	SymbolTypeVariable SymbolType = 'V'
	// SymbolTypeNumber if number
	SymbolTypeNumber SymbolType = 'N'
	// SymbolTypeProduct if product
	SymbolTypeProduct SymbolType = 'P'
)

// Symbol represent basic symbol representation.
type Symbol interface {
	Childs() []Symbol
	Append(Symbol, Magnitude) Symbol
	Factors() []Magnitude
	Type() SymbolType
	Term() Terms

	Add(Magnitude) Symbol

	Multiply(Magnitude) Symbol

	Evaluate() Symbol

	String() string
}
