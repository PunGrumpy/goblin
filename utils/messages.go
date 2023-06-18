package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintError(message string) {
	fmt.Println(color.RedString("[!]"), message)
}

func PrintSuccess(message string) {
	fmt.Println(color.GreenString("[+]"), message)
}

func PrintInfo(message string) {
	fmt.Println(color.BlueString("[*]"), message)
}
