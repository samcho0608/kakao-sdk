package category

import (
	"fmt"
	"kakao-sdk/core"
	"kakao-sdk/local/common"
	"kakao-sdk/pkg/geo"
	"os"
)

func Example() {
	kakaoRESTAPIKey := os.Getenv("KAKAO_REST_API_KEY")

	if err := core.InitializeSDK(kakaoRESTAPIKey); err != nil {
		fmt.Println("err:", err)
		return
	}

	cli, err := core.Default(
		core.WithAPICollectionPath("/v2/local"),
	)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	resp, err := SearchPlacesByCategory(
		cli,
		common.CategoryGroupCodes.Cafe,
		&geo.Coordinate{
			X: 127.110306812433,
			Y: 37.394245407468,
		},
		10000,
	)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("response received:", resp.Meta.IsEnd)
	// Output: response received: false
}
