package core

// Kernel represent the current execution information.
type Kernel struct {
	CellsIn  int64
	CellsOut int64
}

// NewKernel generate a new kernel with default values
func NewKernel() Kernel {
	newkernel := Kernel{}

	newkernel.CellsIn = 0
	newkernel.CellsOut = 0

	return newkernel
}
