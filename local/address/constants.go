package address

const (
	apiPath = "/v2/local/search/address.json"
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
