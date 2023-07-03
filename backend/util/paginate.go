package util

import (
	"inventory-management/backend/internal/http/response"
	"math"
)

func CreatePagination(currPage int, limit int, totalRecords int64) (pagination response.Pagination) {
	totalPages := math.Ceil(float64(totalRecords) / float64(limit))
	if currPage <= 0 {
		currPage = 1
	}

	if totalPages <= 0 {
		totalPages = 1
	}

	pagination.TotalRecords = int(totalRecords)
	pagination.CurrentPage = currPage
	pagination.TotalPages = int(totalPages)

	return pagination
}
