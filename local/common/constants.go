package common

type Sort string

type CategoryGroupCode string

var Sorts = struct {
	Distance Sort
	Accuracy Sort
}{
	Distance: "distance",
	Accuracy: "accuracy",
}

var QueryParamKeys = struct {
	Query             string
	CategoryGroupCode string
	X                 string
	Y                 string
	Radius            string
	Rect              string
	Page              string
	Size              string
	Sort              string
}{
	Query:             "query",
	CategoryGroupCode: "category_group_code",
	X:                 "x",
	Y:                 "y",
	Radius:            "radius",
	Rect:              "rect",
	Page:              "page",
	Size:              "size",
	Sort:              "sort",
}

var CategoryGroupCodes = struct {
	DepartmentStore   CategoryGroupCode
	Convenience       CategoryGroupCode
	Kindergarten      CategoryGroupCode
	School            CategoryGroupCode
	Education         CategoryGroupCode
	Parking           CategoryGroupCode
	GasStation        CategoryGroupCode
	SubwayStation     CategoryGroupCode
	Bank              CategoryGroupCode
	CultureFacility   CategoryGroupCode
	RealEstateAgent   CategoryGroupCode
	PublicInstitution CategoryGroupCode
	TouristAttraction CategoryGroupCode
	Accommodation     CategoryGroupCode
	Restaurant        CategoryGroupCode
	Cafe              CategoryGroupCode
}{
	DepartmentStore:   "MT1",
	Convenience:       "CS2",
	Kindergarten:      "PS3",
	School:            "SC4",
	Education:         "AC5",
	Parking:           "PK6",
	GasStation:        "OL7",
	SubwayStation:     "SW8",
	Bank:              "BK9",
	CultureFacility:   "CT1",
	RealEstateAgent:   "AG2",
	PublicInstitution: "PO3",
	TouristAttraction: "AT4",
	Accommodation:     "AD5",
	Restaurant:        "FD6",
	Cafe:              "CE7",
}

func (c CategoryGroupCode) String() string {
	return string(c)
}

func (s Sort) String() string {
	return string(s)
}
