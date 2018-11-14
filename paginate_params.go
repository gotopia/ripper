package ripper

// IPaginateParams is an interface for paginate params.
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
func NewPaginateParams(page int, parent string, pageSize int, orderBy string, filter string) *PaginateParams {
	if page == 0 {
		page = defaultPage
	}
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &PaginateParams{
		page,
		parent,
		pageSize,
		orderBy,
		filter,
	}
}

// Itop converts IPaginateParams to PaginateParams.
func Itop(i IPaginateParams) (p *PaginateParams) {
	return NewPaginateParams(defaultPage, i.GetParent(), int(i.GetPageSize()), i.GetOrderBy(), i.GetFilter())
}
