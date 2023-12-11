package address

type Option func(*Options)

type Options struct {
	page        int
	size        int
	analyzeType AnalyzeType
}

func WithPage(page int) Option {
	return func(o *Options) {
		o.page = page
	}
}

func WithSize(resultSize int) Option {
	return func(o *Options) {
		o.size = resultSize
	}
}

func WithAnalyzeType(analyzeType AnalyzeType) Option {
	return func(o *Options) {
		o.analyzeType = analyzeType
	}
}
