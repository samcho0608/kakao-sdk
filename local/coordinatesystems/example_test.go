package coordinatesystems

import (
	"fmt"
	"kakao-sdk/core"
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

	resp, err := ConvertCoordinateSystem(
		cli,
		&geo.Coordinate{
			X: 127.110306812433,
			Y: 37.394245407468,
		},
		geo.CoordinateSystems.WCONGNAMUL,
		geo.CoordinateSystems.WGS84,
	)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("response received:", resp.Documents[0].X, resp.Documents[0].Y)
	// Output: response received: 124.84903171887818 33.47511315667069
}
