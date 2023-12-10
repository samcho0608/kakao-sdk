package geo

type Coordinate struct {
	// X is longitude
	X float64
	// Y is latitude
	Y float64
}

type Rectangle struct {
	// LeftTop is the left top coordinate of the rectangle
	LeftTop *Coordinate
	// RightBottom is the right bottom coordinate of the rectangle
	RightBottom *Coordinate
}
