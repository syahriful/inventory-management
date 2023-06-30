package util

import (
	"github.com/gofiber/fiber/v2"
	"inventory-management/backend/internal/http/response"
	"math"
)

func CreatePagination(ctx *fiber.Ctx, totalRecords int64) (pagination response.Pagination, offset int) {
	currPage := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)
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
	offset = (currPage - 1) * limit

	pagination.TotalRecords = int(totalRecords)
	pagination.Limit = limit
	pagination.CurrentPage = currPage
	pagination.TotalPages = int(totalPages)
	pagination.PreviousPage = prevPage
	pagination.HasPreviousPage = hasPreviousPage
	pagination.NextPage = nextPage
	pagination.HasNextPage = hasNextPage

	return pagination, offset
}
