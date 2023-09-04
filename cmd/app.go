package cmd

import (
	"github.com/PunGrumpy/goblin/external/completion"
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
		completion.GenerateCompletion(cmd)
	},
}

func Execute() error {
	return rootCmd.Execute()
}
