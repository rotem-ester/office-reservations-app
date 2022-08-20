package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewMonthlyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "monthly 2014 02",
		Short: "returns information about expected revenue and capacity for a specific month",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("please provide the requiered args")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			err := RunMonthlyCommand(args)
			if err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}

func RunMonthlyCommand(args []string) error {
	revenue, err := RunRevenueCommand(args)
	if err != nil {
		return err
	}

	capacity, err := RunCapacityCommand(args)
	if err != nil {
		return err
	}

	fmt.Printf("%s-%s: expected revenue: $%s, expected total capacity of the unreserved offices: %s\n", args[0], args[1], revenue, capacity)
	return nil
}
