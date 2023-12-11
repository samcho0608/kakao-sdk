package address

const (
	apiPath = "/search/address.json"
)

type AnalyzeType string

var AnalyzeTypes = struct {
	Similar AnalyzeType
	Exact   AnalyzeType
}{
	Similar: "similar",
	Exact:   "exact",
}

func (a AnalyzeType) String() string {
	return string(a)
}
