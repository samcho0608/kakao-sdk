package local

import (
	"kakao-sdk/core"
	"kakao-sdk/local/common"
	"kakao-sdk/local/keyword"
	"kakao-sdk/pkg/geo"
)

type Client struct {
	*core.KakaoSDKClient
}

func Default() (*Client, error) {
	coreClient, err := core.Default(
		core.WithAPICollectionPath("/v2/local"),
	)
	if err != nil {
		return nil, err
	}

	return &Client{KakaoSDKClient: coreClient}, nil
}

func NewClient(coreClient *core.KakaoSDKClient) (*Client, error) {
	return &Client{KakaoSDKClient: coreClient}, nil
}

func (c *Client) SearchPlacesByKeyword(
	query string,
	categoryGroupCode common.CategoryGroupCode,
	origin *geo.Coordinate,
	radiusInMeters int,
	options ...keyword.Option,
) (*keyword.Response, error) {
	return keyword.SearchPlacesByKeyword(c.KakaoSDKClient, query, categoryGroupCode, origin, radiusInMeters)
}
