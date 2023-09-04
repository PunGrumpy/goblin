package cmd

import (
	"github.com/PunGrumpy/goblin/external/logger"
	"github.com/PunGrumpy/goblin/internal/app"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              app.Name,
	Short:            app.Description,
	Version:          app.Version,
	SilenceErrors:    true,
	TraverseChildren: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if err := cmd.GenBashCompletionFile(app.Name + ".bash"); err != nil {
			logger.PrintError("Failed to generate bash completion file: " + err.Error())
		}
		if err := cmd.GenZshCompletionFile(app.Name + ".zsh"); err != nil {
			logger.PrintError("Failed to generate bash completion file: " + err.Error())

		}
		if err := cmd.GenFishCompletionFile(app.Name+".fish", true); err != nil {
			logger.PrintError("Failed to generate bash completion file: " + err.Error())
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}
