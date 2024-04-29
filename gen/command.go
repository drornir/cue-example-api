package gen

import "github.com/spf13/cobra"

func MakeRootCommand() *cobra.Command {
	rootCommand := &cobra.Command{
		Use: "gen",
	}

	subCommandGenGoValidate := &cobra.Command{
		Use: "go-validators",
		RunE: func(cmd *cobra.Command, args []string) error {
			return GenerateGoValidators()
		},
	}
	rootCommand.AddCommand(
		subCommandGenGoValidate,
	)

	return rootCommand
}
