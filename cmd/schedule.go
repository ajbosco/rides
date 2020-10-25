package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func newScheduleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schedule",
		Short: "Show your upcoming workout schedule and add/remove to it",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}

	cmd.AddCommand(
		newShowCmd(),
		newAddCmd(),
		newRemoveCmd(),
	)

	return cmd
}
