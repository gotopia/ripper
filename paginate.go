package ripper

import "github.com/pkg/errors"

func paginate(p *PaginateParams, totalSize int) (offset int, limit int, nextPageToken string, err error) {
	if p.Page <= 0 {
		err = errors.New("Page must be larger than zero")
		return
	}
	if p.PageSize <= 0 {
		err = errors.New("PageSize must be larger than zero")
		return
	}
	offset = (p.Page - 1) * p.PageSize
	limit = p.PageSize
	if hasNextPage(p.Page, p.PageSize, totalSize) {
		nextPage := p.Page + 1
		nextPaginateParams, _ := NewPaginateParams(nextPage, p.Parent, p.PageSize, p.OrderBy, p.Filter)
		nextPageToken, err = encodeToPageToken(nextPaginateParams)
		if err != nil {
			err = errors.Wrap(err, "fail to encode nextPaginateParams to nextPageToken")
			return
		}
	}
	return
}

func hasNextPage(page int, pageSize int, totalSize int) bool {
	return page*pageSize < totalSize
}
