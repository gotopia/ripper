package ripper

const defaultPage = 1

var defaultPageSize = 15
var maxPageSize = 100

// Paginate the results by page_token, returns offset, limit and next_page_token.
func Paginate(p *PaginateParams, pageToken string, totalSize int64) (offset int, limit int, nextPageToken string, err error) {
	if pageToken != "" {
		err = validatePageToken(p, pageToken)
		if err != nil {
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

// SetMaxPageSize sets the max limit to pageSize.
func SetMaxPageSize(pageSize int) bool {
	if pageSize < defaultPageSize {
		return false
	}
	maxPageSize = pageSize
	return true
}
