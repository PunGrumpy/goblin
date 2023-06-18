package cmd

import (
	"fmt"

	"github.com/PunGrumpy/goblin/internal/jenkins"
	"github.com/PunGrumpy/goblin/utils"
	"github.com/spf13/cobra"
)

var hashCmd = &cobra.Command{
	Use:   "hash [string]",
	Short: "Hash the given string",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		hash := jenkins.Hash(input)

		if hash != 0 {
			utils.PrintInfo(fmt.Sprintf("Jenkins' one time hash for \"%s\":", input))
			utils.PrintSuccess(fmt.Sprintf("Hexadecimal: 0x%X", hash))
			utils.PrintSuccess(fmt.Sprintf("Decimal: %d", hash))
			utils.PrintSuccess(fmt.Sprintf("Binary: %b", hash))
		} else {
			utils.PrintError("No hash found")
		}
	},
}

func init() {
	rootCmd.AddCommand(hashCmd)
}
