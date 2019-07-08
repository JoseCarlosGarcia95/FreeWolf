package symbolic

// Terms represent terms of symbols.
type Terms interface {
	Type() SymbolType
	String() string
}
