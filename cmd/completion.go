package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [zsh]",
	Short: "Generate shell completion script",
	Long: `Generate shell completion script for cloudamqp CLI.

To load completions:

Zsh:

  # Add to ~/.zshrc:
  source <(cloudamqp completion zsh)

  # To load completions for each session, add the script to your zsh completion directory:
  cloudamqp completion zsh > "${fpath[1]}/_cloudamqp"

  # You may need to restart your shell for completions to take effect.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"zsh"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		switch args[0] {
		case "zsh":
			return cmd.Root().GenZshCompletion(os.Stdout)
		}
		return nil
	},
}
