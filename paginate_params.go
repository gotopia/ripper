package ripper

import "github.com/pkg/errors"

// IPaginateParams is a interface for request params.
type IPaginateParams interface {
	GetParent() string
	GetPageSize() int32
	GetPageToken() string
	GetOrderBy() string
	GetFilter() string
}

// PaginateParams is the request params for pagination.
type PaginateParams struct {
	Page     int
	Parent   string
	PageSize int
	OrderBy  string
	Filter   string
}

// NewPaginateParams returns a PaginateParams pointer.
func NewPaginateParams(page int, parent string, pageSize int, orderBy string, filter string) (p *PaginateParams, err error) {
	if page == 0 {
		page = defaultPage
	}
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	if page < 0 {
		err = errors.New("page must be larger than zero")
		return
	}
	if pageSize < 0 {
		err = errors.New("pageSize must be larger than zero")
		return
	}
	p = &PaginateParams{
		page,
		parent,
		pageSize,
		orderBy,
		filter,
	}
	return
}

// Itop converts IPaginateParams to PaginateParams.
func Itop(i IPaginateParams) (p *PaginateParams) {
	parent := i.GetParent()
	pageSize := i.GetPageSize()
	orderBy := i.GetOrderBy()
	filter := i.GetFilter()
	p, err := NewPaginateParams(defaultPage, parent, int(pageSize), orderBy, filter)
	if err != nil {
		panic("fail to convert IPaginateParams to PaginateParams")
	}
	return
}
