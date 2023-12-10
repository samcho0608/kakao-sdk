package keyword

import (
	"kakao-sdk/local/common"
	"kakao-sdk/pkg/geo"
)

type Option func(*Options)

type Options struct {
	rect *geo.Rectangle
	page int
	size int
	sort common.Sort
}

func WithRect(rect *geo.Rectangle) Option {
	return func(o *Options) {
		o.rect = rect
	}
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

func WithSort(sort common.Sort) Option {
	return func(o *Options) {
		o.sort = sort
	}
}
