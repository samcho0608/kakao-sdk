package coordinatesystems

type Response struct {
	Meta      Meta       `json:"meta"`
	Documents []Document `json:"documents"`
}

type Meta struct {
	TotalCount int `json:"total_count"`
}

type Document struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
