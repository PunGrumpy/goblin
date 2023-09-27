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
	Use:       "completions",
	Short:     "Generate completion script and auto detect shell",
	Args:      cobra.MinimumNArgs(0),
	ValidArgs: []string{"bash", "zsh", "fish"},
	Run: func(cmd *cobra.Command, args []string) {
		var shell string
		if len(args) != 0 {
			shell = args[0]
		} else {
			shell = detectShell()
		}
		home := os.Getenv("HOME")

		switch shell {
		case cmd.ValidArgs[0]:
			if err := rootCmd.GenBashCompletionFile(home + "/.bash_completion"); err != nil {
				logger.PrintError("Error generating bash completion file")
				return
			}
			logger.PrintSuccess("Bash completion file generated into ~/.bash_completion")
		case cmd.ValidArgs[1]:
			if err := rootCmd.GenZshCompletionFile(home + "/.zshrc"); err != nil {
				logger.PrintError("Error generating zsh completion file")
				return
			}
			logger.PrintSuccess("Zsh completion file generated into ~/.zshrc")
		case cmd.ValidArgs[2]:
			if err := rootCmd.GenFishCompletionFile(home+"/.config/fish/completions/goblin.fish", true); err != nil {
				logger.PrintError("Error generating fish completion file")
				return
			}
			logger.PrintSuccess("Fish completion file generated into ~/.config/fish/completions/goblin.fish")
		default:
			logger.PrintError("Unsupported shell")
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
