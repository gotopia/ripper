package ripper

import (
	"github.com/pkg/errors"
)

const defaultPage = 1

var defaultPageSize = 15
var maxPageSize = 15

// Paginate the results by page_token, returns offset, limit and next_page_token.
func Paginate(p *PaginateParams, pageToken string, totalSize int) (offset int, limit int, nextPageToken string, err error) {
	if pageToken != "" {
		err = validatePageToken(p, pageToken)
		if err != nil {
			err = errors.Wrap(err, "invalid token")
			return
		}
	}
	return paginate(p, totalSize)
}

// SetDefaultPageSize sets the default value of pageSize.
func SetDefaultPageSize(pageSize int) bool {
	if pageSize <= 0 {
		return false
	}
	defaultPageSize = pageSize
	return true
}

// SetMaxPageSize sets the max limitation of pageSize.
func SetMaxPageSize(pageSize int) bool {
	if pageSize < defaultPageSize {
		return false
	}
	maxPageSize = pageSize
	return true
}
