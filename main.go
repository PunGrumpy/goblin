package main

import (
	"github.com/PunGrumpy/goblin/cmd"
	"github.com/PunGrumpy/goblin/utils"
)

func main() {
	if err := cmd.Execute(); err != nil {
		utils.PrintError(err.Error())
	}
}
