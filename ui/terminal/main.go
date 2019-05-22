package terminal

import (
	"fmt"

	"github.com/JoseCarlosGarcia95/FreeWolf/core"
)

// InitializeUI generate the first iteration for UI.
func InitializeUI(kernel *core.Kernel) {
	NextStep(kernel)
}

// NextStep is called when one action is taken.
func NextStep(kernel *core.Kernel) {
	var input string

	fmt.Printf("[%d]:= ", kernel.CellsIn)
	fmt.Scanln(&input)

	kernel.CellsIn++

	NextStep(kernel)
}
