package cmd

import (
	"github.com/PunGrumpy/goblin/internal/app"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     app.Name,
	Short:   app.Description,
	Version: app.Version,
}

func init() {
	rootCmd.AddCommand(hashCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
