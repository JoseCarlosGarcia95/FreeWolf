package symbolic

// TermsType represent different types of Terms
type TermsType uint8

const (
	// TermsTypeNumber represent number types
	TermsTypeNumber TermsType = 'N'
	// TermsTypeVariable represent variable types
	TermsTypeVariable TermsType = 'V'
)

// Terms represent terms of symbols.
type Terms interface {
	Type() TermsType
	String() string
}
