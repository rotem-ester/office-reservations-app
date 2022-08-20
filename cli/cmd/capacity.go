package cmd

import (
	"fmt"
	"io"

	httpUtil "github.com/rotem-ester/office-reservations-app/cli/pkg/http_util"
	"github.com/spf13/cobra"
)

func NewCapacityCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "capacity 2014 02",
		Short: "capacity information about expected total capacity of the unreserved offices for a specific month",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("please provide the requiered args")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := RunCapacityCommand(args)
			if err != nil {
				return err
			}

			fmt.Printf("%s-%s: expected total capacity of the unreserved offices: %s\n", args[0], args[1], res)
			return nil
		},
	}

	return cmd
}

func RunCapacityCommand(args []string) (string, error) {
	// TODO add args validation
	params := []httpUtil.QueryParam{
		{
			Key: "year",
			Value: args[0],
		},
		{
			Key: "month",
			Value: args[1],
		},
	}

	res, err := httpUtil.MakeHttpGetRequest("/capacity", params)
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