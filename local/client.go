package local

import (
	"kakao-sdk/core"
	"kakao-sdk/local/address"
	"kakao-sdk/local/category"
	"kakao-sdk/local/common"
	"kakao-sdk/local/coordinatesystems"
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
	return keyword.SearchPlacesByKeyword(c.KakaoSDKClient, query, categoryGroupCode, origin, radiusInMeters, options...)
}

func (c *Client) SearchPlacesByCategory(
	categoryGroupCode common.CategoryGroupCode,
	origin *geo.Coordinate,
	radiusInMeters int,
	options ...category.Option,
) (*category.Response, error) {
	return category.SearchPlacesByCategory(c.KakaoSDKClient, categoryGroupCode, origin, radiusInMeters, options...)
}

func (c *Client) ConvertCoordinateSystem(
	coordinate *geo.Coordinate,
	inputCoordinateSystem geo.CoordinateSystem,
	outputCoordinateSystem geo.CoordinateSystem,
) (*coordinatesystems.Response, error) {
	return coordinatesystems.ConvertCoordinateSystem(c.KakaoSDKClient, coordinate, inputCoordinateSystem, outputCoordinateSystem)
}

func (c *Client) SearchPlacesByAddress(
	query string,
	options ...address.Option,
) (*address.Response, error) {
	return address.SearchPlacesByAddress(c.KakaoSDKClient, query, options...)
}
