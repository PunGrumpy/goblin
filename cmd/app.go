package cmd

import (
	"github.com/PunGrumpy/goblin/internal/app"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           app.Name,
	Short:         app.Description,
	Version:       app.Version,
	SilenceErrors: true,
	BashCompletionFunction: `
	__goblin_bash_completion() {
		local current_word
		COMPREPLY=()
		current_word="${COMP_WORDS[COMP_CWORD]}"
		if [[ ${COMP_CWORD} -eq 1 ]]; then
			COMPREPLY=( $(compgen -W "hash reverse" -- ${current_word}) )
		fi

		return 0

	} && complete -F __goblin_bash_completion goblin
	`,
}

func Execute() error {
	return rootCmd.Execute()
}
