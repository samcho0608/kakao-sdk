package coordinatesystems

import (
	"encoding/json"
	"kakao-sdk/core"
	"kakao-sdk/local/common"
	"kakao-sdk/pkg/geo"
	"strconv"
)

const (
	apiPath = "/geo/transcoord.json"
)

func ConvertCoordinateSystem(
	client *core.KakaoSDKClient,
	coordinate *geo.Coordinate,
	inputCoordinateSystem geo.CoordinateSystem,
	outputCoordinateSystem geo.CoordinateSystem,
) (*Response, error) {

	queryParam := map[string]string{
		common.QueryParamKeys.X:                      strconv.FormatFloat(coordinate.X, 'f', -1, 64),
		common.QueryParamKeys.Y:                      strconv.FormatFloat(coordinate.Y, 'f', -1, 64),
		common.QueryParamKeys.InputCoordinateSystem:  string(inputCoordinateSystem),
		common.QueryParamKeys.OutputCoordinateSystem: string(outputCoordinateSystem),
	}

	respJSON, err := client.NewRequest().
		SetQueryParams(queryParam).
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
