package cmd

import (
	"os"
	"strings"

	"github.com/PunGrumpy/goblin/external/logger"
	"github.com/spf13/cobra"
)

func detectShell() string {
	shell := os.Getenv("SHELL")
	if strings.Contains(shell, "bash") {
		return "bash"
	} else if strings.Contains(shell, "zsh") {
		return "zsh"
	} else if strings.Contains(shell, "fish") {
		return "fish"
	}
	return ""
}

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate completion script and auto detect shell",
	Run: func(cmd *cobra.Command, args []string) {
		shell := detectShell()
		home := os.Getenv("HOME")

		switch shell {
		case "bash":
			rootCmd.GenBashCompletionFile(home + "/.bash_completion")
		case "zsh":
			rootCmd.GenZshCompletionFile(home + "/.zsh_completion")
		case "fish":
			rootCmd.GenFishCompletionFile(home+"/.config/fish/completions/goblin.fish", true)
		default:
			logger.PrintError("Unsupported shell")
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
