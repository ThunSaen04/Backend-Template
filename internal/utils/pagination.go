package utils

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

// PaginationMeta holds pagination information returned in API responses.
type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	TotalRows   int64 `json:"total_rows"`
	TotalPages  int   `json:"total_pages"`
}

// PaginationParams holds parsed pagination query parameters.
type PaginationParams struct {
	Page  int
	Limit int
}

const (
	DefaultPage  = 1
	DefaultLimit = 10
	MaxLimit     = 100
)

// ParsePagination extracts page and limit from query string with sensible defaults.
//
// Query parameters:
//   - page  (default: 1, min: 1)
//   - limit (default: 10, min: 1, max: 100)
//
// Example URL: /api/v1/items?page=2&limit=20
func ParsePagination(c fiber.Ctx) *PaginationParams {
	page, _ := strconv.Atoi(c.Query("page", strconv.Itoa(DefaultPage)))
	limit, _ := strconv.Atoi(c.Query("limit", strconv.Itoa(DefaultLimit)))

	if page < 1 {
		page = DefaultPage
	}
	if limit < 1 {
		limit = DefaultLimit
	}
	if limit > MaxLimit {
		limit = MaxLimit
	}

	return &PaginationParams{
		Page:  page,
		Limit: limit,
	}
}

// Paginate returns a GORM scope function that applies OFFSET and LIMIT
// based on the given PaginationParams. Use with db.Scopes().
//
// Example:
//
//	var items []models.Item
//	db.Scopes(utils.Paginate(params)).Find(&items)
func Paginate(params *PaginationParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (params.Page - 1) * params.Limit
		return db.Offset(offset).Limit(params.Limit)
	}
}

// CalculateMeta computes pagination metadata from params and total row count.
func CalculateMeta(params *PaginationParams, totalRows int64) *PaginationMeta {
	totalPages := int(math.Ceil(float64(totalRows) / float64(params.Limit)))

	return &PaginationMeta{
		CurrentPage: params.Page,
		PerPage:     params.Limit,
		TotalRows:   totalRows,
		TotalPages:  totalPages,
	}
}
