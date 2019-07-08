package symbolic

// Function evaluate functions
type Function interface {
	Terms
	f(val interface{}) interface{}
}
