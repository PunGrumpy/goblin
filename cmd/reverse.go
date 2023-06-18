package cmd

import (
	"fmt"

	"github.com/PunGrumpy/goblin/internal/jenkins"
	"github.com/PunGrumpy/goblin/utils"
	"github.com/spf13/cobra"
)

var reverseCmd = &cobra.Command{
	Use:   "reverse",
	Short: "Reverse the given hash",
	Run: func(cmd *cobra.Command, args []string) {
		targetHash, _ := cmd.Flags().GetInt("target")
		length, _ := cmd.Flags().GetInt("length")
		characters, _ := cmd.Flags().GetString("characters")

		characterList := utils.GetCharacterList(characters)
		preimages := jenkins.FindPreimages(uint32(targetHash), length, characterList)

		if len(preimages) == 0 {
			utils.PrintError("No pre-images found")
			return
		}

		utils.PrintInfo(fmt.Sprintf("Possible pre-images of length %d for hash %d:", length, targetHash))
		for _, preimage := range preimages {
			if preimage != "" {
				utils.PrintSuccess(preimage)
			} else {
				utils.PrintError("No pre-images found")
			}
		}
	},
}

func init() {
	reverseCmd.Flags().Int("target", 0, "The hash to reverse")
	reverseCmd.Flags().Int("length", 0, "The length of the pre-images to find")
	reverseCmd.Flags().String("characters", "", "A list of characters to try (default is alphanumeric)")

	if err := reverseCmd.MarkFlagRequired("target"); err != nil {
		utils.PrintError(err.Error())
	}

	if err := reverseCmd.MarkFlagRequired("length"); err != nil {
		utils.PrintError(err.Error())
	}

	rootCmd.AddCommand(reverseCmd)
}
