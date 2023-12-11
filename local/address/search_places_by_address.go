package address

import (
	"encoding/json"
	"kakao-sdk/core"
	"kakao-sdk/local/common"
	"strconv"
)

const (
	queryParamKeyAnalyzeType = "analyze_type"
)

func SearchPlacesByAddress(client *core.KakaoSDKClient, query string, options ...Option) (*Response, error) {
	o := &Options{}
	for _, opt := range options {
		opt(o)
	}

	queryParams := makeQueryParams(query, o)

	respJSON, err := client.NewRequest().
		SetQueryParams(queryParams).
		Get(apiPath)
	if err != nil {
		return nil, err
	}

	var resp *Response
	if err := json.Unmarshal(respJSON.Body(), &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func makeQueryParams(query string, options *Options) map[string]string {
	queryParams := map[string]string{
		common.QueryParamKeys.Query: query,
	}

	if options.page != 0 {
		queryParams[common.QueryParamKeys.Page] = strconv.Itoa(options.page)
	}

	if options.size != 0 {
		queryParams[common.QueryParamKeys.Size] = strconv.Itoa(options.size)
	}

	if options.analyzeType != "" {
		queryParams[queryParamKeyAnalyzeType] = options.analyzeType.String()
	}

	return queryParams
}
