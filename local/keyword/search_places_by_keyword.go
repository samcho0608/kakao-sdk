package keyword

import (
	"encoding/json"
	"fmt"
	"kakao-sdk/core"
	"kakao-sdk/local/common"
	"kakao-sdk/pkg/geo"
	"strconv"
)

func SearchPlacesByKeyword(
	client *core.KakaoSDKClient,
	query string,
	categoryGroupCode common.CategoryGroupCode,
	origin *geo.Coordinate,
	radiusInMeters int,
	options ...Option,
) (*Response, error) {
	o := &Options{}
	for _, opt := range options {
		opt(o)
	}

	queryParams := makeQueryParams(query, categoryGroupCode, origin, radiusInMeters, o)

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

func makeQueryParams(
	query string,
	categoryGroupCode common.CategoryGroupCode,
	origin *geo.Coordinate,
	radiusInMeters int,
	options *Options,
) map[string]string {
	queryParams := map[string]string{
		common.QueryParamKeys.Query:             query,
		common.QueryParamKeys.CategoryGroupCode: categoryGroupCode.String(),
		common.QueryParamKeys.X:                 strconv.FormatFloat(origin.X, 'f', -1, 64),
		common.QueryParamKeys.Y:                 strconv.FormatFloat(origin.Y, 'f', -1, 64),
		common.QueryParamKeys.Radius:            strconv.Itoa(radiusInMeters),
	}

	if options != nil {
		if options.rect != nil {
			queryParams[common.QueryParamKeys.Rect] = fmt.Sprintf(
				"%v,%v,%v,%v",
				options.rect.LeftTop.X,
				options.rect.LeftTop.Y,
				options.rect.RightBottom.X,
				options.rect.RightBottom.Y,
			)
		}

		if options.page > 0 {
			queryParams[common.QueryParamKeys.Page] = strconv.Itoa(options.page)
		}

		if options.size > 0 {
			queryParams[common.QueryParamKeys.Size] = strconv.Itoa(options.size)
		}

		if options.sort != "" {
			queryParams[common.QueryParamKeys.Sort] = options.sort.String()
		}
	}

	return queryParams
}
