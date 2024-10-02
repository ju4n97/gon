package pagination

import "fmt"

type OffsetPagination[T any] struct {
	Data  []T                    `json:"data,omitempty"`
	Meta  *OffsetPaginationMeta  `json:"meta,omitempty"`
	Links *OffsetPaginationLinks `json:"links,omitempty"`
}

type OffsetPaginationMeta struct {
	Total   int `json:"total"`
	Count   int `json:"count"`
	Pages   int `json:"pages"`
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

type OffsetPaginationLinks struct {
	First    string `json:"first,omitempty"`
	Last     string `json:"last,omitempty"`
	Previous string `json:"previous,omitempty"`
	Next     string `json:"next,omitempty"`
}

const DefaultPageSize = 10

func NewOffsetPagination[T interface{}](data []T, total int, page int) *OffsetPagination[T] {
	pages := calculatePages(total, DefaultPageSize)
	count := len(data)

	meta := &OffsetPaginationMeta{
		Total:   total,
		Count:   count,
		Pages:   pages,
		Page:    page,
		PerPage: DefaultPageSize,
	}

	links := &OffsetPaginationLinks{
		First:    "/?page=1",
		Last:     fmt.Sprintf("/?page=%d", pages),
		Previous: fmt.Sprintf("/?page=%d", max(1, meta.Page-1)),
		Next:     fmt.Sprintf("/?page=%d", min(pages, meta.Page+1)),
	}

	return &OffsetPagination[T]{
		Data:  data,
		Meta:  meta,
		Links: links,
	}
}

func calculatePages(total, pageSize int) int {
	if total == 0 {
		return 0
	}

	return (total + pageSize - 1) / pageSize
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
