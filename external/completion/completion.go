package completion

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/PunGrumpy/goblin/external/logger"
	"github.com/PunGrumpy/goblin/internal/app"
	"github.com/spf13/cobra"
)

func GenerateCompletion(cmd *cobra.Command) {
	zshCompletionPath := "~/.zsh/completions/"
	bashCompletionPath := "/etc/bash_completion.d/"
	fishCompletionPath := "~/.config/fish/completions/"

	switch GetShellType() {
	case "zsh":
		if err := os.MkdirAll(zshCompletionPath, 0755); err != nil {
			logger.PrintError(err.Error())
		}
		if err := cmd.GenZshCompletionFile(filepath.Join(zshCompletionPath, app.Name+".zsh")); err != nil {
			logger.PrintError(err.Error())
		}
	case "bash":
		if err := os.MkdirAll(bashCompletionPath, 0755); err != nil {
			logger.PrintError(err.Error())
		}
		if err := cmd.GenBashCompletionFile(filepath.Join(bashCompletionPath, app.Name)); err != nil {
			logger.PrintError(err.Error())
		}
	case "fish":
		if err := os.MkdirAll(fishCompletionPath, 0755); err != nil {
			logger.PrintError(err.Error())
		}
		if err := cmd.GenFishCompletionFile(filepath.Join(fishCompletionPath, app.Name+".fish"), true); err != nil {
			logger.PrintError(err.Error())
		}
	}
}

func GetShellType() string {
	shell := os.Getenv("SHELL")
	if strings.Contains(shell, "bash") {
		return "bash"
	} else if strings.Contains(shell, "zsh") {
		return "zsh"
	} else if strings.Contains(shell, "fish") {
		return "fish"
	}
	return "unknown"
}
