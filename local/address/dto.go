package address

type Response struct {
	Meta      Meta       `json:"meta"`
	Documents []Document `json:"documents"`
}

type Meta struct {
	TotalCount    int  `json:"total_count"`
	PageableCount int  `json:"pageable_count"`
	IsEnd         bool `json:"is_end"`
}

type Document struct {
	Address     Address     `json:"address"`
	RoadAddress RoadAddress `json:"road_address"`
	AddressName string      `json:"address_name"`
	AddressType string      `json:"address_type"`
	X           string      `json:"x"`
	Y           string      `json:"y"`
}

type Address struct {
	AddressName       string `json:"address_name"`
	Region1DepthName  string `json:"region_1depth_name"`
	Region2DepthName  string `json:"region_2depth_name"`
	Region3DepthName  string `json:"region_3depth_name"`
	Region3DepthHName string `json:"region_3depth_h_name"`
	HCode             string `json:"h_code"`
	BCode             string `json:"b_code"`
	MountainYN        string `json:"mountain_yn"`
	MainAddressNo     string `json:"main_address_no"`
	SubAddressNo      string `json:"sub_address_no"`
	X                 string `json:"x"`
	Y                 string `json:"y"`
}

type RoadAddress struct {
	AddressName      string `json:"address_name"`
	Region1DepthName string `json:"region_1depth_name"`
	Region2DepthName string `json:"region_2depth_name"`
	Region3DepthName string `json:"region_3depth_name"`
	RoadName         string `json:"road_name"`
	UnderGroundYn    string `json:"underground_yn"`
	MainBuildingNo   string `json:"main_building_no"`
	SubBuildingNo    string `json:"sub_building_no"`
	BuildingName     string `json:"building_name"`
	ZipCode          string `json:"zone_no"`
	X                string `json:"x"`
	Y                string `json:"y"`
}
