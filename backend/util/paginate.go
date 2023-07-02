package util

import (
	"inventory-management/backend/internal/http/response"
	"math"
)

func CreatePagination(currPage int, limit int, totalRecords int64) (pagination response.Pagination) {
	totalPages := math.Ceil(float64(totalRecords) / float64(limit))
	prevPage := currPage - 1
	nextPage := currPage + 1

	if currPage >= int(totalPages) {
		currPage = int(totalPages)
	}
	if prevPage <= 0 {
		prevPage = 1
	}
	if currPage >= int(totalPages) {
		nextPage = int(totalPages)
	}

	hasPreviousPage := currPage > 1
	hasNextPage := currPage < int(totalPages)

	pagination.TotalRecords = int(totalRecords)
	pagination.Limit = limit
	pagination.CurrentPage = currPage
	pagination.TotalPages = int(totalPages)
	pagination.PreviousPage = prevPage
	pagination.HasPreviousPage = hasPreviousPage
	pagination.NextPage = nextPage
	pagination.HasNextPage = hasNextPage

	return pagination
}
