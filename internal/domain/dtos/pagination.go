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
	PaginationInput struct {
		Page      int     `query:"page"`
		PageSize  int     `query:"pageSize"`
		SortOrder *string `query:"sortOrder"`
		SortField *string `query:"sortField"`
	}

	PaginationOutput struct {
		TotalPages  int         `json:"total_pages"`
		CurrentPage int         `json:"current_page"`
		PageSize    int         `json:"page_size"`
		TotalItems  int         `json:"total_items"`
		Items       interface{} `json:"items"`
	}
)

func NewPaginationInput(page, pageSize int, sortOrder, sortField *string) *PaginationInput {
	if page <= 0 {
		page = DEFAULT_PAGE_VALUE
	}

	if pageSize <= 0 {
		pageSize = DEFAULT_PAGE_SIZE_VALUE
	}

	return &PaginationInput{
		Page:      page,
		PageSize:  pageSize,
		SortOrder: sortOrder,
		SortField: sortField,
	}
}

func NewPaginationOutput(totalItems, currentPage, pageSize int, items interface{}) *PaginationOutput {
	if pageSize <= 0 {
		pageSize = DEFAULT_PAGE_SIZE_VALUE
	}

	if totalItems <= 0 {
		return &PaginationOutput{
			TotalPages:  DEFAULT_PAGE_VALUE,
			CurrentPage: DEFAULT_PAGE_VALUE,
			PageSize:    pageSize,
			TotalItems:  totalItems,
			Items:       items,
		}
	}

	return &PaginationOutput{
		TotalPages:  (totalItems + pageSize - 1) / pageSize,
		CurrentPage: currentPage,
		PageSize:    pageSize,
		TotalItems:  totalItems,
		Items:       items,
	}
}

func (p *PaginationInput) IsASC() bool {
	return *p.SortOrder == "" || strings.EqualFold(ORDER_ASC, *p.SortOrder)
}

func (p *PaginationInput) Offset() int {
	return (p.Page - 1) * p.PageSize
}
