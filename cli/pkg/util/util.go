package util

import (
	"fmt"
	"regexp"

	httpUtil "github.com/rotem-ester/office-reservations-app/cli/pkg/http_util"
)

func ParseArgs(args []string) ([]httpUtil.QueryParam, error) {
	if ok, _ := regexp.MatchString("[0-9]{4}", args[0]); !ok {
		return nil, fmt.Errorf("invalid year param. please provide a year with format YYYY")
	}

	if ok, _ := regexp.MatchString("^(0?[1-9]|1[012])$", args[1]); !ok {
		return nil, fmt.Errorf("invalid month param. please provide a month as number between 1-12")
	}

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

	return params, nil
}