package util

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"time"

	"github.com/rotem-ester/office-reservation-app/service/pkg/store"
)

func GetDaysNumByYearAndMonth(year int, month time.Month) int {
	// the first day of the month
	t := time.Date(year, month, 1, 23, 0, 0, 0, time.UTC)

	// adding one month and subtracting one day will give the last day of the given month
	t = t.AddDate(0, 1, -1)
	return t.Day()
}

func ParseQueryParams(rawUrl *url.URL) (url.Values, error) {
	u, err := url.Parse(rawUrl.RequestURI())
	if err != nil {
		return nil, fmt.Errorf("failed to parse request URL: %w", err)
	}

	params, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to parse query params: %w", err)
	}

	return params, nil
}

func EnsureParams(params url.Values) error {
	for _, param := range store.Get().RequestParams {
		if params[param] == nil {
			return fmt.Errorf("missing param '%s' in request for revenue", param)
		}
	}

	return nil
}

func ParseParams(params url.Values) (int, time.Month, error) {
	year, err := strconv.Atoi(params["year"][0])
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing year param: %w", err)
	}
	if ok, _ := regexp.MatchString("^([0-9]{4})$", params["year"][0]); !ok {
		return 0, 0, fmt.Errorf("invalid year param. please provide a year with format YYYY")
	}

	monthNum, err := strconv.Atoi(params["month"][0])
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing month param: %w", err)
	}
	if monthNum < 1 || monthNum > 12 {
		return 0, 0, fmt.Errorf("invalid month param. please provide a month as number between 1-12")
	}
	
	month := time.Month(monthNum)

	return year, month, nil
}
