package paginator

import "math"

type PaginatorType string

const (
	OffsetPaginatorType PaginatorType = "offset"
	CursorPaginatorType PaginatorType = "cursor"
)

func ParsePaginatorType(paginator string) PaginatorType {
	if paginator == string(OffsetPaginatorType) || paginator == string(CursorPaginatorType) {
		return PaginatorType(paginator)
	}
	return OffsetPaginatorType
}

type Paginator[T any] interface {
	GetPage() int
	GetLimit() int
	GetOffset() int
	GetTotalPages(count int64) int64

	Paginate(paginate func(limit, offset int) (data T, count int64, err error)) (*PaginatorResult[T], error)
}

func NewPaginator[T any](paginator PaginatorType, page, size int) Paginator[T] {
	if paginator == OffsetPaginatorType {
		return NewOffsetPaginator[T](page, size)
	}
	// TODO(sora): replace with cursor paginator
	return NewOffsetPaginator[T](page, size)
}

type PaginatorResult[T any] struct {
	TotalPages int64
	Count      int64
	Data       T
}

type OffsetPaginator[T any] struct {
	limit int
	page  int
}

func NewOffsetPaginator[T any](page, limit int) Paginator[T] {
	return &OffsetPaginator[T]{
		page:  page,
		limit: limit,
	}
}

func (p *OffsetPaginator[T]) GetPage() int {
	if p.page <= 0 {
		p.page = 1
	}
	return p.page
}

func (p *OffsetPaginator[T]) GetLimit() int {
	if p.limit > 10 || p.limit <= 0 {
		p.limit = 10
	}
	return p.limit
}

func (p *OffsetPaginator[T]) GetOffset() int {
	return p.GetLimit() * (p.GetPage() - 1)
}

func (p *OffsetPaginator[T]) GetTotalPages(count int64) int64 {
	return int64(math.Ceil(float64(count) / float64(p.GetLimit())))
}

func (p *OffsetPaginator[T]) Paginate(paginate func(limit, offset int) (data T, count int64, err error)) (*PaginatorResult[T], error) {
	data, count, err := paginate(p.GetLimit(), p.GetOffset())
	if err != nil {
		return nil, err
	}

	return &PaginatorResult[T]{
		Count:      count,
		TotalPages: p.GetTotalPages(count),
		Data:       data,
	}, nil
}
