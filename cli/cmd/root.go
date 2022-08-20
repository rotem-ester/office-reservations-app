package cmd

import (
	"github.com/rotem-ester/office-reservations-app/cli/pkg/store"
	"github.com/spf13/cobra"
)

func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use: store.Get().BinaryName,
		Short: "Used for getting office reservation information",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(NewRevenueCommand())
	cmd.AddCommand(NewCapacityCommand())
	cmd.AddCommand(NewMonthlyCommand())

	return cmd
}