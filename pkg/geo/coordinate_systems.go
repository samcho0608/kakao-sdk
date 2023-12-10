package geo

type CoordinateSystem string

var CoordinateSystems = struct {
	WGS84      CoordinateSystem
	WCONGNAMUL CoordinateSystem
	CONGNAMUL  CoordinateSystem
	WTM        CoordinateSystem
	TM         CoordinateSystem
	KTM        CoordinateSystem
	UTM        CoordinateSystem
	BESSEL     CoordinateSystem
	WKTM       CoordinateSystem
	WUTM       CoordinateSystem
}{
	WGS84:      "WGS84",
	WCONGNAMUL: "WCONGNAMUL",
	CONGNAMUL:  "CONGNAMUL",
	WTM:        "WTM",
	TM:         "TM",
	KTM:        "KTM",
	UTM:        "UTM",
	BESSEL:     "BESSEL",
	WKTM:       "WKTM",
	WUTM:       "WUTM",
}
