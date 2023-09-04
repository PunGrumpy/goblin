package main

import (
	"github.com/PunGrumpy/goblin/cmd"
	"github.com/PunGrumpy/goblin/external/logger"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logger.PrintError(err.Error())
	}
}
