package address

import (
	"fmt"
	"kakao-sdk/core"
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

	resp, err := SearchPlacesByAddress(
		cli,
		"서울특별시 강남구 역삼동 736-1",
	)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("response received:", resp.Meta.IsEnd)
	// Output: response received: true
}
