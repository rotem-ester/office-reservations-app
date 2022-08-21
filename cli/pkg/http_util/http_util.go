package http_util

import (
	"fmt"
	"net/http"

	"github.com/rotem-ester/office-reservations-app/cli/pkg/store"
)

type QueryParam struct {
	Key string
	Value string
}

func MakeHttpGetRequest(path string, params []QueryParam) (*http.Response, error) {
	qp := ParamsToQuery(params)	
	url := fmt.Sprintf("%s%s%s", store.Get().ServerUrl, path, qp)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func ParamsToQuery(params []QueryParam) string {
	query := "?"
	format := "%s%s=%s"
	for i, param := range params {
		if i > 0 {
			format = "%s&%s=%s"
		}
		query = fmt.Sprintf(format, query, param.Key, param.Value)
	}

	return query
}