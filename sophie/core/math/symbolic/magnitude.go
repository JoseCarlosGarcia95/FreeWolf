package symbolic

// Magnitude could be a number or variable.
type Magnitude interface {
	Terms
	Val() interface{}
	Set(val interface{})
}
