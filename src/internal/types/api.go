package types

import "github.com/soramon0/portfolio/src/internal/database"

type possibleReturns interface {
	database.User | []database.User | []database.ListPublishedProjectsRow | any
}

type APIListResponse[T possibleReturns] struct {
	Data       T         `json:"data"`
	Count      int64     `json:"count"`
	TotalPages int64     `json:"total_pages"`
	Error      *APIError `json:"error,omitempty"`
}

type APIResponse[T possibleReturns] struct {
	Data  T         `json:"data"`
	Error *APIError `json:"error,omitempty"`
}

func NewAPIListResponse[T possibleReturns](data T, count, totalPages int64) APIListResponse[T] {
	return APIListResponse[T]{
		Data:       data,
		Count:      count,
		TotalPages: totalPages,
	}
}

func NewAPIResponse[T possibleReturns](data T) APIResponse[T] {
	return APIResponse[T]{
		Data: data,
	}
}

type APIError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

type APIFieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type APIValidationErrors struct {
	Errors []APIFieldError `json:"errors"`
}
