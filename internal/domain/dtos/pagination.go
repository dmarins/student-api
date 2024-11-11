package dtos

import "strings"

const (
	PAGE                    = "page"
	PAGE_SIZE               = "pageSize"
	SORT_ORDER              = "sortOrder"
	SORT_FIELD              = "sortField"
	ORDER_ASC               = "ASC"
	ORDER_DESC              = "DESC"
	DEFAULT_PAGE_VALUE      = 1
	DEFAULT_PAGE_SIZE_VALUE = 10
)

type (
	PaginationRequest struct {
		Page      int    `json:"page"`
		PageSize  int    `json:"pageSize"`
		SortOrder string `json:"sortOrder"`
		SortField string `json:"sortField"`
	}

	PaginationResponse struct {
		TotalPages  int         `json:"total_pages"`
		CurrentPage int         `json:"current_page"`
		PageSize    int         `json:"page_size"`
		TotalItems  int         `json:"total_items"`
		Items       interface{} `json:"items"`
	}
)

func NewPagination(totalItems, currentPage, pageSize int, items interface{}) *PaginationResponse {
	if pageSize <= 0 {
		pageSize = DEFAULT_PAGE_SIZE_VALUE
	}

	if totalItems <= 0 {
		return &PaginationResponse{
			TotalPages:  DEFAULT_PAGE_VALUE,
			CurrentPage: DEFAULT_PAGE_VALUE,
			PageSize:    pageSize,
			TotalItems:  totalItems,
			Items:       items,
		}
	}

	return &PaginationResponse{
		TotalPages:  (totalItems + pageSize - 1) / pageSize,
		CurrentPage: currentPage,
		PageSize:    pageSize,
		TotalItems:  totalItems,
		Items:       items,
	}
}

func (p *PaginationRequest) IsASC() bool {
	return p.SortOrder == "" || strings.EqualFold(ORDER_ASC, p.SortOrder)
}

func (p *PaginationRequest) Offset() int {
	return (p.Page - 1) * p.PageSize
}
