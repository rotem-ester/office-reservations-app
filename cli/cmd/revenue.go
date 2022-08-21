package cmd

import (
	"fmt"
	"io"

	httpUtil "github.com/rotem-ester/office-reservations-app/cli/pkg/http_util"
	"github.com/rotem-ester/office-reservations-app/cli/pkg/util"
	"github.com/spf13/cobra"
)

func NewRevenueCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revenue 2014 02",
		Short: "returns information about expected revenue for a specific month",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("please provide the requiered args")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := RunRevenueCommand(args)
			if err != nil {
				return err
			}

			fmt.Printf("%s-%s: expected revenue: $%s\n", args[0], args[1], res)
			return nil
		},
	}

	return cmd
}

func RunRevenueCommand(args []string) (string, error) {
	params, err := util.ParseArgs(args)
	if err != nil {
		return "", err
	}

	res, err := httpUtil.MakeHttpGetRequest("/revenue", params)
	if err != nil {
		return "", fmt.Errorf("request to server failed: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	strRes := string(body)

	if res.StatusCode > 200 {
		return "", fmt.Errorf("request failed with %s: %s", res.Status, strRes)
	}

	return strRes, nil
}
