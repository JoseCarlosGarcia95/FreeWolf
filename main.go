package main

import (
	"github.com/JoseCarlosGarcia95/FreeWolf/core"
	"github.com/JoseCarlosGarcia95/FreeWolf/ui/terminal"
)

func main() {
	kernel := core.NewKernel()
	terminal.InitializeUI(&kernel)
}
